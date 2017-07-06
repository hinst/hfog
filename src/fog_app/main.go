package main

import (
	"fmt"
	"fog"
)

func main() {
	fmt.Println("STARTING...")
	var app = (&fog.TApp{}).Create()
	app.AttachmentsModeEnabled = true
	app.AttachmentTestModeEnabled = false
	app.EnumAttachmentsModeEnabled = false
	app.Run()
	fmt.Println("EXITING...")
}
