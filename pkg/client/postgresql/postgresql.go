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
	"github.com/wtkeqrf0/you-together/ent/intercept"
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

	client.Intercept(queryLogger(log.NewLogger(log.InfoLevel, &log.TextFormatter{}, false)))

	if err = client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true)); err != nil {
		log.WithErr(err).Fatal("tables initialization failed")
	}

	return client
}

func queryLogger(l *log.Logger) ent.InterceptFunc {
	return func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			q, err := intercept.NewQuery(query)
			if err != nil {
				return nil, err
			}

			start := time.Now()
			defer func() {
				l.Infof(" Duration=%s | Schema=%s", time.Since(start), q.Type())
			}()
			return next.Query(ctx, query)
		})
	}
}
