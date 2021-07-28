package db 

import (
	"log"
	pg "github.com/go-pg/pg"

)

func Get_By_ID(dbRef *pg.DB){
	newPI := &ProductItem{
		ID : 1 ,
		// Name : "lichi",
	}
	newPI.GetByID(dbRef)
}

func (pi *ProductItem) GetByID(db *pg.DB) error {
	// getErr := db.Select(pi)
	getErr := db.Model(pi).Where("id=?0",pi.ID).Select()
	// getErr := db.Model(pi).Column("name","desc").Where("id=?0",pi.ID).Where("name=?0",pi.Name).Select()
	// getErr := db.Model(pi).Column("name","desc").
	// Where("id=?0",pi.ID).
	// WhereOr("id=?0",2).
	// Offset(0).
	// Limit(2).
	// Order("id asc").
	// Select()
	
	Panicking(getErr)
	log.Printf("get by id successful : %v \n",*pi)
	return nil
}