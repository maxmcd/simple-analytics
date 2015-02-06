package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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
	Id               int64
	SessionId        string
	SessionTimestamp int64
	Url              string `sql:"type:text;"`
	Header           string `sql:"type:text;"`
	Domain           string `sql:"type:text;"`
	Referer          string `sql:"type:text;"`
	UserAgent        string `sql:"type:text;"`
	Ip               string
	AcceptLanguage   string `sql:"type:text;"`
}

func init() {
	var err error

	DB, err = gorm.Open("postgres", "dbname=kayobe sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	query := DB.AutoMigrate(&Request{})
	if query.Error != nil {
		fmt.Println(query.Error)
	}
}
