package main

import (

    "log"
    "net"

    pb "consignment-service/consignment"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

)

//certain port
const(
    port = ":50051"
)

//Define a local variable as Repository
type Repository struct{
    consignments []*pb.Consignment
}

//Define an interface for Repository
type IRepository interface  {
    Create(*pb.Consignment) (*pb.Consignment, error)
}

//Impletement the Create method
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error){
    updated := append(repo.consignments, consignment)
    repo.consignments = updated
    return consignment, nil
}

//service Impletement all the methods defined in .protoc
type service struct{
    repo IRepository
    *pb.UnimplementedShippingServiceServer
}

//service.CreateConsignment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error){
    //save the consignment
    consignment, err := s.repo.Create(req)
    if err != nil {
        return nil, err
    }

    //return Response and defined data structure
    return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignmentList(ctx context.Context, req *pb.GetRequest) (*pb.Response, error){
    //get the consignments
    Consignments, err := s.repo.Create(req)
    if err != nil {
        return nil, err
    }

    //return Response and defined data structure
    return &pb.Response{Created: true, Consignment: consignment}, nil
}
func main() {

    repo := &Repository{}
    unimplete := new(pb.UnimplementedShippingServiceServer)

    // configure the gRPC server
    lis, err := net.Listen("tcp", port)
    if err != nil{
        log.Fatalf("failed to listen:%v", err)
    }
    //set up a gRPC server
    s := grpc.NewServer()

    //register our microservice on the grpc server
    pb.RegisterShippingServiceServer(s, &service{repo, unimplete})

    //reflection registeration
    reflection.Register(s)
    if err := s.Serve(lis); err != nil{
        log.Fatalf("failed to serve:%v", err)
    }
}

