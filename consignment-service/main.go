package main

import (
	"context"
	"log"

	pb "shipper/consignment-service/consignment"

	"go-micro.dev/v4"
)

//Define a local variable as Repository
type Repository struct {
	consignments []*pb.Consignment
}

//Define an interface for Repository towards local variable/database
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetList() []*pb.Consignment
}

//Impletement the Create method towards lcoal varialble/database
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

//Implet the GetList Method towards local variable/database
func (repo *Repository) GetList() []*pb.Consignment {
	return repo.consignments
}

//service Impletement all the methods defined in .protoc, then service will get theses output to database
type service struct {
	repo IRepository
}

//service.CreateConsignment:parse the context from client.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	//save the consignment to the database;
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	//return Response and defined data structure
	res.Created = true
	res.Consignment = consignment
	return nil
}

//service.GetConsignmentList:
func (s *service) GetConsignmentList(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	//get the consignment list from database
	consignments := s.repo.GetList()
	res.Consignments = consignments
	//return Response and defined data structure to the client
	return nil
}
func main() {

	repo := &Repository{}

	//register a gRPC server on gRPC server
	s := micro.NewService(

		micro.Name("shippy.service.consignment"),
	)

	s.Init()

	if err := pb.RegisterShippingServiceHandler(s.Server(), &service{repo}); err != nil {
		log.Panic(err)
	}

	if err := s.Run(); err != nil {
		log.Panic(err)
	}

}
