package models

type Contact struct {
	ID               int64  `json:"id"`
	Contact_name     string `json:"contactname"`
	Contact_address  string `json:"contactaddress"`
	Contact_email    string `json:"contactemail"`
	Contact_mobileno string `json:"contactmobileno"`
	Isactive         int64  `json:"isactive,string,omitempty"`
}

type Contacts struct {
	ID               int64
	Contact_name     string
	Contact_address  string
	Contact_email    string
	Contact_mobileno string
	Isactive         int64
}
