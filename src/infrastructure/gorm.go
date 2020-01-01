package infrastructure

import (
	"log"
	"time"

	"github.com/ezio1119/fishapp-profile/conf"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConn() *gorm.DB {
	mysqlConf := &mysql.Config{
		User:                 conf.C.Db.User,
		Passwd:               conf.C.Db.Pass,
		Net:                  conf.C.Db.Net,
		Addr:                 conf.C.Db.Host + ":" + conf.C.Db.Port,
		DBName:               conf.C.Db.Name,
		ParseTime:            conf.C.Db.Parsetime,
		Loc:                  time.Now().Location(),
		AllowNativePasswords: conf.C.Db.AllowNativePasswords,
	}

	dbConn, err := gorm.Open(conf.C.Db.Dbms, mysqlConf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	if conf.C.Sv.Debug {
		dbConn.LogMode(true)
	}
	return dbConn
}
