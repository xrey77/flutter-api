package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Users struct {
	ID          int64
	Username    string `binding:"required,min=5,max=30"`
	Password    string `binding:"required,min=100,max=100"`
	Userpicture string `binding:"required,min=10,max=10"`
	Isactivated int64
	Otp         int64
	Email       string `binding:"required,min=100,max=100"`
	Role        string `binding:"required,min=10,max=10"`
	// Token       string `binding:"required,min=200,max=200"`
}

type User struct {
	ID          int64     `json:"id"`
	Lastname    string    `json:"lastname"`
	Firstname   string    `json:"firstname"`
	Email       string    `json:"email"`
	Mobile      string    `json:"mobile"`
	Username    string    `json:"username"`
	Userpicture string    `json:"userpicture"`
	Role        string    `json:"role"`
	Password    string    `json:"password"`
	Isactivated int64     `json:"isactivated"`
	Mailtoken   string    `json:"mailtoken"`
	Otp         int64     `json:"otp"`
	Secretkey   string    `json:"secretkey"`
	Qrcode      string    `json:"qrcode"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}

type Userz struct {
	ID          int64
	Lastname    string `binding:"required,min=30,max=30"`
	Firstname   string `binding:"required,min=30,max=30"`
	Email       string `binding:"required,min=150,max=150"`
	Mobile      string `binding:"required,min=30,max=30"`
	Username    string `binding:"required,min=30,max=30"`
	Userpicture string `binding:"required,min=100,max=100"`
	Role        string `binding:"required,min=10,max=10"`
	Password    string `binding:"required,min=420,max=420"`
	Qrcode      string `binding:"required,min=150,max=150"`
	UpdatedAt   time.Time
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type TempUsers struct {
	ID          int64  `json:"id"`
	Lastname    string `json:"lastname"`
	Firstname   string `json:"firstname"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Userpicture string `json:"userpicture"`
	Role        string `json:"role"`
	Otp         int64  `json:"otp"`
	Qrcode      string `json:"qrcode"`
	Secretkey   string `json:"secretkey"`
	Isactivated int64  `json:"isactivated"`
	Mailtoken   string `json:"mailtoken"`
}

type Tempusers struct {
	ID          int64
	Lastname    string `binding:"required,min=30,max=30"`
	Firstname   string `binding:"required,min=30,max=30"`
	Email       string `binding:"required,min=100,max=100"`
	Mobile      string `binding:"required,min=20,max=20"`
	Username    string `binding:"required,min=5,max=30"`
	Userpicture string `binding:"required,min=10,max=10"`
	Role        string `binding:"required,min=10,max=10"`
	Password    string `binding:"required,min=100,max=100"`
	Isactivated int64
	Mailtoken   string `binding:"required,min=50,max=50"`
	Otp         int64
	Secretkey   string    `binding:"required,min=250,max=250"`
	Qrcode      string    `binding:"required,min=350,max=350"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}

type Userlogin struct {
	ID          int
	Username    string `binding:"required,min=5,max=30"`
	Password    string `binding:"required,min=100,max=100"`
	Userpicture string `binding:"required,min=10,max=10"`
	Isactivated int64
	Otp         int64
	Email       string `binding:"required,min=100,max=100"`
	Role        string `binding:"required,min=10,max=10"`
	Token       string `binding:"required,min=200,max=200"`
	// Email       string `json:"email"`
	// Token       string `json:"token"`
	// Role        string `json:"role"`
	// Otp         int64  `json:"otp"`
	// IsActivated int64  `json:"isactivated"`
}

// Token       string `json:"token";sql:"-"`

type TmpLogin struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Token    string `json:"token"`
}

type Mfa struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Otp       int64  `json:"otp"`
	Secretkey string `json:"secretkey"`
	Qrcode    string `json:"qrcode"`
}

type Mfaverify struct {
	Username    string `json:"username"`
	Userpicture string `json:"userpicture"`
	Secretkey   string `json:"secretkey"`
}

type Qcode struct {
	ID     uint   `json:"id"`
	Qrcode string `json:"qrcode"`
}
