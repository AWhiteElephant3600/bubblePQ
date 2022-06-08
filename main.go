package main

import (
	"bubblePQ/dao"
	"bubblePQ/routes"
)

func main() {
	db := dao.DB

	defer db.Close()
	//
	////插入数据
	//stmt, err := db.Prepare("INSERT INTO todo(title,status) VALUES($1,$2) RETURNING id")
	//checkErr(err)
	//
	//res, err := stmt.Exec("吃饭", 0)
	//checkErr(err)
	//
	//fmt.Println(res.RowsAffected())

	r := routes.SetupRouter()

	r.Run()


}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
