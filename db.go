package midware

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Driver  string
	DSN     string
	Prefix  string
	ShowSQL bool
	MaxOpen int
	MaxIdle int
}

//GetDB used to get a db resource
func GetDB(db *Database) (*xorm.Engine, error) {
	if db == nil {
		db = &Database{
			Driver:  DefaultDBDriver,
			DSN:     DefaultDBDSN,
			Prefix:  DefaultDBPrefix,
			ShowSQL: DefaultDBShowSQL,
			MaxOpen: DefaultDBMaxOpen,
			MaxIdle: DefaultDBMaxIdle,
		}
	}
	engine, err := xorm.NewEngine(db.Driver, db.DSN)
	if err != nil {
		fmt.Println("db engine error!")
		return nil, err
	}
	err = engine.Ping()
	if err != nil {
		fmt.Println("db engine timeout error!")
		return nil, err
	}

	engine.ShowSQL(db.ShowSQL)
	//连接池设置
	engine.SetMaxOpenConns(db.MaxOpen)
	engine.SetMaxIdleConns(db.MaxIdle)
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, db.Prefix))
	return engine, nil
}
