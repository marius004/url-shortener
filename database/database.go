package database

import (
	"fmt"

	"github.com/marius004/url-shortener/internal"
	"github.com/marius004/url-shortener/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func ConnectToPSQL(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return &Database{Conn: db}, err
}

func GenerateDatabaseDSN(config *internal.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
}

func (db *Database) AutoMigrate(dst ...interface{}) error {
	return db.Conn.AutoMigrate(dst)
}

func (db *Database) UserService() services.UserService {
	return NewUserService(db)
}
