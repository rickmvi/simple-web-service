package api

import "fmt"

// GetUrl constructs and returns a GitHub API URL for the specified user.
func GetUrl(user string) (result string) {
	result = fmt.Sprintf("https://api.github.com/users/%s", user)
	return
}
