package main

import (

    "encoding/json"
    "io/ioutil"
    "log"
    "os"

    pb "consignment-service/consignment"
    "go-micro.dev/v4"

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

    service := micro.NewService()
    service.Init()

    cli := pb.NewShippingServiceClient(

    //read a json file sended to server
    file := defaultFilename
    if len(os.Args) > 1{
        file = os.Args[1]
    }
    consignment, err := parseFile(file)
    if err != nil {
        log.Fatalf("Could not parse the file:%v", err)
    }

    //Test CreateConsignment function
    r, err := client.CreateConsignment(context.Background(), consignment)
    if err != nil {
        log.Fatalf("Could not create the consignment:%v", err)
    }
    log.Printf("Created:%t", r.Created)

    //Test GetConsignmentList function
    //input: empty GetListRequest
    //output:repo_list, a list of consignments
    repo_list, err_list := client.GetConsignmentList(context.Background(), &pb.GetRequest{})
    if err_list != nil {
        log.Fatalf("Could not list consignments:%v", err_list)
    }
    for _, v := range repo_list.Consignments{
        log.Println(v)
    }

}

