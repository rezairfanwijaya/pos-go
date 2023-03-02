package user

type InputUserLogin struct {
	Username string `json:"username" biding:"required"`
	Password string `json:"password" biding:"required,min=5,max=5"`
}
