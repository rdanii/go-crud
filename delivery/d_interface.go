package delivery

import (
	"go-crud/entity"
	"go-crud/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Delivery interface {
	Create(c *gin.Context)
	GetBooks(c *gin.Context)
	GetBookByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type delivery struct {
	usecase usecase.Usecase
}

func NewDelivery(usecase usecase.Usecase) Delivery {
	return &delivery{usecase}
}

func (d *delivery) Create(c *gin.Context) {
	var reqBook entity.BookRequest

	err := c.ShouldBindJSON(&reqBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, errResult := d.usecase.Create(reqBook)
	if errResult != nil {
		if strings.Contains(errResult.Error(), "data already exists") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "data already exists",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errResult.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": result,
	})
}

func (d *delivery) GetBooks(c *gin.Context) {
	result, err := d.usecase.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (d *delivery) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id - only numbers allowed",
		})
		return
	}

	result, errResult := d.usecase.GetBookByID(idInt)
	if errResult != nil {
		if strings.Contains(errResult.Error(), "connection refused") {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "database connection lost",
			})
			return
		}

		if strings.Contains(errResult.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "data not found",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errResult.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (d *delivery) Update(c *gin.Context) {
	var reqBook entity.BookRequestUpdate
	err := c.ShouldBindJSON(&reqBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, errResult := d.usecase.Update(reqBook)
	if errResult != nil {
		if strings.Contains(errResult.Error(), "data already exists") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no rows updated",
			})
			return
		}
		if strings.Contains(errResult.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "data not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errResult.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": result,
	})
}

func (d *delivery) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, errID := strconv.Atoi(id)
	if errID != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not valid - only number",
		})
		return
	}
	err := d.usecase.Delete(idInt)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "database connection lost",
			})
			return
		}
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "data not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "data has been deleted",
	})
}
