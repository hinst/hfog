package fog

import (
	"encoding/json"
	"fmt"
	"hgo"
	"net/http"
)

type TWebUI struct {
	URL string
	DB  *TDBMan
}

func (this *TWebUI) Create() *TWebUI {
	this.URL = "/FogBugzBackup"
	return this
}

func (this *TWebUI) Start() {
	this.AddRequestHandler("/bugs", this.GetBugs)
	this.AddRequestHandler("/getBug", this.GetBug)
	this.InstallUiFileHandlers()
}

func (this *TWebUI) GetBugs(response http.ResponseWriter, request *http.Request) {
	var bugs = make([]TBugHeaderWebStruct, 0)
	var titles = this.DB.GetTitles()
	for key, value := range titles {
		bugs = append(bugs,
			TBugHeaderWebStruct{
				Number: key,
				Title:  value,
			})
	}
	var data, marshalResult = json.Marshal(&bugs)
	WriteLogResult(marshalResult)
	if marshalResult == nil {
		response.Header().Set("Content-Type", "application/json")
		response.Write(data)
	}
}

func (this *TWebUI) AddRequestHandler(url string, f func(response http.ResponseWriter, request *http.Request)) {
	var wrappedFunc = func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		f(response, request)
	}
	http.HandleFunc(this.URL+url, wrappedFunc)
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
	}
}
