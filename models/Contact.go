package models

type Contact struct {
	ID               uint   `json:"id"`
	Contact_Name     string `json:"contactname"`
	Contact_Address  string `json:"contactaddress"`
	Contact_Email    string `json:"contactemail"`
	Contact_Mobileno string `json:"contactmobileno"`
	IsActive         int64  `json:"isactive,string,omitempty"`
}
