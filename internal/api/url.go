package api

import "fmt"

const (
	CatApiUrl = "https://api.thecatapi.com/v1/images/search"
)

// GetUrl constructs and returns a GitHub API URL for the specified user.
func GetUrl(user string) (result string) {
	result = fmt.Sprintf("https://api.github.com/users/%s", user)
	return
}

// Url formats a URI string by injecting the provided data and returns the resulting URL string.
func Url(uri string, data string) (url string) {
	url = fmt.Sprintf(uri, data)
	return
}

// Uri formats the given uri string and returns the formatted URL string.
func Uri(uri string) (url string) {
	url = fmt.Sprintf(uri)
	return
}
