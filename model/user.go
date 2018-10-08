package model

type User struct {
	Id string `json:"id, primarykey"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func NewUser(name string, pass string) User {
	return User {
		Name: name,
		Pass: pass,
	}
}