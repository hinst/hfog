package main

import (
	"fmt"
	"fog"
	"hgo"
	"runtime/debug"
)

func main() {
	fmt.Println("STARTING...")
	debug.SetGCPercent(10)
	var app = (&fog.TApp{}).Create()
	var afDoc = []string{".doc", ".docx"}
	hgo.Unuse(afDoc)
	app.AttachmentFilter = afDoc
	app.AttachmentsModeEnabled = false
	app.AttachmentTestModeEnabled = false
	app.EnumAttachmentsModeEnabled = true
	app.ImageCompressionTestModeEnabled = false
	app.RunAllowImagesMode = false
	app.Run()
	fmt.Println("EXITING...")
}
