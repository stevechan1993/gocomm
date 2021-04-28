package idgen

import (
	"testing"

	"github.com/stevechan1993/gocomm/pkg/log"
)

func Test_Next(t *testing.T){
	m :=make(map[int64]int64)
	num :=1000
	for i:=0;i<num;i++{
		id :=Next()
		if _,ok:=m[id];ok{
			log.Fatal("exists id:",id,len(m))
		}
	}
}
