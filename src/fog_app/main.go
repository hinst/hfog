package main

import (
	"fmt"
	"fog"
)

func main() {
	fmt.Println("STARTING...")
	var app = (&fog.TApp{}).Create()
	app.AttachmentTestModeEnabled = true
	app.Run()
	fmt.Println("EXITING...")
}
