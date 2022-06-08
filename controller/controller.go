package controller

import (
	"bubblePQ/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}

func PostHandler(c *gin.Context) {
	var todo model.Todo
	c.BindJSON(&todo)

	err := model.Create(&todo)
	checkErr(err)

	c.JSON(http.StatusOK,todo)
}

func GetHandler(c *gin.Context) {
	all, err := model.GetAll()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,all)
	}
}

func UpdateHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error" : "无效的id"})
		return
	}

	todo, err := model.GetATodo(id)
	fmt.Println(todo)
	if  err != nil {
		c.JSON(http.StatusOK,gin.H{"error" : err.Error()})
		return
	}

	c.BindJSON(todo)
	fmt.Println(todo)
	err = model.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error" : err.Error()})
	} else {
		c.JSON(http.StatusOK,todo)
	}
}


func DeleteHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error" : "无效的id"})
		return
	}
	err := model.DeleteATOdo(id)
	if  err != nil {
		c.JSON(http.StatusOK,gin.H{"error" : err.Error()})
	} else {
		c.JSON(http.StatusOK,gin.H{id: "deleted"})
	}
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}