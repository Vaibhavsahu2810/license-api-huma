package models

type License struct {
	ShortName string `json:"shortname" doc:"Short name of the license" required:"true"`
	FullName  string `json:"fullname" doc:"Full name of the license" required:"true"`
	Text      string `json:"text" doc:"Text of the license" required:"true"`
}
