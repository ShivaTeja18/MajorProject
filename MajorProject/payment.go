package MajorProject

import (
	"ecommerce/details"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) Paybyid(c *gin.Context) {
	id := c.Request.FormValue("customer_number")
	var pay []details.Payment
	var pub details.Payment
	atpi, _ := strconv.Atoi(id)
	parse, err := time.Parse("2006-01-02", c.PostForm("payment_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "CHECK REQUIRED",
			Code:   http.StatusBadRequest,
			Error:  err.Error(),
			Data:   nil,
		})
		return
	}
	fv, _ := strconv.ParseFloat(c.PostForm("Amount"), 64)
	pub = details.Payment{
		CustomerNumber: atpi,
		CheckNumber:    c.PostForm("CheckNumber"),
		PaymentDate:    parse,
		Amount:         fv,
	}
	if err := c.Bind(&pub); err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "CheckRequired",
			Error:  err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		log.Println(err)
		return
	}
	pay = append(pay, pub)
	//if result:= h.DB.Model(&details.Payment{}).Where("customer")
	if err := h.DB.Model(&details.Payment{}).Where("customer_number = ? ", id).Find(&pay).Create(&pub).Error; err != nil {
		fmt.Println(err.Error())
		c.PureJSON(http.StatusNoContent, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  err.Error(),
			Code:   http.StatusNoContent,
			Data:   nil,
		})
		log.Println(err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Code:   http.StatusOK,
		Data:   &pub,
	})
	return
}

func (h Handler) Remv(c *gin.Context) {
	var pays details.Payment
	id := c.Request.FormValue("check_number")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, details.Response{
			Status: "Failuer",
			Error:  "check value entered not be empty",
		})
	}

	if err := h.DB.Model(&details.Payment{}).Where("check_number = ?", id).Last(&pays).Delete(&pays).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response{Status: "unsuccessful", Error: err.Error(), Code: http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Error:  "",
		Code:   http.StatusOK,
		Data:   &pays,
	})
	return
}

func (h Handler) FetchPay(c *gin.Context) {
	var pays []details.Payment
	id := c.Request.FormValue("customer_number")
	if err := h.DB.Model(&details.Payment{}).Where("customer_number = ?", id).Find(&pays).Error; err != nil {
		c.JSON(http.StatusNotFound, details.Response{
			Status: "CHECK ID PROVIDED",
			Error:  err.Error(),
			Code:   http.StatusNotAcceptable,
			Data:   nil,
		})
		return
	}
	c.JSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Error:  "",
		Code:   http.StatusOK,
		Data:   &pays,
	})
	return
}

func (h Handler) Upda(c *gin.Context) {
	var pay details.Payment
	eid := c.Request.PostFormValue("CustomerNumber")
	id, _ := strconv.Atoi(eid)

	CheckNumber := c.Request.FormValue("CheckNumber")

	if id == 0 || CheckNumber == "" {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "UNSUCCESSFUL",
			Error:  "Fields Should Not Be Empty",
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		return
	}
	if err := h.DB.Model(&details.Payment{}).Where("customer_number = ?", id).Update("check_number", CheckNumber).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response1{
			Status: "UNSUCCESSFUL",
			Error:  err,
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		return
	}
	h.DB.Model(&pay).Save(CheckNumber)
	c.JSON(http.StatusOK, details.Response{
		Status: "Successful",
		Error:  "",
		Code:   http.StatusOK,
		Data:   &pay,
	})
	return
}

//dd := time.Date(2000, 01, 01, 00, 00, 00, 00, time.UTC)
//d:=strconv.Itoa(dd.Day())
//m:=strconv.Itoa(int(dd.Month()))
//y:=strconv.Itoa(dd.Year())
//a := fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v", dd.Year(), dd.Format("01"), dd.Day(), dd.Hour(), dd.Minute(), dd.Second(), dd.UTC())
//if err := h.DB.Model(&details.Payment{}).Where("check_number = ?", id).Delete(&pays).Error; err != nil {
//	c.JSON(http.StatusBadRequest, details.Response{
//		Status: "UNSUCCESSFUL",
//		Error:  err.Error(),
//		Code:   http.StatusBadRequest,
//		Data:   nil,
//	})
//	return
//}
