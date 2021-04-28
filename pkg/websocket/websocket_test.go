package websocket

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stevechan1993/gocomm/pkg/log"
	"github.com/stevechan1993/gocomm/pkg/mybeego"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func Test_RunWebSocket(t *testing.T) {
	//InitWebsocketConnmgrs(10)
	//http.HandleFunc("/upgrade",join)
	//http.HandleFunc("/",home)
	//
	//go TimerSendData()
	//go TimerStatus()
	//log.Fatal(http.ListenAndServe(":8080",nil))
}

func TimerSendData(){
	t :=time.NewTicker(10*time.Second)
	ch :=make(chan int,1)
	for {
		select {
			case <-t.C:
				uid :=rand.Int63n(500)
				SendDataByConnmgr(uid,2,time.Now())
		}
	}
	<-ch
}

func TimerStatus(){
	t :=time.NewTicker(10*time.Second)
	ch :=make(chan int,1)
	for {
		select {
		case <-t.C:
			buf :=bytes.NewBuffer(nil)
			buf.WriteString("")
			for i:=0;i<len(DefaultConnmgrs);i++{
				if v ,ok :=DefaultConnmgrs[i].(*MemoryConnmgr);ok{
					buf.WriteString(fmt.Sprintf("id:%d clients:%d connect:%d\n",i,v.Clients.Size(),v.Connections.Size()))
				}
			}
			log.Info(buf.String())
		}
	}
	<-ch
}

var upgrader = websocket.Upgrader{}

func join(w http.ResponseWriter, r *http.Request) {
	requestHead := &mybeego.RequestHead{}
	requestHead.Uid, _ = strconv.ParseInt(r.Header.Get("uid"), 10, 64)
	requestHead.AppId, _ = strconv.Atoi(r.Header.Get("appid"))
	requestHead.Token = r.Header.Get("token")
	if !validToken(requestHead.Token) {
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	wsConn := NewWebsocketConnection(conn, requestHead, onReceive)
	wsConn.Serve()
}

func onReceive(data []byte) *mybeego.Message {
	return mybeego.NewMessage(0)
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/join")
}

func validToken(token string) bool {
	return true
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
