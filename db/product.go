package db

import (
	orm "github.com/go-pg/pg/orm"
	pg "github.com/go-pg/pg"
	"time"
	"log"
)

type ProductItem struct {
	RefPointer int `sql:"-"`
	tableName struct{} `sql:"product_item_collection"`
	ID int `sql:"id,pk"`
	Name string `sql:"name,unique"`
	Desc string  `sql:"desc"`
	Image string `sql:"image"`
	Price float64 `sql:"price,type:real"`
	Features struct {
		Name string 
		Desc string 
		Imp int 
	} `sql:"feature,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive bool `sql:"is_active"`
}
func CreateProductItemTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr:= db.CreateTable(&ProductItem{},opts)
	Panicking(createErr)
	log.Println("table created successfully")
	return nil 

}