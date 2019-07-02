package handler

import (
	"echoweb-master/server"
	"fmt"
	_ "mysql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type blog struct {
	Id    string
	Title string
	Isi   string
}

var data []blog

func Baca_data(c echo.Context) error {

	DataBlog()
	return c.JSON(http.StatusOK, data)
}
func Baca_data_id(c echo.Context) error {

	var id = c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
	}
	DataBlogId(i)
	return c.JSON(http.StatusOK, data)
}
func DataBlogId(x int) {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	id := x
	rows, err := db.Query("SELECT `id`, `title`, `isi` FROM `blog` WHERE id = ?", id)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	defer rows.Close()
	for rows.Next() {
		var each = blog{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Isi)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func DataBlog() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from blog")
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	defer rows.Close()
	for rows.Next() {
		var each = blog{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Isi)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
func TambahData(c echo.Context) error {

	db, err := server.Koneksi()
	defer db.Close()

	var Title = c.FormValue("title")
	var Isi = c.FormValue("isi")

	_, err = db.Exec("insert into blog values (?,?,?)", nil, Title, Isi)

	if err != nil {
		return c.JSON(http.StatusOK, "gagal Menambahkan")
		} else {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

}
func UbahData(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()

	var Title = c.FormValue("title")
	var Isi = c.FormValue("isi")
	var Id = c.FormValue("id")

	_, err = db.Exec("update blog set title = ? ,isi = ? where id = ? ", Title, Isi, Id)

	if err != nil {
		return c.JSON(http.StatusOK, "gagal update")
	} else {
		return c.JSON(http.StatusOK, "berhasil update")

	}
}
func HapusData(c echo.Context) error {

	db, err := server.Koneksi()
	defer db.Close()

	var id = c.FormValue("id")

	_, err = db.Exec("DELETE FROM `blog` WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusOK, "gagal delete")

	} else {
		return c.JSON(http.StatusOK, "berhasil delete")

	}

}
