package db 

import (
	"log"
	pg "github.com/go-pg/pg"

)

func DeleteItem(dbRef *pg.DB){
	newPI := &ProductItem{
		Name : "mango",
	} 
	newPI.delete(dbRef)
}

func (pi *ProductItem) delete(db *pg.DB) error {
	_ , deleteErr := db.Model(pi).Where("name= ?name").Delete()
	Panicking(deleteErr)
	log.Printf("Delete successful for %s ,item \n",pi.Name)
	return nil
}