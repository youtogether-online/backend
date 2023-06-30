package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller"
	"github.com/wtkeqrf0/you-together/internal/repo/postgres"
	redisRepo "github.com/wtkeqrf0/you-together/internal/repo/redis"
	"github.com/wtkeqrf0/you-together/internal/service"
	"github.com/wtkeqrf0/you-together/pkg/client/email"
	"github.com/wtkeqrf0/you-together/pkg/client/postgresql"
	redisInit "github.com/wtkeqrf0/you-together/pkg/client/redis"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"github.com/wtkeqrf0/you-together/pkg/middleware/bind"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"github.com/wtkeqrf0/you-together/pkg/middleware/query"
	"github.com/wtkeqrf0/you-together/pkg/middleware/session"
	"github.com/wtkeqrf0/you-together/pkg/ws"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := conf.GetConfig()

	pClient, rClient, mailClient := getClients(cfg)

	h, sess := initHandler(pClient, rClient, mailClient, cfg)
	r := gin.New()

	h.InitRoutes(createSetter(r, sess))

	run(cfg.Listen.Port, r, pClient, rClient, mailClient)
}

// run the Server with graceful shutdown
func run(port int, r *gin.Engine, pClient *ent.Client, rClient *redis.Client, mailClient *email.MailClient) {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithErr(err).Fatalf("error occurred while running http server")
		}
	}()
	log.Infof("Server Started On Port %d", port)

	<-quit

	log.Info("Server Shutting Down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.WithErr(err).Fatal("Server Shutdown Failed")
	}

	if err := rClient.Close(); err != nil {
		log.WithErr(err).Fatal("Redis Connection Shutdown Failed")
	}

	if err := pClient.Close(); err != nil {
		log.WithErr(err).Fatal("PostgreSQL Connection Shutdown Failed")
	}

	if mailClient != nil {
		if err := mailClient.Close(); err != nil {
			log.WithErr(err).Fatal("Email Connection Shutdown Failed")
		}
	}

	log.LastInfo("Server Exited Properly")
}

func getClients(cfg *conf.Config) (*ent.Client, *redis.Client, *email.MailClient) {
	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisInit.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Password, cfg.DB.Redis.Port, cfg.DB.Redis.DbId)

	mailClient := email.NewMailClient(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)

	return pClient, rClient, mailClient
}

func initHandler(pClient *ent.Client, rClient *redis.Client, mailClient *email.MailClient, cfg *conf.Config) (*controller.Handler, *session.Auth) {
	pUser := postgres.NewUserStorage(pClient.User)
	pRoom := postgres.NewRoomStorage(pClient.Room)
	rConn := redisRepo.NewRClient(rClient)
	webSocket := ws.NewManager(context.Background(), rConn)

	user := service.NewUserService(pUser, rConn)
	room := service.NewRoomService(pRoom)
	auth := service.NewAuthService(pUser, rConn)

	sess := session.NewAuth(auth, cfg)

	return controller.NewHandler(
		user,
		room,
		auth,
		mailClient,
		sess,
		webSocket,
		cfg,
	), sess
}

func createSetter(r *gin.Engine, sess *session.Auth) *controller.Setter {
	return controller.NewSetter(
		r,
		bind.NewValidator(),
		errs.NewErrHandler(),
		query.NewQueryHandler(),
		sess,
	)
}
