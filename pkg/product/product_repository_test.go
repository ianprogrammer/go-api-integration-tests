package product_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/ianprogrammer/go-api-integration-test/config"
	"github.com/ianprogrammer/go-api-integration-test/internal/database"
	"github.com/ianprogrammer/go-api-integration-test/pkg/product"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

const succeed = "\u2713"
const failed = "\u2717"

var (
	_sut  *product.Repository
	_conn *gorm.DB
)

func TestMain(m *testing.M) {
	log.Println("Starting postgres container...")
	postgresPort := nat.Port("5432/tcp")

	postgres, err := tc.GenericContainer(context.Background(),
		tc.GenericContainerRequest{
			ContainerRequest: tc.ContainerRequest{
				Image:        "postgres:12.2-alpine",
				ExposedPorts: []string{postgresPort.Port()},
				Env: map[string]string{
					"POSTGRES_PASSWORD": "pass",
					"POSTGRES_USER":     "user",
					"POSTGRES_DB":       "product",
				},
				WaitingFor: wait.ForAll(
					wait.ForLog("database system is ready to accept connections"),
					wait.ForListeningPort(postgresPort),
				),
			},
			Started: true,
		})

	if err != nil {
		log.Fatal("start:", err)
	}

	hostPort, err := postgres.MappedPort(context.Background(), postgresPort)
	if err != nil {
		log.Fatal("map:", err)
	}

	port, err := strconv.Atoi(hostPort.Port())

	if err != nil {
		log.Fatal("convert port to int went wrong", err)
	}

	databaseConfig := config.DatabaseConfig{
		Host:         "localhost",
		UserName:     "user",
		Password:     "pass",
		DatabaseName: "postgres",
		DatabasePort: port,
	}

	_conn, err = database.NewDatabase(databaseConfig)

	if err != nil {
		log.Fatal("could not be possible to connect to database")
	}

	path, err := filepath.Abs("../../db/migrations")

	if err != nil {
		log.Fatal("could not be possible to get migration path")
	}

	path = fmt.Sprintf("file://%s", path)

	err = database.MigrateDB(databaseConfig, path)

	if err != nil {
		log.Fatal("could not be possible to migrate database ", err)
	}
	_sut = &product.Repository{
		DB: _conn,
	}
	os.Exit(m.Run())
}

func TestRepoImpl(t *testing.T) {

	t.Run("create and get a product", func(t *testing.T) {

		t.Log("should save the product into database and return the product.")
		{
			products := []product.Product{
				{
					Name:  "Test product 1",
					Price: 100,
				},
				{
					Name:  "Test product 3",
					Price: 100,
				},
				{
					Name:  "Test product 2",
					Price: 100000,
				},
			}
			var countInserted int
			for i, p := range products {
				t.Logf("\tTest %d: When passing product  with name of %s and price of %d", i, p.Name, p.Price)
				{
					product, err := _sut.Insert(p)
					require.NoError(t, err)
					if product.Name == p.Name {
						t.Logf("\t%s\t  Product %s was inserted.", succeed, p.Name)
					} else {
						t.Logf("\t%s\t  Product %s was failed.", failed, p.Name)
					}
					countInserted++
				}
			}
			assert.Equal(t, countInserted, len(products))
		}
	})
}
