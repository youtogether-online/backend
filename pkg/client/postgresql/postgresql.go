package postgresql

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/wtkeqrf0/you-together/ent"
	_ "github.com/wtkeqrf0/you-together/ent/runtime"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"time"
)

// Open postgres connection, check it and create tables (if not exist)
func Open(username, password, host string, port int, DBName string) *ent.Client {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		username, password, host, port, DBName)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.WithErr(err).Fatal("error occurred while opening PostgreSQL connection")
	}

	if err = db.Ping(); err != nil {
		log.WithErr(err).Fatal("unable to connect to the postgres database")
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err = client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true)); err != nil {
		log.WithErr(err).Fatal("tables initialization failed")
	}

	client.Use(dbLogger)

	return client
}

var logger = log.NewLogger(log.InfoLevel, &log.TextFormatter{}, false)

func dbLogger(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		start := time.Now()
		defer func() {
			logger.Infof("Op=%s\tType=%s\tTime=%s", m.Op(), m.Type(), time.Since(start))
		}()
		return next.Mutate(ctx, m)
	})
}
