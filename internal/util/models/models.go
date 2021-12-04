package models

import (
	"GoServerDemo/internal/util/settings"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := settings.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	// db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8md4",
	// 	user,
	// 	password,
	// 	host,
	// ))

	if err != nil {
		log.Fatalln(err)
	}

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func closeDB() {
	defer db.Close()
}

func GetMdList(pageNum int, pageSize int, maps interface{}) (files []Article) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&files)

	return
}

func GetMdNum(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}
