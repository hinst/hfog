package fog

import "net/http"

type TWebUI struct {
}

func (this *TWebUI) Create() *TWebUI {
	return this
}

func (this *TWebUI) Start() {

}

func (this *TWebUI) ReadBugList() {
	ReadBugsFromFile("data/bugs.xml")
}

func (this *TWebUI) GetBugs(response http.ResponseWriter, request *http.Request) {
	var bugs struct {
		Number int
		Title  string
	}
	Unuse(bugs)
}
