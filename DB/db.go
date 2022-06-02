package DB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func NewConnectionDB(driverDB string, database string, host string, user string, password string, port int) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)

	return db, nil

}
func DbInit() (*sql.DB, error) {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}

	db, err := NewConnectionDB(viper.GetString("database.driver"), viper.GetString("database.schema"),
		viper.GetString("database.hostname"), viper.GetString("database.username"), viper.GetString("database.password"),
		viper.GetInt("database.port"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
