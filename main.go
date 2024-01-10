package main

import (
	"go-crud/delivery"
	"go-crud/repository"
	"go-crud/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error", err)
	}

	repoBook := repository.NewRepository(db)
	usecaseBook := usecase.NewUsecase(repoBook)
	deliveryBook := delivery.NewDelivery(usecaseBook)

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	v1 := router.Group("/v1")
	v1.POST("/book-create", deliveryBook.Create)
	v1.GET("/books", deliveryBook.GetBooks)
	v1.GET("/book/:id", deliveryBook.GetBookByID)
	v1.POST("/book/update", deliveryBook.Update)
	v1.DELETE("/delete/:id", deliveryBook.Delete)

	router.Run()
}
