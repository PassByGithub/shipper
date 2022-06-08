package main


import(

	"fmt"
	"log"

	pb "shipper/vessel-service/vessel"

	"go-micro.dev/4"

)

type Repository struct {
	vessels []*pb.Vessel
}

type IRepository interface {
	Update(*pb.Vessel)(*pb.Vessel, error)
}



type service struct {
	repo Repository
}

type (s *service) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error){
	req 
}
