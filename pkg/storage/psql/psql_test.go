package psql_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	host   = "0.0.0.0"
	user   = "root"
	passwd = "supersecret"
	dbName = "test"
)

var conn *pgx.Conn

func connectionString(port, db string) string {
	cs := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, host, port, db)
	log.Println("Connection string:", cs)
	return cs
}

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13.1-alpine",
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", passwd),
			fmt.Sprintf("POSTGRES_USER=%s", user),
		},
		Mounts: []string{},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	resource.Expire(600)

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		ctx := context.Background()
		pi := connectionString(resource.GetPort("5432/tcp"), "")
		var err error
		conn, err = pgx.Connect(ctx, pi)
		if err != nil {
			return err
		}
		return conn.Ping(ctx)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// create database
	ctx := context.Background()
	if _, err := conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", dbName)); err != nil {
		log.Fatalf("error creating database %s: %v", dbName, err)
	}

	pi := connectionString(resource.GetPort("5432/tcp"), dbName)
	conn, err = pgx.Connect(ctx, pi)
	if err != nil {
		log.Fatalf("Error connecting to db %v: %v", dbName, err)
	}

	_, err = conn.Exec(ctx, `
CREATE TABLE beers(
    id BIGSERIAL PRIMARY KEY
    ,name VARCHAR(100) NOT NULL
    ,price real NOT NULL
    ,cost real
    ,size real
    ,vol real
    ,brewer VARCHAR(200)
)`)
	if err != nil {
		log.Fatalf("Error creating beers table: %v", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSomething(t *testing.T) {
	// db.Query()
}
