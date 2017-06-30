package fog

import (
	"hgo"
	"sync"
	"time"
)

type TWebApp struct {
	WebUI  *TWebUI
	DB     *TDBMan
	Holder sync.WaitGroup
}

func (this *TWebApp) Create() *TWebApp {
	return this
}

func (this *TWebApp) Run() {
	WriteLog("Starting...")
	this.DB = (&TDBMan{}).Create()
	this.DB.FilePath = hgo.AppDir + "/data/db-sh.bolt"
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
