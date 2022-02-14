package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"grpc_streaming1/pb"
	"grpc_streaming1/repository"
	"grpc_streaming1/service"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func initDB() *sqlx.DB {
	godotenv.Load("db.env")
	conn_str := fmt.Sprintf("host=localhost port=5432 user=postgres password=%s dbname=postgres sslmode=disable", os.Getenv("POSTGRES_PASSWORD"))
	db, err := sqlx.Open("pgx", conn_str)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	product_repository := repository.NewRepository(initDB())
	products_service := service.NewProductService(product_repository)

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, products_service)

	listener, err := net.Listen("tcp", "127.0.0.1:8060")
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
