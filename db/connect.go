package db

import (
	"fmt"
	"log"
	"os"
	"time"

	pg "github.com/go-pg/pg"
)
func Panicking( err error){
	if err != nil {
		log.Println(" error : ",err)
	}
}

func Connect(){
	opts := &pg.Options{
		User : "dummy",
		Password : "123456",
		Addr : "localhost:5432",
    	Database : "people",
		DialTimeout: 30*time.Second,
		ReadTimeout: 1*time.Minute,
		WriteTimeout: 1*time.Minute,
		IdleTimeout: 30*time.Minute,
		// MaxAge : 1*time.Minute,
		PoolSize: 20,
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Println("failed  to connect database")
		os.Exit(100)
	}
	log.Println("connection to database successfully !!!")

	fmt.Println("#### CreateProductItemTable ### \n ")
	CreateProductItemTable(db)
	fmt.Println("### saveProduct ### \n ")
	saveProduct(db)
	fmt.Println("### saveAndReturnProduct ### \n ")
	saveAndReturnProduct(db)
	fmt.Println("### saveMultipleProduct ### \n ")
	saveMultipleProduct(db)
	fmt.Println("### placeHolderDemo ### \n ")
	placeHolderDemo(db)
	fmt.Println("### UpdateItemPrice ### \n ")
	UpdateItemPrice(db)
	fmt.Println("### Get_By_ID ### \n ")
	Get_By_ID(db)


	closeErr := db.Close()
	Panicking(closeErr)
	log.Println("Connection closed successfully !!! ")


}