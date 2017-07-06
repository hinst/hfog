package main

import (
	"fmt"
	"fog"
	"hgo"
)

func main() {
	fmt.Println("STARTING...")
	var app = (&fog.TApp{}).Create()
	var afDoc = []string{".doc", ".docx"}
	hgo.Unuse(afDoc)
	var afImg = []string{".png", ".gif", ".jpg"}
	hgo.Unuse(afImg)
	app.AttachmentFilter = afImg
	app.AttachmentsModeEnabled = false
	app.AttachmentTestModeEnabled = false
	app.EnumAttachmentsModeEnabled = true
	app.Run()
	fmt.Println("EXITING...")
}
