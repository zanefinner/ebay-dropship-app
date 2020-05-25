package accounts

//User is the format for users
type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	IsPremium bool   `json:"premium"`
}
