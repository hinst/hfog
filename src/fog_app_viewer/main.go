package main

import (
	"fmt"
	"fog"
	"hgo"
	_ "net/http/pprof"
	"runtime/debug"
	"time"
)

func main() {
	fmt.Println("STARTING...")
	debug.SetGCPercent(10)
	hgo.EnablePeriodicReleaseMemory(time.Minute)
	var app = (&fog.TWebApp{}).Create()
	app.Run()
	fmt.Println("EXITING...")
}

// Bug with PULL-LIST.JPG attachment file: 128379
