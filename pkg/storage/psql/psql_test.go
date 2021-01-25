package psql_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"testing"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/storage"
	"github.com/filariow/bshop/pkg/storage/psql"
	"github.com/jackc/pgx/v4"
	"github.com/matryer/is"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

const (
	host   = "0.0.0.0"
	user   = "root"
	passwd = "supersecret"
	dbName = "test"
)

var (
	conn   *pgx.Conn
	dbPort string
)

func connectionString(port, db string) string {
	cs := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, host, port, db)
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
	dbPort = resource.GetPort("5432/tcp")

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

type testData struct {
	conn    *pgx.Conn
	ctx     context.Context
	counter int32
	repo    storage.BeerRepository
}

var tdcounter int32 = 0

func newTestData() *testData {
	return &testData{
		ctx:     context.Background(),
		counter: atomic.AddInt32(&tdcounter, 1),
	}
}

func (t *testData) databaseName() string {
	return fmt.Sprintf("%s_%d", dbName, t.counter)
}

func arrange(t *testing.T) *testData {
	// create database
	td := newTestData()
	db := td.databaseName()
	if _, err := conn.Exec(td.ctx, fmt.Sprintf("CREATE DATABASE %s", db)); err != nil {
		log.Fatalf("error creating database %s: %v", db, err)
	}

	var err error
	pi := connectionString(dbPort, db)
	td.conn, err = pgx.Connect(td.ctx, pi)
	if err != nil {
		log.Fatalf("Error connecting to db %v: %v", dbName, err)
	}

	_, err = td.conn.Exec(td.ctx, `
CREATE TABLE beers(
    id BIGSERIAL PRIMARY KEY
    ,name VARCHAR(100) NOT NULL
    ,price real NOT NULL
    ,cost real
    ,size real
    ,vol real
    ,brewer VARCHAR(200)
	,is_deleted boolean DEFAULT false
)`)
	if err != nil {
		log.Fatalf("Error creating beers table: %v", err)
	}

	td.repo = psql.New(td.conn)
	return td
}

func TestCreateRead(t *testing.T) {
	td := arrange(t)
	b := bshop.Beer{
		Product: bshop.Product{
			Name:  "Name",
			Price: 4.4,
			Cost:  2.0,
		},
		Size:   33.0,
		Vol:    3.5,
		Brewer: "Brewer",
	}
	is := is.New(t)

	id, err := td.repo.Create(td.ctx, b)
	is.NoErr(err)

	rb, err := td.repo.Read(td.ctx, id)
	is.NoErr(err)

	is.Equal(id, rb.ID)
	is.Equal(b.Name, rb.Name)
	assert.InDelta(t, b.Price, rb.Price, 0.001)
	assert.InDelta(t, b.Cost, rb.Cost, 0.001)
	is.Equal(b.Size, rb.Size)
	is.Equal(b.Vol, rb.Vol)
	is.Equal(b.Brewer, rb.Brewer)
}

func TestCreateDelete(t *testing.T) {
	td := arrange(t)
	b := bshop.Beer{
		Product: bshop.Product{
			Name:  "Name",
			Price: 4.4,
			Cost:  2.0,
		},
		Size:   33.0,
		Vol:    3.5,
		Brewer: "Brewer",
	}

	is := is.New(t)
	id, err := td.repo.Create(td.ctx, b)
	is.NoErr(err)
	_, err = td.repo.Read(td.ctx, id)
	is.NoErr(err)
	err = td.repo.Delete(td.ctx, id)
	is.NoErr(err)
	_, err = td.repo.Read(td.ctx, id)
	is.Equal(err, storage.ErrorNotFound)
}

func TestUpdate(t *testing.T) {
	td := arrange(t)
	b := bshop.Beer{
		Product: bshop.Product{
			Name:  "Name",
			Price: 4.4,
			Cost:  2.0,
		},
		Size:   33.0,
		Vol:    3.5,
		Brewer: "Brewer",
	}
	is := is.New(t)
	id, err := td.repo.Create(td.ctx, b)
	is.NoErr(err)

	bu := bshop.Beer{
		Product: bshop.Product{
			ID:    id,
			Name:  "Name Update",
			Price: 5.4,
			Cost:  3.0,
		},
		Size:   66.0,
		Vol:    5.5,
		Brewer: "Brewer TestUpdate",
	}

	err = td.repo.Update(td.ctx, bu)
	is.NoErr(err)

	rb, err := td.repo.Read(td.ctx, id)
	is.NoErr(err)

	is.Equal(bu.ID, rb.ID)
	is.Equal(bu.Name, rb.Name)
	assert.InDelta(t, bu.Price, rb.Price, 0.001)
	assert.InDelta(t, bu.Cost, rb.Cost, 0.001)
	is.Equal(bu.Size, rb.Size)
	is.Equal(bu.Vol, rb.Vol)
	is.Equal(bu.Brewer, rb.Brewer)
}

func TestList(t *testing.T) {
	td := arrange(t)
	b := bshop.Beer{
		Product: bshop.Product{
			Name:  "Name",
			Price: 4.4,
			Cost:  2.0,
		},
		Size:   33.0,
		Vol:    3.5,
		Brewer: "Brewer",
	}
	is := is.New(t)
	id1, err := td.repo.Create(td.ctx, b)
	is.NoErr(err)

	id2, err := td.repo.Create(td.ctx, b)
	is.NoErr(err)

	rrb, err := td.repo.List(td.ctx)
	is.NoErr(err)

	for k, r := range rrb {
		if k == 0 {
			is.Equal(id1, r.ID)
		} else {
			is.Equal(id2, r.ID)
		}
		is.Equal(b.Name, r.Name)
		assert.InDelta(t, b.Price, r.Price, 0.001)
		assert.InDelta(t, b.Cost, r.Cost, 0.001)
		is.Equal(b.Size, r.Size)
		is.Equal(b.Vol, r.Vol)
		is.Equal(b.Brewer, r.Brewer)
	}
}
