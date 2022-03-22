package userdto

type Register struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResRegister struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}
