package main

import (
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)
// import "net/http"

func main() {
	r := gin.Default()
	// router := gin.Default()/\
	r.LoadHTMLGlob("template/*.html")
	db, err := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/userlist")
	
	fmt.Println("this is database",db)
	// results, err := db.Query("SELECT id, name FROM tags")
    // if there is an error opening the connection, handle it
    if err != nil {
		fmt.Println("error")
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
	r.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{})
	})
	r.GET("/data", func(c *gin.Context) {
		c.HTML(200,"data.html",gin.H{"data":"Helo Gin World"})
	})
	r.GET("/form", func(c *gin.Context) {
		c.HTML(200,"form.html",gin.H{"data":"Helo Gin World"})
	})
	r.GET("json", func (c *gin.Context){
		c.JSON(200 ,gin.H{
			"result":"ok",
			"data":"Helo Go/go/ginWorld",
			"developer":"komavideo",
		})
	})
	r.POST("/service", func (c *gin.Context){
		uname :=c.PostForm("uname")
		password :=c.PostForm("password")

		c.JSON(200 ,gin.H{
			"result":"ok",
			"hello":uname,
			"password":password,

		})
	})
	
	
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
