// model/user.go
package model

type User struct { // User struct has two fields ID and Name, both of type int and string respectively
	ID   int    `json:"id"`
	Name string `json:"name"`
}
