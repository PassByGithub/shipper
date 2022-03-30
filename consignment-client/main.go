package main

import (

    "encoding/json"
    "io/ioutil"
    "log"
    "os"

    pb "consignment-service/consignment"
    "golang.org/x/net/context"
    "google.golang.org/grpc"

)

//certain port
const(
    address = "localhost:50051"
    defaultFilename = "consignment.json"
)

//Define a local variable as Repository

func parseFile(file string) (*pb.Consignment, error){
    var consignment *pb.Consignment

    data, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, err
    }
    json.Unmarshal(data, &consignment)
    return consignment, err

}


func main() {

    // configure the gRPC connetcion
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil{
        log.Fatalf("failed to connect:%v", err)
    }
    //set up a gRPC server
    defer conn.Close()
    client := pb.NewShippingServiceClient(conn)

    file := defaultFilename
    if len(os.Args) > 1{
        file = os.Args[1]
    }

    consignment, err := parseFile(file)

    if err != nil {
        log.Fatalf("Could not parse the file:%v", err)
    }

    r, err := client.CreateConsignment(context.Background(), consignment)
    if err != nil {
        log.Fatalf("Could not create the consignment:%v", err)
    }
    log.Printf("Created:%t", r.Created)

}

