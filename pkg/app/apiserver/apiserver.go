package apiserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GaponovAlexey/learn-rest/pkg/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // ...
	"github.com/sirupsen/logrus"
)

// database_url = "host=localhost dbname=restapi_dev sslmode=disable"
// session_key = "8109612acafb4ae0ff34f5f1fa549577f4ca3a4a294f559498c111cc7d92973e5dde4eb64f086b49e063708705338f29b662047c09c850f5bb21da65f37036b4"

// Start ...
func Start(config *Config) error {
	db, err := newDB("host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte("8109612acafb4ae0ff34f5f1fa549577f4ca3a4a294f559498c111cc7d92973e5dde4eb64f086b49e063708705338f29b662047c09c850f5bb21da65f37036b4"))
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(":3000", srv)
}

func newDB(dbURL string) (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USERNAME"), os.Getenv("DBNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("SSLMODE")))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
