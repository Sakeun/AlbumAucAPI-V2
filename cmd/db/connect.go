package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Album struct {
	Id         int    `json:"id"`
	SellerId   int    `json:"sellerId"`
	Name       string `json:"name"`
	Genre      string `json:"genre"`
	Condition  string `json:"condition"`
	EndingTime string `json:"endingTime"`
	IsDone     bool   `json:"isDone"`
	Bids       int    `json:"bids"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Rating   int    `json:"rating"`
	Country  string `json:"country"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
	Password string `json:"password"`
}

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

func GetUser(name string) User {
	c := getCredentials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM user WHERE username = ?", name)
	if err != nil {
		panic(err.Error())
	}

	var user User

	for results.Next() {
		err = results.Scan(&user.Id, &user.Username, &user.Rating, &user.Country, &user.Email, &user.IsAdmin, &user.Password)
		if err != nil {
			panic(err.Error())
		}
	}

	if user.Username == "" {
		return User{}
	}

	return user
}

func getCredentials() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Username = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.DBName = os.Getenv("DB_NAME")
	config.Database.Port, err = strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	return config
}
