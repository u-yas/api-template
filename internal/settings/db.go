package settings

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func loadDbEnv() string {

	result := "user=" + os.Getenv("POSTGRES_USER") + " port=" + os.Getenv("POSTGRES_PORT") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " host=" + os.Getenv("POSTGRES_HOST") + " sslmode=" + os.Getenv("POSTGRES_SSLMODE") + " TimeZone=" + os.Getenv("POSTGRES_TIMEZONE")
	return result
}

func SetupDB() *sql.DB {
	// postgres settings
	con, err := sql.Open("postgres", loadDbEnv())
	if err != nil {
		panic(err)
	}
	// connection pool settings
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(300 * time.Second)

	// global connection setting
	boil.SetDB(con)

	return con
}
