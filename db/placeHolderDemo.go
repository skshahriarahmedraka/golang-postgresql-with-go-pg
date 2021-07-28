package db 

import (
	"log"
	pg "github.com/go-pg/pg"

)
type Params struct {
	Param1 string 
	Param2 string 
}

func placeHolderDemo(db *pg.DB) error {
	var value int
	params := Params {
		Param1: "this is one",
		Param2: "this is two",
	} 

	var query string ="SELECT ?param2"
	_,selectErr := db.Query(pg.Scan(&value),query,params)
	Panicking(selectErr)

	log.Printf("Scan Successful , scaned value %d \n ",value)
	return nil

}