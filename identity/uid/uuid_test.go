package uid

import (
	"fmt"
	"strings"
	"testing"
)

func TestUID(t *testing.T){
	 uid :=NewV1()
	 //t.Fatal(uid)
	 fmt.Println(uid)
	udata,err := uid.MarshalBinary()
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("MarshalBinary:",udata)
	fmt.Println("uuid version:",uid.Version())

	uidStr36 :=uid.String()
	uidStr32 :=uid.StringNoDash()
	if strings.Replace(uidStr36,"-","",-1) != uidStr32{
		t.Fatal("no equal",uidStr36,uidStr32)
	}
	t.Log(uidStr36,uidStr32)
}

func Test_StringNoDash(t *testing.T){
	for i:=0;i<100;i++{
		uid :=NewV1()
		uidStr36 :=uid.String()
		uidStr32 :=uid.StringNoDash()
		if strings.Replace(uidStr36,"-","",-1) != uidStr32{
			t.Fatal("no equal",uidStr36,uidStr32)
		}
	}
}

func Test_NewV1(t *testing.T){
	num :=10000
	mUid :=make(map[string]int,num)
	for i:=0;i<num;i++{
		uid :=NewV1()
		uidStr36 :=uid.String()
		if _,ok:=mUid[uidStr36];ok{
			t.Fatal("repeat uid",uidStr36)
		}else{
			mUid[uidStr36]=0
		}
	}
	if len(mUid)!=num{
		t.Fatal("map num error")
	}
}
