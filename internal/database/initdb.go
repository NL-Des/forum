package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(path string) *sql.DB {
	DB, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("❌ Erreur ouverture DB :", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Erreur connexion DB :", err)
	}

	log.Println("✅ Base de données ouverte")
	return DB
}
