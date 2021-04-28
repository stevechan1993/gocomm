package mybeego

import (
	"encoding/json"
	"fmt"
	"gitlab.fjmaimaimai.com/mmm-go/gocomm/pkg/log"
	"strconv"

	"github.com/astaxie/beego"
	"gitlab.fjmaimaimai.com/mmm-go/gocomm/time"
)

// BaseController
type BaseController struct {
	beego.Controller
	Query       map[string]string
	JSONBody    map[string]interface{}
	ByteBody    []byte
	RequestHead *RequestHead
}

func assertCompleteImplement() {
	var _ beego.ControllerInterface = (*BaseController)(nil)
}

func (this *BaseController) Options() {
	this.AllowCross() //允许跨域
	this.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	this.ServeJSON()
}

func (this *BaseController) AllowCross() {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "*")
	//this.Ctx.WriteString("")
}

func (this *BaseController) Prepare() {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "*")
	if this.Ctx.Input.Method() == "OPTIONS" {
		this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//this.Ctx.WriteString("")
		return
	}

	this.Query = map[string]string{}
	input := this.Input()
	for k := range input {
		this.Query[k] = input.Get(k)
	}
	if this.Ctx.Input.RequestBody != nil {
		this.ByteBody = this.Ctx.Input.RequestBody[:]
		if len(this.ByteBody) < 1 {
			this.ByteBody = []byte("{}")
		}
		this.RequestHead = this.GetRequestHead()
		this.RequestHead.SetRequestId(fmt.Sprintf("%v.%v.%s", this.RequestHead.Uid, time.GetTimeByYyyymmddhhmmss(), this.Ctx.Request.URL))
		log.Debug(fmt.Sprintf("====>Recv data from uid(%d) client:\nHeadData: %s\nRequestId:%s BodyData: %s", this.RequestHead.Uid, this.Ctx.Request.Header, this.RequestHead.GetRequestId(), string(this.ByteBody)))
	}
	//key := SWITCH_INFO_KEY
	//str := ""
	//switchInfo := &TotalSwitchStr{}
	//if str, _ = redis.Get(key); str == "" {
	//	switchInfo.TotalSwitch = TOTAL_SWITCH_ON
	//	switchInfo.MessageBody = "正常运行"
	//	redis.Set(key, switchInfo, redis.INFINITE)
	//} else {
	//	json.Unmarshal([]byte(str), switchInfo)
	//}
	//if switchInfo.TotalSwitch == TOTAL_SWITCH_OFF {
	//	var msg *Message
	//	msg = NewMessage(3)
	//	msg.Errmsg = switchInfo.MessageBody
	//	log.Info(msg.Errmsg)
	//	this.Data["json"] = msg
	//	this.ServeJSON()
	//	return
	//}
}

func (this *BaseController) GetRequestHead() *RequestHead {
	reqHead := &RequestHead{}
	reqHead.Token = this.Ctx.Input.Header("token")
	reqHead.Version = this.Ctx.Input.Header("version")
	reqHead.Os = this.Ctx.Input.Header("os")
	reqHead.From = this.Ctx.Input.Header("from")
	reqHead.Screen = this.Ctx.Input.Header("screen")
	reqHead.Model = this.Ctx.Input.Header("model")
	reqHead.Channel = this.Ctx.Input.Header("channel")
	reqHead.Net = this.Ctx.Input.Header("net")
	reqHead.DeviceId = this.Ctx.Input.Header("deviceid")
	reqHead.Uid, _ = strconv.ParseInt(this.Ctx.Input.Header("uid"), 10, 64)
	reqHead.AppId, _ = strconv.Atoi(this.Ctx.Input.Header("appid"))
	reqHead.LoginIp = this.Ctx.Input.IP()
	reqHead.Jwt = this.Ctx.Input.Header("jwt")
	return reqHead
}

func (this *BaseController) Resp(msg *Message) {

	this.Data["json"] = msg
	this.ServeJSON()
}

func (this *BaseController) Finish() {
	if this.Ctx.Input.Method() == "OPTIONS" {
		return
	}
	strByte, _ := json.Marshal(this.Data["json"])
	length := len(strByte)
	if length > 5000 {
		log.Debug(fmt.Sprintf("<====Send to uid(%d) client: %d byte\nRequestId:%s RspBodyData: %s......", this.RequestHead.Uid, length, this.RequestHead.GetRequestId(), string(strByte[:5000])))
	} else {
		log.Debug(fmt.Sprintf("<====Send to uid(%d) client: %d byte\nRequestId:%s RspBodyData: %s", this.RequestHead.Uid, length, this.RequestHead.GetRequestId(), string(strByte)))
	}
}
