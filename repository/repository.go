package repository

import (
	"grpc_streaming1/pb"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

type Repository struct{
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		db: db,
	}
}


func (r *Repository) AddProduct(product *pb.Product) (string, error){
	id := uuid.NewString()
	_, err := r.db.Exec("insert into products (id, title, price, category, rating, description) values ($1, $2, $3, $4, $5, $6)", 
	id, product.Title, product.Price, product.Category, product.Rating, product.Description)
	if err != nil{
		log.Println(err)
	}
	return id, err	
}

func (r *Repository) SerchProduct(filter *pb.ProductFilter) ([]pb.Product, error){
	var product []pb.Product
	err := r.db.Select(&product, "select * from products where category=$1 AND price<=$2 AND rating>=$3", filter.Category, filter.MaxPrice, filter.Rating)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return product, nil
}