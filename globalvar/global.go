package globalvar

var USERNAME string
var USERID string
var MESSAGE string
var USERPIC string

func SetUserpic(pic string) {
	USERPIC = pic
}

func GetUserpic() string {
	return USERPIC
}

func SetMessage(message string) {
	MESSAGE = message
}

func GetMessage() string {
	return MESSAGE
}

func SetUsername(username string) {
	USERNAME = username
}

func GetUsername() string {
	return USERNAME
}

func SetUserid(userid string) {
	USERID = userid
}

func GetUserid() string {
	return USERID
}
