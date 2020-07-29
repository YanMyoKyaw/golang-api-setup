package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"test/model"
	"test/utils"
)

var db *gorm.DB

func connectDB(config model.DBProperties) bool {
	var err error
	if db == nil {
		if config.DBType == utils.MYSQL {
			settingString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.UserName, config.Password, config.Host, config.Port, config.Database)
			log.Println(settingString)
			db, err = gorm.Open("mysql", settingString)
			if err != nil {
				log.Println(err.Error())
				return false
			}
		}
	}
	return true

}

func InitializeDB(config model.DBProperties) bool {
	log.Println("start InitializeDB function")
	if ok := connectDB(config); !ok {
		log.Println("Error in connecting databse. Quit InitializeDB function")
		return false
	}

	if !db.HasTable(model.User{}) {
		db.CreateTable(model.User{})
	}
	if !db.HasTable(model.Product{}) {
		db.CreateTable(model.Product{})
	}
	return true
}

func CreateUser(user *model.User) error {
	db.Create(user)
	return errors.New("empty errro")
}

func GetUserByMail(mail string) *model.User {
	var user model.User
	db.Where("email = ?", mail).First(&user)
	return &user
}

func readFileAsString(filePath string) ([]byte, error) {
	var content []byte
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return content, err
	}
	return content, nil
}

func GetDBProperties() (model.DBProperties, error) {
	var result = model.DBProperties{}
	content, err := readFileAsString(utils.DB_CONFIG)
	if err != nil {
		return result, err
	}
	if err = json.Unmarshal(content, &result); err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}
