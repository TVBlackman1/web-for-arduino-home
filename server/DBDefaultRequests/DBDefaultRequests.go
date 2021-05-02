package DBDefaultRequests

import (
	dbDefaultEssence "../DBDefaultEssence"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB
var err error

func Connect() {
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ArduinoSmartHome")
	if err != nil {
		panic(err.Error())
	}
}

func CreateUser(user *dbDefaultEssence.User) error {

	_, err := GetUserByLogin(user.Login)
	if err == nil {
		// we found user in database
		return errors.New("user already exists")
	}

	query := fmt.Sprintf("INSERT INTO Users (Login, Password) VALUES ('%s', '%s')", user.Login, user.Password)
	fmt.Println(query)
	rows, err := Db.Query(query)
	if err != nil {
		return err
	}

	fmt.Println(rows)
	return nil
}

func GetUserByLogin(login string) (*dbDefaultEssence.User, error) {
	//user
	query := fmt.Sprintf("select Login, Password from Users where Login = '%s'", login)
	fmt.Println(query)


	results, err := Db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var resultUser dbDefaultEssence.User

		err = results.Scan(&resultUser.Login, &resultUser.Password)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(resultUser.Login)
		log.Printf(resultUser.Password)

		return &resultUser, nil
	}

	return nil, errors.New("user is not exists")
}
