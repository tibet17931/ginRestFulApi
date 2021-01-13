package db

import (
	"ginRestFulApi/utils"
	"log"

	"github.com/goonode/mogo"
)

var mongoConnection *mogo.Connection = nil

//GetConnection is for get mongo connection
func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		dbName := utils.EnvVar("DB_NAME", "")
		config := &mogo.Config{
			ConnectionString: "mongodb+srv://yeen:<password>@cluster0.tezwz.mongodb.net/" + dbName + "?retryWrites=true&w=majority",
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}
	return mongoConnection
}
