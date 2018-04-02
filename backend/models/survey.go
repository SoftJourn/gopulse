package models

type Survey struct {
	ID          int64
	Title       string `orm:"size(128);"`
	Description string `orm:"size(128);"`
	Schedule    string `orm:"size(128);"`
	Questions   string `orm:"size(128);"`
	Recipients  string `orm:"size(128);"`
}
