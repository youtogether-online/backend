package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller"
	"github.com/wtkeqrf0/you-together/internal/repo/postgres"
	redisRepo "github.com/wtkeqrf0/you-together/internal/repo/redis"
	"github.com/wtkeqrf0/you-together/internal/service"
	"github.com/wtkeqrf0/you-together/pkg/bind"
	"github.com/wtkeqrf0/you-together/pkg/client/email"
	"github.com/wtkeqrf0/you-together/pkg/client/postgresql"
	redisInit "github.com/wtkeqrf0/you-together/pkg/client/redis"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"github.com/wtkeqrf0/you-together/pkg/middleware/query"
	"github.com/wtkeqrf0/you-together/pkg/middleware/session"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// global validator of incoming values
	binding.Validator = bind.NewValid(validator.New())

	cfg := conf.GetConfig()

	pClient, rClient, mailClient := getClients(cfg)

	h := initHandler(pClient, rClient, mailClient)
	m := initMiddlewares()

	r := gin.New()

	m.InitGlobalMiddleWares(r)
	h.InitRoutes(r.Group(cfg.Listen.MainPath), mailClient != nil)

	run(cfg.Listen.Port, r, pClient, rClient, mailClient)
}

// run the Server with graceful shutdown
func run(port int, r *gin.Engine, pClient *ent.Client, rClient *redis.Client, mailClient *smtp.Client) {
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

	if err := mailClient.Quit(); err != nil {
		log.WithErr(err).Fatal("Email Connection Shutdown Failed")
	}

	log.Info("Server Exited Properly")
}

func getClients(cfg *conf.Config) (*ent.Client, *redis.Client, *smtp.Client) {
	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisInit.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Port, cfg.DB.Redis.DbId)

	mailClient := email.Open(cfg.Email.User, cfg.Email.Password, cfg.Email.Host, cfg.Email.Port)

	return pClient, rClient, mailClient
}

func initHandler(pClient *ent.Client, rClient *redis.Client, mailClient *smtp.Client) *controller.Handler {
	pUser := postgres.NewUserStorage(pClient.User)
	rConn := redisRepo.NewRClient(rClient)
	mailConn := service.NewEmailSender(mailClient)

	auth := service.NewAuthService(pUser, rConn)
	user := service.NewUserService(pUser, rConn)

	return controller.NewHandler(
		user,
		auth,
		mailConn,
		session.NewAuth(auth),
		bind.NewValid(validator.New()),
	)
}

func initMiddlewares() *controller.Middlewares {
	return controller.NewMiddleWares(
		errs.NewErrHandler(),
		query.NewQueryHandler(),
	)
}
