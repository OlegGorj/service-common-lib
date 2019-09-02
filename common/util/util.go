package util

import (
	"os"
)

//@params
//	environment variable
//@return
//	value of environment variable

func GetENV(k string) string {
	v := os.Getenv(k)
	if v == "" {
		//loggerEntry.Fatalf("%s environment variable not set.", k) // disable for now
	}
	return v
}
