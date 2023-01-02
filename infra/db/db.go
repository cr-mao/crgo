package db

import (
	"crgo/infra/conf"
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"crgo/infra/log"
)

type ConnectionMap struct {
	mapping map[string]*gorm.DB
}

var connectionMap ConnectionMap
var once sync.Once

func InitDB() {
	once.Do(func() {
		vip := conf.GetViper()

		m := make(map[string]*gorm.DB)
		for bind, _ := range vip.GetStringMap("database") {
			instanceConf := vip.Sub("database." + bind)
			dsn := instanceConf.GetString("dsn")
			maxOpenConns := instanceConf.GetInt("max_open_conns")
			maxIdleConns := instanceConf.GetInt("mac_idle_conns")
			maxLifeTime := instanceConf.GetInt("max_life_seconds")
			debug := instanceConf.GetBool("debug")
			log.Debugf("preparing MySQL gorm.Connection -> %s @ %s", bind, dsn)

			db, err := gorm.Open("mysql", dsn)

			if err != nil {
				panic("parse MySQL DSN failed!" + err.Error())
			}
			if debug {
				db.LogMode(true)
			}
			//  表名 不适用复数形式
			db.SingularTable(true)
			db.DB().SetMaxIdleConns(maxIdleConns)
			db.DB().SetMaxOpenConns(maxOpenConns)
			db.DB().SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)
			m[bind] = db
		}
		connectionMap = ConnectionMap{mapping: m}
	})
}

func GetDb(bind string) *gorm.DB {
	if db, ok := connectionMap.mapping[bind]; ok {
		return db
	}
	panic(fmt.Sprintf("db bind %s get failed", bind))
}
