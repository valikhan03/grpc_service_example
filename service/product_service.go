package service

import (
	"context"
	"errors"
	"grpc_streaming1/pb"
	"grpc_streaming1/repository"
	"log"
)

type ProductService struct{
	repository *repository.Repository
}

func NewProductService(rep *repository.Repository) *ProductService{
	return &ProductService{
		repository: rep,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error){
	product := req.GetProduct()

	id, err := s.repository.AddProduct(product)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	res := &pb.CreateProductResponse{
		Id: id,
	}

	return res, nil
}

func (s *ProductService) SearchProduct(req *pb.SearchProductRequest, stream pb.ProductService_SearchProductServer) error{
		products, err := s.repository.SerchProduct(req.GetFilter())
		if err != nil{
			return errors.New("Error occured during reading from DB")
		}

		for _, p := range products{
			if err = stream.Send(&pb.SearchProductResponse{
				Product: &p,
			}); err != nil{
				return err
			}
		}

		return nil
}