package main

import (
	"ecommerce/MajorProject"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	a, err := os.Create("log.text")
	custerr := fmt.Errorf("failed to create file %v", a)
	if err != nil {
		fmt.Println(custerr, err)
	} else {
		log.SetOutput(a)
	}

}
func main() {
	h := MajorProject.Dbhand()
	var r = gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.53:8000"})
	r.GET("/fetchorders", h.Fetch)
	r.GET("/fetchbynum", h.FbyOrderNum)
	r.POST("/newEmp", h.CreateEmp)
	r.POST("/newProd", h.NewProd)
	r.DELETE("/remove/:employee_number", h.Delet)
	r.GET("/fetchemps", h.FetchEmp)
	r.POST("/payments", h.Paybyid)
	r.DELETE("/rmpay", h.Remv)
	r.GET("/fpaybyid", h.FetchPay)
	r.GET("/fromtotill", h.FromtoTill)
	r.GET("/fetchprod", h.Fproducts)
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
