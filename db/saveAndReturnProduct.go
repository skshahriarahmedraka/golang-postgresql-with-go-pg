package db

import (
	//"fmt"
	"log"
	"time"

	pg "github.com/go-pg/pg"
)

func saveAndReturnProduct(dbRef *pg.DB){

	newProduct := &ProductItem{
		Name : "lichi",
		Desc : "sweet lichi fruit",
		Image : " file path of image ",
		Price : 23,
		Features : struct{Name string; Desc string; Imp int}  {
			Name : "F1",
			Desc : "F1 description",
			Imp : 3 ,
		},
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		IsActive : true ,
	}
	updatedPi, err :=newProduct.saveAndReturn(dbRef)
	Panicking(err)
	
	log.Println("returned value   : ", updatedPi)
}

func (pi *ProductItem) saveAndReturn(dBase *pg.DB) (*ProductItem,error) {
	InsertResult,insertErr := dBase.Model(pi).Returning("*").Insert()
	Panicking(insertErr)
	log.Println("successfully inserted and return productItem : ", pi.Name)
	log.Println("successfully inserted and return productItem : ", InsertResult)
	
	return pi, nil
}

