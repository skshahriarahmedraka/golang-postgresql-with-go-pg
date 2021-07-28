package db

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
)

func saveProduct(dbRef *pg.DB){

	newProduct := &ProductItem{
		Name : "mango",
		Desc : "sweet fruit",
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
	newProduct.Save(dbRef)
}

func (pi *ProductItem) Save(dBase *pg.DB) error {
	insertErr := dBase.Insert(pi)
	Panicking(insertErr)
	log.Println("successfully inserted productItem : ", pi.Name)
	return nil
}

