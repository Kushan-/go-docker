package database

import (
	"github.com/Kushan-/go-docker/pkg/config"
	"github.com/Kushan-/go-docker/pkg/model"

	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DBConn, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))	)	

	if err != nil {
		fmt.Println("Unable to connect to DB")
		panic(err.Error())
	}

	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&model.Product{}, &model.User{})
	fmt.Println("Database Migrated")
}

func Close() (err error) {
	return DBConn.Close()
}