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
	app.AttachmentFilter = fog.ImageFileNameSuffixes
	app.AttachmentsModeEnabled = true
	app.AttachmentTestModeEnabled = false
	app.EnumAttachmentsModeEnabled = false
	app.ImageCompressionTestModeEnabled = false
	app.RunAllowImagesMode = false
	app.Run()
	fmt.Println("EXITING...")
}
