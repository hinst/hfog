package fog

import (
	"encoding/json"
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
