package fog

type TConfig struct {
	RootURL string
}

func (this *TConfig) Create() *TConfig {
	return this
}
