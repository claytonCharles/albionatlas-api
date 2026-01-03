package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/claytonCharles/albionatlas-api/internal/auth"
	_ "github.com/lib/pq"
)

type ContainerDB struct {
	DB       *sql.DB
	AuthRepo auth.AuthRepository
}

func NewConnection() *ContainerDB {
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database conectada com sucesso!")
	return &ContainerDB{DB: db}
}

func (cdb *ContainerDB) InitializeRepositories() {
	cdb.AuthRepo = auth.NewRepository(cdb.DB)
}
