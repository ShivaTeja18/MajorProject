package main

import (
	"ecommerce/MajorProject"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

//func init() {
//	a, err := os.OpenFile("log.text", os.O_CREATE, os.ModePerm)
//	custerr := fmt.Errorf("failed to create file %v", a)
//	if err != nil {
//		fmt.Println(custerr, err)
//	}
//
//	wrt := io.MultiWriter(os.Stdout, a)
//	log.SetOutput(wrt)
//}
func init() {
	file, err := openLogFile("./mylog.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

}
func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
func main() {
	h := MajorProject.Dbhand()
	//var c *gin.Context
	var r = gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.53:8000"})
	r.GET("/fetchorders", h.Fetch)
	r.GET("/fetchbynum", h.FbyOrderNum)
	r.POST("/newEmp", h.CreateEmp)
	r.POST("/newProd", h.NewProd)
	r.DELETE("/remove/:employee_number", h.Delet)
	r.GET("/fetchemps/:office_code", h.FetchEmp)
	r.POST("/payments", h.Paybyid)
	r.DELETE("/rmpay", h.Remv)
	r.GET("/fpaybyid", h.FetchPay)
	r.GET("/fromtotill", h.FromtoTill)
	r.GET("/fetchprod", h.Fproducts)
	r.GET("/fet", h.Older)
	r.PATCH("/update", h.Upda)
	r.POST("/new", MajorProject.New)
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
