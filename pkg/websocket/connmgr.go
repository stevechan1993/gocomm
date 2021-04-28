package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"gitlab.fjmaimaimai.com/mmm-go/gocomm/pkg/mybeego"
	"reflect"
	"sync"
)

var (
	ErrorNotFound =fmt.Errorf("conn not exist")
)

var DefaultConnmgrs Connmgrs
//初始化websocket链路管理器数量,通过uid取模放入不同的管理器
func InitWebsocketConnmgrs(mgrsSize int){
	connmgrs :=make(map[int]IConnmgr,mgrsSize)
	for i:=0;i<mgrsSize;i++{
		connmgrs[i] = NewMemoryConnmgr()
	}
	DefaultConnmgrs = Connmgrs(connmgrs)
}

type IConnmgr interface {
	Put(key,value interface{})(bool)
	Remove(key interface{}) error
	Get(key interface{})(value interface{},err error)
}
//连接管理器
type Connmgrs map[int]IConnmgr
//将连接从指定连接管理器中移除
func (m Connmgrs)Remove(connmgrId int,key interface{})(err error){
	//删除特定链接管理的连接
	if mgr,ok:= m[connmgrId];ok{
		err =mgr.Remove(key)
		if err!=nil{
			return err
		}
	}
	return
}
//将连接装载到指定 连接管理器
func (m Connmgrs)Put(connmgrId int,key,value interface{})(result bool){
	result =false
	if mgr,ok:= m[connmgrId];ok{
		result =mgr.Put(key,value)
		if !result{
			return
		}
	}
	return
}


type MemoryConnmgr struct {
	mutex sync.RWMutex
	Connections *JMap //conn
	Clients *JMap  // key=uid(int64) value(*WebsocketConnection)
	//rooms  //房间1
}

func NewMemoryConnmgr()*MemoryConnmgr{
	keyType := reflect.TypeOf(&websocket.Conn{})
	valueType := reflect.TypeOf(&WebsocketConnection{})
	return &MemoryConnmgr{
		Connections:NewJMap(keyType, valueType),
		Clients:NewJMap(reflect.TypeOf("1:1"), valueType),
	}
}

func(m *MemoryConnmgr)Put(key,value interface{})(result bool){
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if c,ok :=value.(*WebsocketConnection);ok{
		idKey := fmt.Sprintf("%d:%d", c.Uid, c.AppId)
		return m.Connections.Put(c.Conn,c) && m.Clients.Put(idKey,c)
	}
	return false
}

func(m *MemoryConnmgr)Get(key interface{})(value interface{},err error){
	var ok bool
	switch reflect.TypeOf(key).Kind() {
	case reflect.String:
		if value,ok = m.Clients.Get(key);!ok{
			err = ErrorNotFound
			return
		}
	case reflect.Struct:
		if value,ok = m.Connections.Get(key);!ok{
			err = ErrorNotFound
			return
		}
	default:
		err = ErrorNotFound
	}
	return
}

func(m *MemoryConnmgr)Remove(key interface{})(err error){
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if c,ok :=key.(*WebsocketConnection);ok{
		key := fmt.Sprintf("%d:%d", c.Uid, c.AppId)
		m.Connections.Remove(c.Conn)
		m.Clients.Remove(key)
	}
	return
}

//发送数据
func SendDataByConnmgr(uid int64, appId int, sendMsg interface{}) bool {
	if sendMsg == nil || uid < 1 || appId < 1 {
		return false
	}
	var mgrId int =int(uid % int64(len(DefaultConnmgrs)))
	connmgr,ok :=(DefaultConnmgrs)[mgrId]
	if !ok{
		return false
	}
	msg := &mybeego.Message{
		Errno:  0,
		Errmsg: mybeego.NewMessage(0).Errmsg,
		Data:   sendMsg,
	}
	msgByte, err := json.Marshal(msg)
	if err != nil {
		beego.Error(err)
		return false
	}
	key := fmt.Sprintf("%d:%d", uid, appId)
	if connI,err := connmgr.Get(key); err==nil {
		if conn, ok := connI.(*WebsocketConnection); ok {
			conn.Send(string(msgByte))
			return true
		}
	}
	return false
}
