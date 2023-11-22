package connect

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"main.go/structures"
)

var connection *gorm.DB

const engine_sql string = "mysql"

const userName string = ""
const password string = ""
const database string = ""

func InitializeDB() {
	connection = ConnectORM(DBString())
	log.Println("Successfully connected")
}

func CloseConnection() {
	connection.Close()
	log.Println("Connection closed")
}

func ConnectORM(stringConnection string) *gorm.DB {

	connection, err := gorm.Open(engine_sql, stringConnection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

func GetUser(id string) structures.User {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	return user
}

func CreateUser(user structures.User) structures.User {
	connection.Create(&user)
	return user
}

func UpdateUser(id string, user structures.User) structures.User {
	currentUser := structures.User{}
	connection.Where("id = ?", id).First(&currentUser)

	currentUser.Name = user.Name
	currentUser.Email = user.Email

	connection.Save(&currentUser)

	return currentUser
}

func DeleteUser(id string) {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	connection.Delete(&user)
}

func DBString() string {
	return userName + ":" + password + "@/" + database
}
