package fog

type TApp struct {
	RootURL string
}

func (this *TApp) Create() *TApp {
	return this
}

func (this *TApp) Run() {

}