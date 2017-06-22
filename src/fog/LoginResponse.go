package fog

type TLoginResponse struct {
	Response struct {
		Error string `xml:"error"`
		Token string `xml:"token"`
	} `xml:"response"`
}
