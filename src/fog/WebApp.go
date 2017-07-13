package fog

import (
	"encoding/json"
	"fmt"
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
	if true {
		fmt.Println(hgo.MakeRandomString(10))
	}
	this.Config.Address = ":9000" // default
	return this
}

func (this *TWebApp) Run() {
	this.LoadConfig()
	this.DB = (&TDBMan{}).Create()
	this.DB.ReadOnly = false
	this.DB.Start()
	this.WebUI = (&TWebUI{}).Create()
	this.WebUI.DB = this.DB
	this.WebUI.AccessKey = this.Config.AccessKey
	this.WebUI.SecondaryAccessKey = this.Config.SecondaryAccessKey
	this.WebUI.Start()
	var server = hgo.StartHttpServer(this.Config.Address)
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
