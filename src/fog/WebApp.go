package fog

import (
	"encoding/json"
	"hgo"
	"io/ioutil"
	"sync"
	"time"
)

type TWebApp struct {
	WebUI  *TWebUI
	DB     *TDBMan
	Holder sync.WaitGroup
	Config TWebAppConfig
}

func (this *TWebApp) Create() *TWebApp {
	println(hgo.MakeRandomString(10))
	return this
}

func (this *TWebApp) Run() {
	this.LoadConfig()
	this.DB = (&TDBMan{}).Create()
	this.DB.Start()
	this.WebUI = (&TWebUI{}).Create()
	this.WebUI.DB = this.DB
	this.WebUI.Start()
	var server = hgo.StartHttpServer(":9000")
	this.Holder.Add(1)
	hgo.InstallShutdownReceiver(this.ReceiveShutdownSignal)
	this.Holder.Wait()
	var stopServerHttpResult = hgo.StopHttpServer(server, 1*time.Second)
	WriteLogResult(stopServerHttpResult)
	this.DB.Stop()
}

func (this *TWebApp) ReceiveShutdownSignal() {
	this.Holder.Done()
}

func (this *TWebApp) LoadConfig() {
	var data, readFileResult = ioutil.ReadFile(hgo.AppDir + "/fog_app_viewer.json")
	AssertResult(readFileResult)
	var unmarshalResult = json.Unmarshal(data, &this.Config)
	AssertResult(unmarshalResult)
}
