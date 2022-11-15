package models

type Products struct {
	ID             uint   `json:"id"`
	Prod_name      string `json:"prod_name"`
	Prod_desc      string `json:"prod_desc"`
	Prod_stockqty  string `json:"prod_stockqty"`
	Prod_unit      string `json:"prod_unit"`
	Prod_cost      string `json:"prod_cost"`
	Prod_sell      string `json:"prod_sell"`
	Prod_pic       string `json:"prod_pic"`
	Prod_category  string `json:"prod_category"`
	Prod_saleprice string `json:"prod_saleprice"`
}

type Sale struct {
	ID             uint   `json:"id"`
	Prod_pic       string `json:"prod_pic"`
	Prod_desc      string `json:"prod_desc"`
	Prod_saleprice string `json:"prod_saleprice"`
}
