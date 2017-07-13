package fog

import (
	"encoding/json"
	"fmt"
	"hgo"
	"net/http"
	"strings"
)

type TWebUI struct {
	URL                string
	ApiURL             string
	DB                 *TDBMan
	AccessKey          string
	SecondaryAccessKey string
}

func (this *TWebUI) Create() *TWebUI {
	this.URL = "/FogBugzBackup"
	this.ApiURL = "/FogBugzBackupApi"
	return this
}

func (this *TWebUI) Start() {
	this.AddRequestHandler("/bugs", this.GetBugs)
	this.AddRequestHandler("/getBug", this.GetBug)
	this.AddRequestHandler("/getBugsFiltered", this.GetBugsFiltered)
	this.AddRequestHandler("/getAtt/", this.DownloadAttachment)
	this.InstallUiFileHandlers()
}

func (this *TWebUI) GetBugs(response http.ResponseWriter, request *http.Request) {
	var titles = this.DB.GetTitles()
	var bugs = TBugHeaderWebStruct{}.GetFromMap(titles)
	var data, marshalResult = json.Marshal(&bugs)
	WriteLogResult(marshalResult)
	if marshalResult == nil {
		response.Header().Set("Content-Type", "application/json")
		response.Write(data)
	}
}

func (this *TWebUI) GetBugsFiltered(response http.ResponseWriter, request *http.Request) {
	var filterString = request.URL.Query().Get("filter")
	var titles = this.DB.GetTitlesFiltered(
		TDBSearchFilter{
			Words:           strings.Split(filterString, " "),
			CommentsEnabled: request.URL.Query().Get("ce") == "y",
		},
	)
	var bugs = TBugHeaderWebStruct{}.GetFromRankedMap(titles)
	var data, marshalResult = json.Marshal(&bugs)
	WriteLogResult(marshalResult)
	if marshalResult == nil {
		response.Header().Set("Content-Type", "application/json")
		response.Write(data)
	}
}

func (this *TWebUI) AddRequestHandler(url string, f func(response http.ResponseWriter, request *http.Request)) {
	var wrappedFunc = func(response http.ResponseWriter, request *http.Request) {
		var incomingAccessKey = request.URL.Query().Get("AccessKey")
		if incomingAccessKey == this.AccessKey ||
			len(this.SecondaryAccessKey) > 0 && incomingAccessKey == this.SecondaryAccessKey {
			response.Header().Set("Access-Control-Allow-Origin", "*")
			f(response, request)
		} else {
			response.Write([]byte("ERROR: AccessKey mismatch"))
		}
	}
	http.HandleFunc(this.ApiURL+url, wrappedFunc)
}

func (this *TWebUI) InstallUiFileHandler(subDir string) {
	var directoryPath = hgo.AppDir + "/hfog-ui/build" + subDir
	var url = this.URL + subDir + "/"
	var fileDirectory = http.Dir(directoryPath)
	var fileServerHandler = http.FileServer(fileDirectory)
	fmt.Println(url + " -> " + directoryPath)
	http.HandleFunc(url,
		hgo.WrapFixJavaScriptContentType(
			http.StripPrefix(url, fileServerHandler),
		),
	)
}

func (this *TWebUI) InstallUiFileHandlers() {
	this.InstallUiFileHandler("")
	this.InstallUiFileHandler("/static/css")
	this.InstallUiFileHandler("/static/js")
	this.InstallUiFileHandler("/static/media")
}

func (this *TWebUI) GetBug(response http.ResponseWriter, request *http.Request) {
	var bug = this.DB.LoadBugData(request.URL.Query().Get("id"))
	if bug != nil {
		var wBug = bug.ToBugDataWeb()
		var data, marshalResult = json.Marshal(wBug)
		AssertResult(marshalResult)
		response.Header().Set("Content-Type", "application/json")
		response.Write(data)
	} else {
		response.Write([]byte("No such bug"))
	}
}

func (this *TWebUI) DownloadAttachment(response http.ResponseWriter, request *http.Request) {
	var key = request.URL.Query().Get("key")
	if len(key) > 0 {
		func() {
			var tx = this.DB.StartTx(false)
			defer tx.Commit()
			var op TDBAttachmentOp
			op.Tx = tx
			op.Key = key
			op.Read()
			if op.Data != nil {
				if len(op.Data) > 0 {
					if op.ImageType == "png" {
						response.Header().Set("Content-Type", "image/png")
					}
					if op.ImageType == "jpeg" {
						response.Header().Set("Content-Type", "image/jpeg")
					}
					if op.ImageType == "gif" {
						response.Header().Set("Content-Type", "image/gif")
					}
					response.Write(op.Data)
				} else {
					response.Header().Set("Content-Type", "text/plain")
					response.Write([]byte("No data\n" + key))
					response.Write([]byte("\nAllowed: " + hgo.BoolToStr(op.Allowed)))
					response.Write([]byte("\nFileName: '" + op.FileName + "'"))
				}
			} else {
				response.Header().Set("Content-Type", "text/plain")
				response.Write([]byte("Does not exist\n" + key))
			}
		}()
	}
}
