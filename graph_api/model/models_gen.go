// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Order struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Weight      string `json:"weight"`
	Phonenumber string `json:"phonenumber"`
	Countrycode string `json:"countrycode"`
}

type OrderInput struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Weight      string `json:"weight"`
	Phonenumber string `json:"phonenumber"`
	Countrycode string `json:"countrycode"`
}
