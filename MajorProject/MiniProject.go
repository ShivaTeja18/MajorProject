package MajorProject

import (
	"ecommerce/details"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//type Handler struct {
//	DB *gorm.DB
//}

func (h Handler) Fetch(c *gin.Context) {
	var cust []details.Orders
	id := c.Param("customer_number")
	//if id != "customer_number" {
	//	err := errors.New("id not found no customer belongs to this id")
	//	if err != nil {
	//		c.AbortWithStatusJSON(http.StatusBadRequest, details.Response{Status: "UNSUCCESSFUL", Error: err, Code: http.StatusBadRequest, Data: &cust})
	//	}

	if err := h.DB.Model(details.Orders{}).Where("customer_number = ?", &id).Find(&cust).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		return
	}

	if len(cust) <= 0 {
		log.Println(len(cust))
		c.IndentedJSON(http.StatusOK, details.Response{
			Status: "failed",
			Error:  "No data found",
			Code:   http.StatusNotFound,
			Data:   nil,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, details.Response{
		Status: "Success",
		Error:  "",
		Code:   http.StatusOK,
		Data:   &cust,
	})
	return
}
func (h Handler) FbyOrderNum(c *gin.Context) {
	var order []details.Orderdetail
	id := c.Param("order_number")
	fmt.Println("got id ", id)
	if err := h.DB.Model(details.Orderdetail{}).Preload("Products").Where("order_number =?", id).Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
			Code:   http.StatusBadRequest,
			Data:   &order,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, details.Response{
		Status: "Success",
		Error:  "",
		Code:   http.StatusOK,
		Data:   &order,
	})

	return
}

//
//func Dbhand() Handler {
//	url := os.Getenv("dns")
//	return Handler{DB: dbc.Dbinit(url)}
//}
