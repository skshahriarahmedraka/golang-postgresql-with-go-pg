package db

import (
	//"fmt"
	"log"
	"time"

	pg "github.com/go-pg/pg"
)

func saveMultipleProduct(dbRef *pg.DB){
	newProduct1 := &ProductItem{
		Name : "lotkon",
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
	newProduct2 := &ProductItem{
		Name : "banana",
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
	totalItems := []*ProductItem{
		newProduct1,newProduct2,
	}
	newProduct1.saveMultiple(dbRef , totalItems)
}

func (pi *ProductItem) saveMultiple(db *pg.DB , items []*ProductItem) error {
	_ , insertErr := db.Model(items[0],items[1]).Insert() 
	Panicking(insertErr)
	log.Println("bulk insert successfully")
	return nil 

}