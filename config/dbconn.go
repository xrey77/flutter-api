// *** Author : Reynald Marquez-Gragasin
// *** Email : reynald88@yahoo.com
// *** Program Name : FLUTTER-API

package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func Connect() *pg.DB {

	var (
		opts *pg.Options
		err  error
	)
	opts, err = pg.ParseURL(os.Getenv("DATABASE_URL"))

	// opts = &pg.Options{
	// 	User:     "rey",
	// 	Password: "rey",
	// 	Addr:     "127.0.0.1:8090",
	// 	Database: "global-api-10",
	// }

	if err != nil {
		log.Print("Unable to connect to Postgres Database")
		return nil
	} else {
		log.Print("Connected to Postgres DAtabase")
	}
	db := pg.Connect(opts)
	return db

}
