package handler

import (
	"encoding/json"
	"net/http"

	// "strconv"
	"github.com/labstack/echo"
)

func DetailsHandler(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	var id = c.Param("id")
	var datax, err = ambil_data_details(id)
	if err != nil {
	}

	return c.Render(http.StatusOK, "details.html", map[string]interface{}{
		"name": c.Param("id"),
		"msg":  "Hello Saya Dari Niomic",
		"data": datax,
	})
}

// var baseURL_deta = "http://localhost:1323"
func ambil_data_details(x string) ([]blog, error) {

	var err error
	var client = &http.Client{}
	var data []blog
	var id = x
	request, err := http.NewRequest("GET", baseURL+"/api/d/"+id,nil)
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
