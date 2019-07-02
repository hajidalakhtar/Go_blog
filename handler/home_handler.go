package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

var baseURL = "http://localhost:1323"

func HomeHandler(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go

	var datax, err = ambil_data()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "HOME",
		"msg":  "Framework Echo Workedsss!sss",
		"data": datax,
	})
}

func ambil_data() ([]blog, error) {
	var err error
	var client = &http.Client{}
	var data []blog

	request, err := http.NewRequest("GET", baseURL+"/api", nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
