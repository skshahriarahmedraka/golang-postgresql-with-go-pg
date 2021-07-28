package db 

import (
	"log"
	pg "github.com/go-pg/pg"

)

func UpdateItemPrice(dbRef *pg.DB){
	newPI := &ProductItem{
		ID : 1 ,
		Price: 3443,
	}
	newPI.updatePrice(dbRef)
}

func (pi *ProductItem)updatePrice(db *pg.DB) error{
	_ , updateErr := db.Model(pi).Set("price = ?price").Where("id = ?id ").Update()
	Panicking(updateErr)

	log.Printf("Price updated successfully for ID %d \n",pi.ID)

	return nil

}