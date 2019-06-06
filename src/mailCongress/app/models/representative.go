package models

type Representative struct {
	Id	 string `json:'id'`
	Name string `json:'name'`
}

type Representatives []Representative