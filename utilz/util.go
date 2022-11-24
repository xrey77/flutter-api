package utilz

import (
	"encoding/json"
	"log"
	"net/http"
	"src/flutter-api/config"
	"src/flutter-api/models"

	"golang.org/x/crypto/bcrypt"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		Respond(w, Message(false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}

func ValidateUserName(usrname string) int {
	db := config.Connect()
	defer db.Close()
	Tmpusers := new(models.Users)
	if err22 := db.Model(Tmpusers).Where("username = ?", usrname).Select(); err22 != nil {
		log.Print("Username not found.", err22.Error())
		return 0
	}
	return 1
}

func IsActivated(usname string) int64 {
	db := config.Connect()
	defer db.Close()
	tmpUser1 := new(models.Users)
	if err23 := db.Model(tmpUser1).Where("username = ?", usname).Select(); err23 != nil {
		return 0
	}
	return tmpUser1.Isactivated
}

func GetNextval() int64 {
	db := config.Connect()
	defer db.Close()
	var users models.Users
	err := db.Model(&users).Last()
	log.Print("Count : ", users.ID)
	if err != nil {
		log.Print(err)
		return 0
	} else {
		return int64(users.ID) + 1
	}
}

func ValidateEmail(mail string) int {
	db := config.Connect()
	defer db.Close()
	var tmpUser = new(models.Users)
	if err2 := db.Model(tmpUser).Where("email = ?", mail).Select(); err2 != nil {
		return 0
	}
	return 1
}

func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err24 := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err24 != nil {
		return false
	} else {
		return true
	}
}

func EncryptPassword(pwd string) []byte {
	return []byte(pwd)
}

func GetPasswd(usrname string) string {
	db := config.Connect()
	defer db.Close()
	var tmpUser = new(models.Users)
	if err2 := db.Model(tmpUser).Where("username = ?", usrname).Select(); err2 != nil {
		return ""
	} else {
		return tmpUser.Password
	}
}

func GetPIC(usrname string) string {
	db := config.Connect()
	defer db.Close()
	var tmpUser = new(models.Users)
	if err2 := db.Model(tmpUser).Where("username = ?", usrname).Select(); err2 != nil {
		return ""
	} else {
		return tmpUser.Userpicture
	}
}

func GetUID(usname string) string {
	db := config.Connect()
	defer db.Close()
	var tmpUser1 = new(models.Users)
	if err2 := db.Model(tmpUser1).Where("username = ?", usname).Select(); err2 != nil {
		return ""
	} else {
		return string(rune(tmpUser1.ID))
	}
}

func GetOTP(usname string) int64 {
	db := config.Connect()
	defer db.Close()
	var tmpUser2 = new(models.Users)
	if err2 := db.Model(tmpUser2).Where("username = ?", usname).Select(); err2 != nil {
		return 0
	} else {
		return tmpUser2.Otp
	}
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
