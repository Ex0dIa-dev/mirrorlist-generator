package helpers

import "strings"

// CheckErr panics if there is a error
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

// ReturnAsArrays returns the given string-comma-separated in a array
func ReturnAsArrays(s string) []string {

	return strings.Split(s, ",")

}
