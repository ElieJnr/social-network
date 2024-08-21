package models


type NewError struct {
	Mess  string
	Error bool
}
var Errors = map[string]NewError{
	"400":{
		"Bad Request",
		true,
	},
	"500":{
		"Internal Server Error",
		true,
	},
	"405":{
		"Not Allowed",
		true,
	},
	"404":{
		"Page Not Found",
		true,
	},
	"401":{
		"StatusUnauthorized",
		true,
	},
} 

