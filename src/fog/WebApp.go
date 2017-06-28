package fog

import (
	"hgo"
	"sync"
	"time"
)

type TWebApp struct {
	WebUI  *TWebUI
	Holder *sync.WaitGroup
}

func (this *TWebApp) Create() *TWebApp {
	return this
}

func (this *TWebApp) Run() {
	WriteLog("Starting...")
	this.WebUI = (&TWebUI{}).Create()
	this.WebUI.Start()
	var server = hgo.StartHttpServer(":9000")
	this.Holder.Add(1)
	hgo.InstallShutdownReceiver(this.ReceiveShutdownSignal)
	this.Holder.Wait()
	var stopServerHttpResult = hgo.StopHttpServer(server, 1*time.Second)
	WriteLogResult(stopServerHttpResult)
}

func (this *TWebApp) ReceiveShutdownSignal() {
	this.Holder.Done()
}
