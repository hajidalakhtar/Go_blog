package server

import(
	"database/sql"
	_"mysql"
)

func Koneksi() (*sql.DB ,error) {
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/Go_blog")
	if err != nil {
		return nil,err
	}
	return db ,nil


}