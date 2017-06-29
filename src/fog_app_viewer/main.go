package main

import (
	"fmt"
	"fog"
	"runtime/debug"
)

func main() {
	fmt.Println("STARTING...")
	debug.SetGCPercent(10)
	var app = (&fog.TWebApp{}).Create()
	app.Run()
	fmt.Println("EXITING...")
}
