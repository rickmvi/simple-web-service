package domain

import (
	"encoding/json"
	"github.com/rickmvi/simple-web-service/internal/api"
	"io"
	"net/http"
)

type CatApi struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func HandlerCatApi(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, api.CatApiUrl, nil)

	if err != nil {
		http.Error(w, "error creating request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, "error when querying CatApi: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(Body io.ReadCloser) {
		err := resp.Body.Close()
		if err != nil {
			http.Error(w, "error closing response body: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "CatApi returned: "+string(body), resp.StatusCode)
		return
	}

	if err != nil {
		http.Error(w, "error reading response from CatApi: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var cats []CatApi
	if err := json.Unmarshal(body, &cats); err != nil {
		http.Error(w, "error when unmarshal JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(cats) == 0 {
		http.Error(w, "no cat found", http.StatusInternalServerError)
		return
	}

	response := struct {
		Image string `json:"image"`
	}{
		Image: cats[0].Url,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "error when encoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
