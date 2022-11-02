package dbc

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//var Db *gorm.DB

func Dbinit(Url string) *gorm.DB {
	//var Data *gorm.DB
	//const DBURL = `Host = psql-mock-database-cloud.postgres.database.azure.com User = jzoxvjombiuzvjolrtikkns@psql-mock-database-cloud Password = yrddycltuxgnukwshiiftbfi DBname = ecom1666192780245qvwvcnhcdzktnoum`
	//postgres := `host=psql-mock-database-cloud.postgres.database.azure.com user=jzoxvjombiuzvjolvrtikkns@psql-mock-database-cloud password=yrddycltuxgnukwshiiftbfi dbname = ecom1666192780245qbwvcnhcdzktnoum sslmode = disable`
	DB, err := gorm.Open(postgres.Open(Url), &gorm.Config{})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("connected")
	}
	//cerr := errors.New("failed Migration")
	//if err != nil {
	//	return
	//}
	//err = DB.AutoMigrate()
	//if err != nil {
	//	log.Println(err)
	//	Db = DB
	return DB
}
