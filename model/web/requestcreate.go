package web

type RequestCreate struct {
	Firstname string `validate:"required" json:"firstname"`
	Lastname  string `validate:"required" json:"lastname"`
	Username  string `validate:"required" json:"username"`
	Password  string `validate:"required" json:"password"`
}
