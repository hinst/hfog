package fog

type TBugDataWeb struct {
	Id     string
	Title  string
	Events []TBugDataWebEvent
}

type TBugDataWebEvent struct {
	Moment      string
	Verb        string
	Description string
	Person      string
	Text        string
	HTML        string
}
