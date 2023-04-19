package handlers

import (
	"devstreamline/converter"
	"devstreamline/figma"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Handles request /convert?file=<figma_url> by fetching the figma file and
// converting it into frontend code.
func HandleConvertRequest(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Query().Get("file")
	re, err := regexp.Compile(`https?://www\.figma\.com/file/([^/]+)/([^/]+)`)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	matches := re.FindStringSubmatch(url)
	var key string
	if len(matches) > 2 {
		key = matches[1]
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	file, err := fetchFigmaFile(key)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	// TODO: Remove hard coded figma node id.
	converter.Convert(file.GetNode("1:1652"))

	type ConvertTemplateData struct {
		Markup string
	}

	convertTemplate, _ := template.ParseFiles("./templates/convert.html")
	convertTemplateData := ConvertTemplateData{Markup: "Hello world"}

	convertTemplate.Execute(writer, convertTemplateData)
}

// Handles request /json?file=<figma_url> by fetching the figma file
// represented as json.
func HandleJsonRequest(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Query().Get("file")
	re, err := regexp.Compile(`https?://www\.figma\.com/file/([^/]+)/([^/]+)`)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	matches := re.FindStringSubmatch(url)
	var key string
	if len(matches) > 2 {
		key = matches[1]
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	file, err := fetchFigmaFile(key)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, _ := json.Marshal(file)
	writer.Write([]byte(resp))
}

// Fetches the figma file from the figma server.
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
