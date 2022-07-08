package main

import (
	"fmt"
	"food_delivery/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id int 
}


func main() {

	dsn := helpers.GetDsn("MYSQL_CONNECTION")
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	if err := serve(db); err != nil {
		log.Fatalln(err)
	}
}


func serve(db *gorm.DB) error {
	
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r.Run(`:8080`);
}