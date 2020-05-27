package accounts

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
	"github.com/martini-contrib/render"
)

var (
	err      error
	db       *gorm.DB
	username = "zane"
	password = "5245"
	dbname   = "dash"
)

//LoginForm is used for form binding
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//New makes a new account
func New() {
	//
}

//Update updates an account
func Update() {
	//
}

//Del deletes an account
func Del() {
	//
}

//Match checks user input with current data
func Match(login LoginForm, ren render.Render) {
	fmt.Println("POST: /login")
	db, err = gorm.Open("mysql", username+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	var usr User
	db.AutoMigrate(&User{})
	res := db.First(&usr, "username = ? and password = ?", login.Username, login.Password)
	if !res.NewRecord(&usr) {
		//make session
	} else {
		//try again, nerd!
	}
	defer db.Close()
	//can use login.Username || login.Password
	if err != nil {
		fmt.Println(err)
		return
	}

	//usr := []User{}
}

//PresentLoginForm shows a html login template to send a post request
func PresentLoginForm(r render.Render) {
	r.HTML(200, "forms/login", nil)
}
