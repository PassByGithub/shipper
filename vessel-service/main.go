package main

import (
	"context"
	"log"

	pb "shipper/vessel-service/vessel"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

type Repository struct {
	vessels []*pb.Vessel
}

//Database interface
type IRepository interface {
	Update(*pb.Vessel) (*pb.Vessel, error)
}

func (repo *Repository) Update(ves *pb.Vessel, capa int, weight int) (*pb.Vessel, error) {
	ves.Capacity = ves.Capacity + capa
	ves.Weight = ves.Weight + weight
	return ves, nil
}

//Service etcd interface
type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	for _, vessel := range s.repo.vessels {
		if in.Capacity < vessel.Capacity && in.MaxWeight < vessel.MaxWeight {
			s.repo.update(vessel, in.Capacity, in.MaxWeight)
			return vessel, nil
		}
	}
	return nil, error.New("No vessel found by that specification")
}

func main() {
	newvessels := []*pb.Vessel{
		&pb.Veesel{Id: "vessel001", Capacity: 500, Max_weight: 20000, Name: "FirstVessel"},
	}
	repo := &Repository{newvessels}

	s := micro.NewService(micro.Name("shippy.service.vessel"))
	s.Init()

	if err := pb.RegisterVesselServiceHandler(s.Server(), &service{repo}); err != nil {
		log.Panic(err)
	}

	if err := s.Run(); err != nil {
		log.Panic(err)
	}
}
