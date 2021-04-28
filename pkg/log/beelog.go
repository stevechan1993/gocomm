package log

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/stevechan1993/gocomm/config"
	"path/filepath"
	"strconv"
)

type beegoLog struct {
	log *logs.BeeLogger
}

func newbeelog(conf config.Logger)Log{
	filename := `{"filename":"` + filepath.ToSlash(conf.Filename) + `"}`

	l :=&beegoLog{
		log:logs.GetBeeLogger(),
	}
	l.log.SetLogger(logs.AdapterFile,filename)
	ilv,err :=strconv.Atoi(conf.Level)
	if err!=nil{
		ilv = logs.LevelDebug
	}
	l.log.SetLevel(ilv)
	l.log.EnableFuncCallDepth(true)
	l.log.SetLogFuncCallDepth(6)
	return l
}

func(this *beegoLog)Debug(args ...interface{}){
	//this.log.Debug(args...)
	logs.Debug(nil,args...)
}

func(this *beegoLog)Info(args ...interface{}){
	logs.Info(nil,args...)
}

func(this *beegoLog)Warn(args ...interface{}){
	logs.Warn(nil,args...)
}

func(this *beegoLog)Error(args ...interface{}){
	logs.Error(nil,args...)
}

func(this *beegoLog)Panic(args ...interface{}){
	logs.Error(nil,args...)
}

func(this *beegoLog)Fatal(args ...interface{}){
	logs.Error(nil,args...)
}
