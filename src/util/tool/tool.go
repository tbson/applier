package tool

import (
	"strings"
	"github.com/Machiel/slugify"
)


func ToSlug (input string) string {
	return slugify.Slugify(input)
}

func ToAscii (input string) string {
	result := slugify.Slugify(input)
	return strings.Replace(result, "-", " ", -1)
}

func ToInterface (args []string) []interface{} {
	args1 := make([]interface{}, len(args))
	for i, v := range args {
		args1[i] = interface{}(v)
	}
	return args1
}

func DirectionParse (input string) string {
	order := ">"
	if input == "prev" {
		order = "<"
	}
	return order
}