package handlers

import (
	"devstreamline/figma"
	"encoding/json"
	"errors"
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
	file, err := fetchFigmaFile(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, _ := json.Marshal(file)
	w.Write([]byte(resp))
}

func fetchFigmaFile(key string) (file figma.File, err error) {
	file = figma.File{}

	figmaEndpoint := fmt.Sprintf("https://api.figma.com/v1/files/%s", key)

	figmaRequest, figmaErr := http.NewRequest("GET", figmaEndpoint, nil)
	if figmaErr != nil {
		err = errors.New("Error while fetching figma file")
		return
	}
	figmaRequest.Header.Set("X-Figma-Token", FIGMA_ACCESS_TOKEN)
	client := http.Client{}
	figmaResponse, figmaErr := client.Do(figmaRequest)
	if err != nil {
		err = errors.New("Error while fetching figma file")
		return
	}
	defer figmaResponse.Body.Close()

	responseBody, figmaErr := ioutil.ReadAll(figmaResponse.Body)
	if err != nil {
		err = errors.New("Error while parsing figma file")
		return
	}
	json.Unmarshal(responseBody, &file)
	return
}
