package details

import "time"

type Customers struct {
	CustomerNumber   string
	CustomerName     string
	ContactLastName  string
	ContactFirstName string
	Phone            string
	AddressLine1     string
}

type Orders struct {
	OrderNumber    int
	OrderDate      string
	RequiredDate   string
	ShippedDate    string
	Status         string
	Comments       string
	CustomerNumber int16 `json:"customerNumber"`
}

type Orderdetail struct {
	OrderNumber int
	ProductCode string
	Products    Product `gorm:"foreignKey:ProductCode;references:ProductCode"`
	//ProductCode string
	QuantityOrdered int16
	PriceEach       float64
	OrderLineNumber int16
}

type Product struct {
	ProductCode     string
	ProductName     string
	ProductLine     string
	ProductScale    string
	ProductVendor   string
	QuantityInStock int
	BuyPrice        float64
	Msrp            float64
}

type Response struct {
	Status string
	Error  string
	Code   int
	Data   interface{}
}

type Employee struct {
	EmployeeNumber int
	LastName       string
	FirstName      string
	Extension      string
	Email          string
	OfficeCode     string
	ReportsTo      int
	JobTitle       string `gorm:"column:job_Title"`
}

type ProductLine struct {
	ProductLine     string
	TextDescription string
	HtmlDescription string
	Image           []byte //`gorm:"-:all"`
}

type Payment struct {
	CustomerNumber int
	CheckNumber    string    `gorm:"primary_key"`
	PaymentDate    time.Time `json:"paymentDate"`
	Amount         float64
}
