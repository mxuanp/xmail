package model


type User struct {
	Password string `json:password`
	Email    string `json:email`
	MailHost string `json:host`
	Port     string `json:port`
}
