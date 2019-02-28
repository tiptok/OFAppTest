package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type IConsignment interface {
	Create(*Consignment) (*Consignment, error)
	GetAll() []*Consignment
}

type ConsignmentCache struct {
	Cache []*Consignment
}

func (repo *ConsignmentCache) Create(consignment *Consignment) (*Consignment, error) {
	updated := append(repo.Cache, consignment)
	repo.Cache = updated
	return consignment, nil
}

func (repo *ConsignmentCache) GetAll() []*Consignment {
	log.Println("GetAll")
	return repo.Cache
}

type ConsignmentService struct {
	repo IConsignment
}

func (s *ConsignmentService) CreateConsignment(ctx context.Context, req *Consignment) (*Response, error) {
	c, err := s.repo.Create(req)
	if err != nil {
		log.Println("CreateConsignment Error.", err)
		return nil, err
	}

	return &Response{Created: true, Consignment: c}, nil
}

func (s *ConsignmentService) GetAllConsignment(ctx context.Context, req *RequestAll) (*ResponseAll, error) {
	consignments := s.repo.GetAll()
	return &ResponseAll{Consignments: consignments}, nil
}

func main() {
	cache := &ConsignmentCache{}

	lis, err := net.Listen("tcp", ":9927")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer()
	RegisterShipServiceServer(s, &ConsignmentService{cache})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server:%v", err)
	}
}
