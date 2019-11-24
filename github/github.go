package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type File struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Size     int    `json:"size"`
	FileName string `json:"file_name"`
}

type License struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Body string `json:"body"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func List() ([]File, error) {
	var files []File

	err := getJson("https://api.github.com/repos/github/gitignore/contents", &files)
	if err != nil {
		return nil, err
	}

	var filesOut []File
	for _, item := range files {
		i := strings.Split(item.Name, ".")
		if len(i) > 0 {
			if strings.HasSuffix(item.Name, "gitignore") {
				item.FileName = i[0]
				filesOut = append(filesOut, item)
			}
		}
	}

	return filesOut, nil
}

func ListLicense() ([]License, error) {
	var lic []License
	err := getJson("https://api.github.com/licenses", &lic)
	if err != nil {
		return nil, err
	}

	return lic, nil
}

func GetLicense(l string) (License, error) {
	var lic License
	err := getJson(fmt.Sprintf("https://api.github.com/licenses/%s", l), &lic)
	if err != nil {
		return License{}, err
	}

	return lic, nil
}

func Download(url, file string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
