package MajorProject

import (
	"ecommerce/details"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

func (h Handler) CreateEmp(c *gin.Context) {
	var newEmp details.Employee
	a, _ := strconv.Atoi(c.PostForm("EmployeeNumber"))
	b, _ := strconv.Atoi(c.PostForm("ReportsTo"))
	newEmp = details.Employee{
		EmployeeNumber: a,
		LastName:       c.PostForm("LastName"),
		FirstName:      c.PostForm("FirstName"),
		Extension:      c.PostForm("Extension"),
		Email:          c.PostForm("Email"),
		OfficeCode:     c.PostForm("OfficeCode"),
		ReportsTo:      b,
		JobTitle:       c.PostForm("Test"),
	}
	if err := c.Bind(&newEmp); err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		fmt.Println(err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(&newEmp); err != nil {
		c.JSON(http.StatusUnauthorized, details.Response{
			Status: "Check Fields Required",
			Error:  err.Error(),
			Code:   http.StatusUnauthorized,
			Data:   &newEmp,
		})
		return
	}

	if err := h.DB.Model(&details.Employee{}).Create(&newEmp).Error; err != nil {
		c.JSON(http.StatusNotFound, details.Response{
			Status: "Unsuccessful",
			Error:  err.Error(),
			Code:   http.StatusNotFound,
			Data:   nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, details.Response{
			Status: "SUCCESSFUL",
			Error:  "",
			Code:   http.StatusOK,
			Data:   &newEmp,
		})
	}
	return
}

func (h Handler) Delet(c *gin.Context) {
	id := c.Param("employee_number")
	var nemp details.Employee
	if err := h.DB.Model(&details.Employee{}).Where("employee_number=?", id).First(&nemp).Delete(&nemp).Error; err != nil {
		c.JSON(http.StatusNotFound, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
			Code:   http.StatusNotFound,
			Data:   nil,
		})
		log.Println(err)
		return
	}

	//if err := h.DB.Model(&details.Employee{}).Where("employee_number", id).Error; err != nil {
	//	c.JSON(http.StatusGone, err)
	//}

	c.JSON(http.StatusOK, details.Response{Status: "successful", Error: "", Code: http.StatusOK, Data: &nemp})
	return
}

func (h Handler) FetchEmp(c *gin.Context) {
	var emps []details.Employee
	id := c.Request.FormValue("office_code")
	if id == "" {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  "ID should not be empty",
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		return
	}

	if err := h.DB.Model(&details.Employee{}).Where("office_code = ?", id).Find(&emps).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		return
	}
	c.JSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Error:  "",
		Code:   http.StatusOK,
		Data:   emps,
	})
	return
}
