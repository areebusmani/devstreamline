package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func HandleConvertRequest(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("file")
	re, err := regexp.Compile(`https?://www\.figma\.com/file/([^/]+)/([^/]+)`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	matches := re.FindStringSubmatch(url)
	var key string
	if len(matches) > 2 {
		key = matches[1]
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	figmaEndpoint := fmt.Sprintf("https://api.figma.com/v1/files/%s", key)

	figmaRequest, err := http.NewRequest("GET", figmaEndpoint, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	figmaRequest.Header.Set("X-Figma-Token", FIGMA_ACCESS_TOKEN)
	client := http.Client{}
	figmaResponse, err := client.Do(figmaRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer figmaResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(figmaResponse.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(responseBody)
	return
}
