package model

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	Id             int64
// 	Email          string `sql:"type:text;"`
// 	HashedPassword string
// 	CreatedAt      time.Time
// 	UpdatedAt      time.Time
// }

// type Site struct {
// 	Domain string `sql:"type:text;"`
// 	Key    string
// }

var DB gorm.DB

type Request struct {
	Id             int64
	SessionId      string
	Url            string `sql:"type:text;"`
	Header         string `sql:"type:text;"`
	Domain         string `sql:"type:text;"`
	Referer        string `sql:"type:text;"`
	UserAgent      string `sql:"type:text;"`
	Ip             string
	AcceptLanguage string `sql:"type:text;"`
	CreatedAt      time.Time
}

type User struct {
	Id             int64
	Email          string `sql:"type:text;"`
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func clear(b []byte) {
	// excessive?
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func hashPassword(password string) (hash string, err error) {
	bytePassword := []byte(password)
	defer clear(bytePassword)
	byteHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(byteHash), err
}

func NewUser(email string, password string) (user User, err error) {
	hash, err := hashPassword(password)
	if err != nil {
		return
	}
	user = User{
		Email:          email,
		HashedPassword: hash,
	}
	query := DB.Create(&user)
	if query.Error != nil {
		return user, query.Error
	}
	return
}

func ValidateUserPassword(email, password string) (user User, err error) {
	DB.Where("email = ?", email).First(&user)
	bytePassword := []byte(password)
	byteHash := []byte(user.HashedPassword)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	return user, err
}

func init() {
	var err error

	host := os.Getenv("KAYOBE_HOST")
	if host != "" {
		port := os.Getenv("KAYOBE_PORT")
		dbname := os.Getenv("KAYOBE_DBNAME")
		username := os.Getenv("KAYOBE_USERNAME")
		password := os.Getenv("KAYOBE_PASSWORD")
		configString := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " sslmode=disable dbname=" + dbname
		DB, err = gorm.Open("postgres", configString)
	} else {
		DB, err = gorm.Open("postgres", "dbname=kayobe sslmode=disable")
	}

	if err != nil {
		fmt.Println(err)
	}
	query := DB.AutoMigrate(&Request{}, &User{})
	if query.Error != nil {
		fmt.Println(query.Error)
	}
}
