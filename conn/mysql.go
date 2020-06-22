package conn

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hulujia/config"
	"hulujia/database"
	"hulujia/util/log"
	"time"
)

var (
	db *gorm.DB
)

func SetupMysql()  {
	var err error
	host := config.Database.Host
	user := config.Database.User
	password := config.Database.Password
	name := config.Database.Name
	charset := config.Database.CharSet

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, name, charset)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect mysql %s", err.Error()))
	} else {
		log.Info("Connect to MySQL successfully, database: %s.", name)
		db.DB().SetMaxIdleConns(config.Database.Max)
		db.DB().SetMaxOpenConns(config.Database.Min)
		db.DB().SetConnMaxLifetime(time.Minute)
		if gin.Mode() != gin.ReleaseMode {
			db.LogMode(true)
		}
	}

	db.SingularTable(true) //禁用表名复数

	// database
	if err := database.Migrate(db); err != nil {
		log.Error("auto database tables failed")
	}

	// Seeder
	database.Seeder(db)

	// foreign
	database.Roreign(db)
}

// Shutdown - close database connection
func Shutdown() error {
	log.Info("Closing database's connections")
	return db.Close()
}

// GetDb - get a database connection
func DB() *gorm.DB {
	return db
}

// 事务环绕
func Tx(txFunc func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	err = txFunc(tx)
	return err
}