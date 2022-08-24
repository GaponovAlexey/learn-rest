package store

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Store struct {
	// config *Config
	db     *sqlx.DB
	userRepository *UserRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// // Open
// func (s *Store) Open() error {
// 	err := godotenv.Load()
// 	if err != nil {
// 		logrus.Fatal("Error loading .env file")
// 	}
// 	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USERNAME"), os.Getenv("DBNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("SSLMODE")))
// 	if err != nil {
// 		return err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return err
// 	}

// 	s.db = db

// 	return nil
// }

// // Clothe
// func (s *Store) Close() {
// 	s.db.Close()
// }

//User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
