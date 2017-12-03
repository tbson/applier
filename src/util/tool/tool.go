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