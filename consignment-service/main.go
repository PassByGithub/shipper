package main

import (

    "log"
    "net"

    pb "shipper/consignment-service/consignment"
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
//Define an interface for Repository towards local variable/database
type IRepository interface  {
    Create(*pb.Consignment) (*pb.Consignment, error)
    GetList() []*pb.Consignment
}
//Impletement the Create method towards lcoal varialble/database
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error){
    updated := append(repo.consignments, consignment)
    repo.consignments = updated
    return consignment, nil
}
//Implet the GetList Method towards local variable/database
func (repo *Repository) GetList() []*pb.Consignment{
    return repo.consignments
}

//service Impletement all the methods defined in .protoc, then service will get theses output to database
type service struct{
    repo IRepository
    *pb.UnimplementedShippingServiceServer
}

//service.CreateConsignment:parse the context from client.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error){
    //save the consignment to the database;
    consignment, err := s.repo.Create(req)
    if err != nil {
        return nil, err
    }
    //return Response and defined data structure
    return &pb.Response{Created: true, Consignment: consignment}, nil
}
//service.GetConsignmentList:
func (s *service) GetConsignmentList(ctx context.Context, req *pb.GetRequest) (*pb.Response, error){
    //get the consignment list from database
    consignments := s.repo.GetList()
    //return Response and defined data structure to the client
    return &pb.Response{Consignments: consignments}, nil
}
func main() {

    repo := &Repository{}
    unimplete := new(pb.UnimplementedShippingServiceServer)

    // configure the gRPC server
    lis, err := net.Listen("tcp", port)
    if err != nil{
        log.Fatalf("failed to listen:%v", err)
    }
    //register a gRPC server on gRPC server
    s := grpc.NewServer()

    //register our microservice on the grpc server
    pb.RegisterShippingServiceServer(s, &service{repo, unimplete})

    //reflection registeration
    reflection.Register(s)
    if err := s.Serve(lis); err != nil{
        log.Fatalf("failed to serve:%v", err)
    }
}

