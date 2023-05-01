package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	_ "github.com/wtkeqrf0/you-together/docs"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller"
	"github.com/wtkeqrf0/you-together/internal/middleware/authorization"
	"github.com/wtkeqrf0/you-together/internal/repo/postgres"
	redis2 "github.com/wtkeqrf0/you-together/internal/repo/redis"
	"github.com/wtkeqrf0/you-together/internal/service"
	"github.com/wtkeqrf0/you-together/pkg/client/email"
	"github.com/wtkeqrf0/you-together/pkg/client/postgresql"
	redisDB "github.com/wtkeqrf0/you-together/pkg/client/redis"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:32:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "status",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})

	logrus.SetReportCaller(true)
}

// @title You Together API
// @version 1.0
// @description It's an API interacting with You Together using Golang
// @accept application/json
// @produce application/json
// @schemes http

// @host localhost:3000
// @BasePath /api

// @authorization.docs.description Authorization, registration and authentication
func main() {
	cfg := conf.GetConfig()

	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisDB.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Port, cfg.DB.Redis.DbId)

	pConn, rConn := postgres.NewUserStorage(pClient.User), redis2.NewRClient(rClient)
	auth := service.NewAuthService(pConn, rConn)

	mailClient := email.Open(cfg.Email.User, cfg.Email.Password, cfg.Email.Host, cfg.Email.Port)

	h := controller.NewHandler(
		service.NewUserService(pConn, rConn),
		authorization.NewAuth(auth),
		auth,
		service.NewEmailSender(mailClient),
	)

	r := gin.New()
	h.InitRoutes(r, mailClient != nil)

	Run(cfg.Listen.Port, r, pClient, rClient, mailClient)
}

// Run the Server with graceful shutdown
func Run(port int, r *gin.Engine, pClient *ent.Client, rClient *redis.Client, mailClient *smtp.Client) {
	srv := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatalf("error occurred while running http server")
		}
	}()

	logrus.Infof("Server Started On Port %d", port)

	<-quit

	logrus.Info("Server Shutting Down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("Server Shutdown Failed")
	}

	if err := rClient.Close(); err != nil {
		logrus.WithError(err).Fatal("Redis Connection Shutdown Failed")
	}

	if err := pClient.Close(); err != nil {
		logrus.WithError(err).Fatal("PostgreSQL Connection Shutdown Failed")
	}

	if err := mailClient.Quit(); err != nil {
		logrus.WithError(err).Fatal("Email QUIT Failed")
	}

	if err := mailClient.Close(); err != nil {
		logrus.WithError(err).Fatal("Email Connection Shutdown Failed")
	}

	logrus.Info("Server Exited Properly")
}
