package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	_ "github.com/wtkeqrf0/you-together/docs"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller"
	"github.com/wtkeqrf0/you-together/internal/repo/postgres"
	redis2 "github.com/wtkeqrf0/you-together/internal/repo/redis"
	"github.com/wtkeqrf0/you-together/internal/service"
	"github.com/wtkeqrf0/you-together/pkg/bind"
	"github.com/wtkeqrf0/you-together/pkg/client/email"
	"github.com/wtkeqrf0/you-together/pkg/client/postgresql"
	redisDB "github.com/wtkeqrf0/you-together/pkg/client/redis"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/middleware/sessions"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006/01/02 15:32:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "status",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})
	logrus.SetReportCaller(true)

	logrus.SetReportCaller(true)

	// global validator of incoming values
	binding.Validator = bind.NewValid(validator.New())
}

// @title You Together API
// @version 1.0
// @description It's an API interacting with You Together using Golang
// @accept application/json
// @produce application/json
// @schemes http

// @host localhost:3000
// @BasePath /api

// @sessions.docs.description Authorization, registration and authentication
func main() {
	cfg := conf.GetConfig()

	pClient, rClient, mailClient := getClients(cfg)

	pConn, rConn := postgres.NewUserStorage(pClient.User), redis2.NewRClient(rClient)
	auth := service.NewAuthService(pConn, rConn)

	h := controller.NewHandler(
		service.NewUserService(pConn, rConn),
		sessions.NewAuth(auth),
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
		logrus.WithError(err).Fatal("Email Connection Shutdown Failed")
	}

	logrus.Info("Server Exited Properly")
}

func getClients(cfg *conf.Config) (*ent.Client, *redis.Client, *smtp.Client) {
	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisDB.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Port, cfg.DB.Redis.DbId)

	mailClient := email.Open(cfg.Email.User, cfg.Email.Password, cfg.Email.Host, cfg.Email.Port)

	return pClient, rClient, mailClient
}
