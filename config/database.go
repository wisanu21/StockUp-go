package config

import (
	"database/sql"
	"log"
	"stockup-go/helper"
	"stockup-go/util"

	_redis "github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorm.DB

//RedisClient ...
var RedisClient *_redis.Client

func DbConn() {
	helper.Log("DbConn()")
	configs := util.Configs()
	dbDriver := configs.Find("DB_CONNECTION")
	dbUser := configs.Find("DB_USERNAME")
	dbPass := configs.Find("DB_PASSWORD")
	dbhost := configs.Find("DB_HOST") + ":" + configs.Find("DB_PORT")
	dbName := configs.Find("DB_DATABASE")

	dbcon, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbhost+")/"+dbName+"?parseTime=true")
	if err != nil {
		log.Println(err.Error())
	} else {
		print("DbConn ---- ok")
	}
	if err = dbcon.Ping(); err != nil {
		log.Println(err)
	}
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbcon,
	}), &gorm.Config{})

}

func GetDB() *gorm.DB {
	helper.Log("GetDB()")
	return db
}

//Init

//GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}
