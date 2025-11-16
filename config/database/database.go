package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host		string
	Port		string
	Password	string
	User		string
	DBName		string
	SSLMode		string
}

var DB *gorm.DB

func NewConnection(config *Config)(error){
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("An error occured while establishing a connection with the database!")
		return  err
	}
	DB = db
	return nil
}

func Migrate(tables ...interface{}) error{
	return DB.AutoMigrate(tables...)
}