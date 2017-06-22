package fog

type TConfig struct {
	RootURL  string
	Email    string
	Password string
}

func (this *TConfig) Create() *TConfig {
	return this
}
