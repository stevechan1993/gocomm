package websocket

import (
	"github.com/gorilla/websocket"
	"testing"
)

func TestNewMemoryConnmgr(t *testing.T){
	connmgr :=NewMemoryConnmgr()
	num:=100
	var listConn  []*WebsocketConnection
	for i:=0;i<num;i++{
		listConn = append(listConn,&WebsocketConnection{
			Uid:int64(i),
			AppId:1,
			Conn:&websocket.Conn{},
		})
	}
	for i:=range listConn{
		connmgr.Put(listConn[i],listConn[i])
	}
	if connmgr.Clients.Size() !=num && connmgr.Connections.Size()!=num{
		t.Fatal("size error :",connmgr.Clients.Size(),connmgr.Connections.Size())
	}
	for i:=range listConn{
		connmgr.Remove(listConn[i])
	}
	if connmgr.Clients.Size() !=0 && connmgr.Connections.Size()!=0{
		t.Fatal("size error :",connmgr.Clients.Size(),connmgr.Connections.Size())
	}
}

func BenchmarkNewMemoryConnmgr(b *testing.B) {
	connmgr :=NewMemoryConnmgr()
	var listConn  []*WebsocketConnection
	for i:=0;i<b.N;i++{
		listConn = append(listConn,&WebsocketConnection{
			Uid:int64(i),
			AppId:1,
			Conn:&websocket.Conn{},
		})
	}
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		connmgr.Put(listConn[i],listConn[i])
	}
	for i:=0;i<b.N;i++{
		connmgr.Remove(listConn[i])
	}
	if connmgr.Clients.Size() !=0 && connmgr.Connections.Size()!=0{
		b.Fatal("size error :",connmgr.Clients.Size(),connmgr.Connections.Size())
	}
}
