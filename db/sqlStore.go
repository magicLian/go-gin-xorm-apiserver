package db

import (
	"errors"
	"fmt"
	"gin-test-go/configs"
	"gin-test-go/models"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
)

var x *xorm.Engine

func InitDB(c configs.Conf) {
	pgInfo := "user=" + c.Db.Postgres.Username +
		" password=" + c.Db.Postgres.Password +
		" dbname=" + c.Db.Postgres.DBName +
		" host=" + c.Db.Postgres.Host +
		" port=" + c.Db.Postgres.Port +
		" sslmode=disable"
	fmt.Println("pgInfo = {}", pgInfo)

	err := errors.New("")
	x, err = xorm.NewEngine("postgres", pgInfo)
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}
	if err = x.Sync2(new(models.User)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

func GetDbInstance() *xorm.Engine {
	return x
}
