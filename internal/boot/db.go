package boot

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"go-project-initial/configs"
	"sync"
)

func init() {
	new(sync.Once).Do(initDb)
}

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		initDb()
	}
	return db
}

func initDb() {
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		configs.Conf.Database.Username,
		configs.Conf.Database.Password,
		configs.Conf.Database.Host,
		configs.Conf.Database.Port,
		configs.Conf.Database.DBName,
		configs.Conf.Database.Charset)
	log.Printf("db connect, source: %s\n", source)
	var err error
	db, err = gorm.Open("mysql", source)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}
	// Enable Logger, show detailed log
	db.LogMode(true)
	//db.SetLogger(gorm.Logger{revel.TRACE})
	db.SetLogger(NewLogger("db"))
}

func DBClose() {
	if db != nil {
		log.Println("db close")
		_ = db.Close()
		db = nil
	}
}

//使用相同事务
func DBTransaction(tx *gorm.DB, fc func(tx *gorm.DB) error) (err error) {
	if tx == nil {
		err = DB().Transaction(fc)
	} else {
		err = fc(tx)
	}
	return
}
