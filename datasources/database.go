package datasources

import (
	"fmt"
	"gcm-push/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pelletier/go-toml"
	"time"
	"unsafe"
)

var (
	DB          = New()
	MaxLifetime = 10
)

func New() *gorm.DB {
	driver := config.Conf.Get("database.driver").(string)
	configTree := config.Conf.Get(driver).(*toml.Tree)
	userName := configTree.Get("databaseUserName").(string)
	password := configTree.Get("databasePassword").(string)
	databaseName := configTree.Get("databaseName").(string)
	databaseHost := configTree.Get("databaseHost").(string)
	databasePort := configTree.Get("databasePort").(string)
	connect := userName + ":" + password + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"

	DB, err := gorm.Open(driver, connect)

	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}

	// use connection pool
	idleConns := configTree.Get("setMaxIdleConns")
	openConns := configTree.Get("setMaxOpenConns")
	DB.DB().SetMaxIdleConns(*(*int)(unsafe.Pointer(&idleConns)))
	DB.DB().SetMaxOpenConns(*(*int)(unsafe.Pointer(&openConns)))
	DB.DB().SetConnMaxLifetime(time.Second * time.Duration(MaxLifetime))
	return DB
}
