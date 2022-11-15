package utilz

import (
	"encoding/json"
	"log"
	"net/http"
	"src/flutter-api/config"
	"src/flutter-api/models"
	"strconv"

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

func ValidateUserName(usrname string) int64 {
	db := config.Connect()
	defer db.Close()
	var tmpUser models.TempUsers
	sqlUname := `SELECT username FROM USERS WHERE username=$1;`
	userNamerow := db.QueryRow(sqlUname, usrname)
	err2 := userNamerow.Scan(&tmpUser.Username)
	if err2 != nil {
		log.Print(err2.Error())
		return 0
	} else {
		return 1
	}
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
	var tmpUser models.TempUsers
	sqlFindUsername := `SELECT password FROM USERS WHERE username=$1;`
	rowUsername := db.QueryRow(sqlFindUsername, usrname)
	err2 := rowUsername.Scan(&tmpUser.Password)
	if err2 != nil {
		return ""
	} else {
		return tmpUser.Password
	}
}

func GetPIC(usrname string) string {
	db := config.Connect()
	defer db.Close()
	var tmpUser models.TempUsers
	sqlFindUsername := `SELECT userpicture FROM USERS WHERE username=$1;`
	rowUsername := db.QueryRow(sqlFindUsername, usrname)
	err2 := rowUsername.Scan(&tmpUser.Userpicture)
	if err2 != nil {
		return ""
	} else {
		return tmpUser.Userpicture
	}
}

func IsActivated(usname string) int64 {
	db := config.Connect()
	defer db.Close()
	var tmpUser1 models.TempUsers
	sqlFindUname := `SELECT isactivated FROM USERS WHERE username=$1;`
	rowUname := db.QueryRow(sqlFindUname, usname)
	err23 := rowUname.Scan(&tmpUser1.Isactivated)
	if err23 != nil {
		return 0
	} else {
		return tmpUser1.Isactivated
	}
}

func GetUID(usname string) string {
	db := config.Connect()
	defer db.Close()
	var tmpUser1 models.TempUsers
	sqlFindUname := `SELECT id FROM USERS WHERE username=$1;`
	rowUname := db.QueryRow(sqlFindUname, usname)
	err23 := rowUname.Scan(&tmpUser1.ID)
	if err23 != nil {
		return ""
	} else {
		return strconv.FormatInt(tmpUser1.ID, 10)
	}
}

func GetOTP(usname string) int64 {
	db := config.Connect()
	defer db.Close()
	var tmpUser2 models.TempUsers
	sqlOTP := `SELECT otp FROM USERS WHERE username=$1;`
	rowOTP := db.QueryRow(sqlOTP, usname)
	err25 := rowOTP.Scan(&tmpUser2.Otp)
	if err25 != nil {
		return 0
	} else {
		return tmpUser2.Otp
	}
}

func ValidateEmail(mail string) int64 {
	db := config.Connect()
	defer db.Close()
	var tmpUser models.TempUsers
	//FIND EMAIL IF IT IS ALREADY EXISTS
	sqlFindEmail := `SELECT email FROM USERS WHERE email=$1;`
	rowEmail := db.QueryRow(sqlFindEmail, mail)
	err1 := rowEmail.Scan(&tmpUser.Email)
	if err1 != nil {
		return 0
	}
	return 1
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
