package model

import (
	"bubblePQ/dao"
	"fmt"
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}


func Create(todo *Todo) (err error ){
	stmt, err := dao.DB.Prepare("INSERT INTO todo(title,status) VALUES($1,$2) RETURNING id")
	checkErr(err)

	res, err := stmt.Exec(todo.Title, todo.Status)
	checkErr(err)

	fmt.Println(res.RowsAffected())
	return
}

func GetAll() (todoList []*Todo,err error) {
	rows, err := dao.DB.Query("select * from todo")
	checkErr(err)
	for rows.Next(){
		var todo Todo
		err = rows.Scan(&todo.Id,&todo.Title,&todo.Status)
		fmt.Println(todo)
		todoList = append(todoList, &todo)
		checkErr(err)
	}
	return
}


func GetATodo(id string) (todo *Todo,err error) {
	rows, err := dao.DB.Query("select * from todo where id = "+id)
	checkErr(err)
	todo = new(Todo) // 一定要new 解决野指针问题
	//fmt.Println(rows.Next())
	for rows.Next() {
		rows.Scan(&todo.Id,&todo.Title,&todo.Status)
	}
	fmt.Println(todo,"geiAto")
	return
}

func UpdateATodo(todo *Todo) (err error) {
	stmt, err := dao.DB.Prepare("update todo set status=$1 where id=$2")
	checkErr(err)

	res, err := stmt.Exec(todo.Status, todo.Id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	return
}


func DeleteATOdo(id string) (err error) {
	stmt, err := dao.DB.Prepare("delete from todo where id = $1")
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	affected, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affected)
	return
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
