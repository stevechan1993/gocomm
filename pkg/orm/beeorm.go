package orm

import (
	"github.com/astaxie/beego/orm"
	"github.com/stevechan1993/gocomm/config"
	"github.com/stevechan1993/gocomm/pkg/log"
)

func NewBeeormEngine(conf config.Mysql) {
	aliasName := "default"
	if len(conf.AliasName) > 0 {
		aliasName = conf.AliasName
	}
	err := orm.RegisterDataBase(aliasName, "mysql", conf.DataSource)
	if err != nil {
		log.Error(err)
	} else {
		//log.Debug("open db address:",conf.DataSource)
	}
	orm.SetMaxIdleConns(aliasName, conf.MaxIdle)
	orm.SetMaxOpenConns(aliasName, conf.MaxOpen)
	//orm.DefaultTimeLoc = time.Local
	//orm.Debug = true
}
