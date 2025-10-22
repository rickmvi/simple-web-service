package domain

import (
	"encoding/json"
	"fmt"
	"github.com/rickmvi/simple-web-service/internal/api"
	"io"
	"net/http"
	"time"
)

// GithubUser represents a GitHub user profile with relevant information retrieved from the GitHub API.
type GithubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	AvatarUrl   string `json:"avatar_url"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
}

// Handler processes HTTP GET requests to fetch and return GitHub user profile data in JSON format.
func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	url := api.GetUrl(username)

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		http.Error(w, "error creating request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("User-Agent", "go-github-client")

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, "error when querying GitHub: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "error closing response body: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("GitHub returned: %d - %s", resp.StatusCode, string(body)), resp.StatusCode)
		return
	}

	if err != nil {
		http.Error(w, "error reading response from GitHub: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user GithubUser
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "error when unmarshal JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	out := GithubUser{
		Login:       user.Login,
		Name:        user.Name,
		AvatarUrl:   user.AvatarUrl,
		Bio:         user.Bio,
		PublicRepos: user.PublicRepos,
		Followers:   user.Followers,
		Following:   user.Following,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(out); err != nil {
		http.Error(w, "error when encoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
