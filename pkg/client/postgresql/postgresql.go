package postgresql

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/you_together/ent"
	"github.com/wtkeqrf0/you_together/ent/migrate"
)

// Open postgres connection, check it and create tables (if not exist). Returns the client of defined postgres database
func Open(userName, password, host, port, DBName string) *ent.Client {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		userName, password, host, port, DBName)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		logrus.Fatalf("error occured while opening PostgreSQL connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		logrus.Fatalf("Unable to connect to the postgres database: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		logrus.Fatalf("Tables Initialization Failed: %s\n", err)
	}
	return client
}
