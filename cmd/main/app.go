package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/internal/controller"
	"github.com/wtkeqrf0/you_together/internal/repo"
	"github.com/wtkeqrf0/you_together/internal/service"
	"github.com/wtkeqrf0/you_together/pkg/client/postgresql"
	redisDB "github.com/wtkeqrf0/you_together/pkg/client/redis"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"github.com/wtkeqrf0/you_together/pkg/middlewares/authorization"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	{
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

	cfg := conf.GetConfig()

	pClient := postgresql.Open(cfg.DB.Postgres.UserName, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisDB.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Port, cfg.DB.Redis.DB)

	pConn := repo.NewUserStorage(pClient.User)

	h := controller.NewHandler(
		service.NewUserService(pConn),
		service.NewRClient(rClient),
		authorization.NewAuth(pConn),
	)

	r := gin.New()
	h.InitRoutes(r)

	Run(cfg.Listen.Port, r, pClient, rClient)
}

// Run the Server with graceful shutdown
func Run(port string, r *gin.Engine, pClient *ent.Client, rClient *redis.Client) {
	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occurred while running http server: %s", err)
		}
	}()

	logrus.Infof("Server Started On Port %s", port)

	<-quit

	logrus.Info("Server Shutting Down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown Failed: %s", err)
	}

	if err := rClient.Close(); err != nil {
		logrus.Fatalf("Redis Connection Shutdown Failed: %s", err)
	}

	if err := pClient.Close(); err != nil {
		logrus.Fatalf("PostgreSQL Connection Shutdown Failed: %s", err)
	}

	logrus.Info("Server Exited Properly")
}
