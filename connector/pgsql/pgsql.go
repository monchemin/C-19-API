package pgsql

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func Open() (*DB, error) {
	var (
		hostName     = os.Getenv("POSTGRES_HOSTNAME")
		userName     = os.Getenv("POSTGRES_USERNAME")
		password     = os.Getenv("POSTGRES_PASSWORD")
		port         = os.Getenv("POSTGRES_PORT")
		dataBaseName = os.Getenv("POSTGRES_DB")
	)
	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	return OpenWithConfig(Config{
		HostName:     hostName,
		UserName:     userName,
		Password:     password,
		Port:         p,
		DataBaseName: dataBaseName,
	})
}

func OpenWithConfig(config Config) (*DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.HostName, config.Port, config.UserName, config.Password, config.DataBaseName)

	println(connectionString)
	connection := sqlx.MustOpen("postgres", connectionString)
	err := connection.Ping()
	if err != nil {
		panic(err)
	}


	return &DB{connection}, nil
}
