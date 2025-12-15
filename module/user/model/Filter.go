package usermodel

type Filter struct {
	Id      int    `json:"id,omitempty" form:"id"`
	Name    string `json:"name,omitempty" form:"name"`
	Email   string `json:"email,omitempty" form:"email"`
	Address string `json:"address,omitempty" form:"address"`
}
