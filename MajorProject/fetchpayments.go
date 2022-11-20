package MajorProject

import (
	"ecommerce/details"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h Handler) FromtoTill(c *gin.Context) {
	var fpay []details.Payment
	from := c.Request.FormValue("payment_date")
	fmt.Println("from ", from)
	parse, err := time.Parse("2006-01-02", from)
	if err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "check requires in parsing time parse",
			Error:  err.Error(),
		})
		return
	}
	toDate := c.Request.FormValue("to")
	to, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "check requires in parsing time to",
			Error:  err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
		return
	}
	if from == "" || toDate == "" {
		c.JSON(http.StatusNotAcceptable, details.Response{
			Status: "FAILURE",
			Error:  "Fields should not be empty",
			Code:   http.StatusNotAcceptable,
			Data:   nil,
		})
		return
	}

	if err := h.DB.Model(&details.Payment{}).Where("payment_date BETWEEN ? AND ?", parse, to).Find(&fpay).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "check require",
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, details.Response{
		Status: "SUCCESSFUL",
		Error:  "",
		Code:   http.StatusOK,
		Data:   &fpay,
	})
	return
}

//v := parse.mo
//to := c.Request.FormValue("payment_date")
//if err := h.DB.Model(details.Payment{}).Where("payment_date BETWEEN ? AND ?", parse, to).Find(&fpay).Error; err != nil {
//	c.JSON(http.StatusNotAcceptable, details.Response{
//		Status: "UNSUCCESSFUL",
//		Error:  err.Error(),
//		Code:   http.StatusNotAcceptable,
//		Data:   nil,
//	})
//	return
//	//	//} else if from != "" {
//	//	//	err := h.DB.Model(details.Payment{}).Find(&fpay, "payment_date = ? AND payment_date = ?", from, to).Error
//	//	//	if err != nil {
//	//	//		c.JSON(http.StatusNotAcceptable, details.Response{
//	//	//			Status: "UNSUCCESSFUL",
//	//	//			Error:  err.Error(),
//	//	//			Code:   http.StatusNotAcceptable,
//	//	//			Data:   nil,
//	//	//		})
//	//	//		return
//	//	//	}
//	//	//	return
//	//	//}
//}
//for i, v := range fpay {
//	fmt.Println("val = ", i, v)
//}
//for parse.Before(to) {
//	if err := h.DB.Model(&details.Payment{}).Where("payment_date BETWEEN ? AND ?", parse, to).Find(fpay).Error; err != nil {
//		c.JSON(http.StatusBadRequest, details.Response{
//			Status: "FAILURE",
//			Error:  err.Error(),
//		})
//	}
//	//pa = append(pa, fpay)
//	fmt.Println("value", err)
//	return
//}

func (h Handler) Older(c *gin.Context) {
	var old []details.Payment
	date := c.Request.FormValue("date")
	older, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, details.Response{
			Status: "CHECK FORMAT",
			Error:  err.Error(),
		})
		return
	}
	if err := h.DB.Model(&details.Payment{}).Where("payment_date<?", older).Order("payment_date Desc").Find(&old).Error; err != nil {
		c.JSON(http.StatusBadRequest, details.Response{
			Status: "unsu c",
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, details.Response{
		Status: "Successful",
		Error:  "",
		Code:   http.StatusOK,
		Data:   old,
	})
}
