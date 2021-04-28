package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stevechan1993/gocomm/pkg/log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"
)
//go test -v example_test.go -test.run TestWebSocketClient
func WebSocketClient(t *testing.T){
	var clientlist []*websocket.Conn
	var num = 1000

	doRead :=func(c *websocket.Conn,done chan struct{},key string){
		for {
			defer close(done)
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Info(key," read:", err)
				return
			}
			log.Info(key," recv: ", string(message))
		}
	}

	doWrite :=func(c *websocket.Conn,done chan struct{},key string){
		ticker := time.NewTicker(time.Second*60)
		defer ticker.Stop()
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
				if err != nil {
					log.Info(key," write:", err)
					return
				}
			case <-interrupt:
				log.Info(key," interrupt")

				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Info(key," write close:", err)
					return
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
				return
			}
		}
	}

	for i:=1;i<=num;i++{
		u :=url.URL{Scheme:"ws",Host:"127.0.0.1:8080",Path:"/upgrage"}
		requestHeader :=http.Header{}
		requestHeader.Add("uid",fmt.Sprintf("%d",i))
		requestHeader.Add("appid",fmt.Sprintf("%d",2))
		conn,_,err:=websocket.DefaultDialer.Dial(u.String(),requestHeader)
		if err!=nil{
			log.Fatal(err)
		}
		do:= make(chan struct{})
		key :=fmt.Sprintf("%v:%v",i,2)
		go doRead(conn,do,key)
		go doWrite(conn,do,key)
		clientlist = append(clientlist, conn)
	}
	time.Sleep(time.Second*3600)
}