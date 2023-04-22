package postgresql

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/ent/migrate"
	"log"
	"time"
)

// Open postgres connection, check it and create tables (if not exist). Returns the client of defined postgres database
func Open(username, password, host string, port int, DBName string) *ent.Client {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		username, password, host, port, DBName)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		logrus.WithError(err).Fatal("error occured while opening PostgreSQL connection")
	}

	if err = db.Ping(); err != nil {
		logrus.WithError(err).Fatal("Unable to connect to the postgres database")
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		logrus.WithError(err).Fatal("Tables Initialization Failed")
	}

	client.Use(logger)

	return client
}

func logger(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		start := time.Now()
		defer func() {
			log.Printf("Op=%s\tType=%s\tTime=%s\n", m.Op(), m.Type(), time.Since(start))
		}()
		return next.Mutate(ctx, m)
	})
}
