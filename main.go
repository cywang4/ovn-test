package main

import (
	"context"
	"log"
	"net"

	"ovn-test/kubeovn/pb"

	"google.golang.org/grpc"
)

// ... (保留 NetworkManager 结构体以及所有方法)

type server struct {
	pb.UnimplementedNetworkServiceServer
	networkManager *NetworkManager
}

func (s *server) CreateVPC(ctx context.Context, req *pb.CreateVPCRequest) (*pb.CreateVPCResponse, error) {
	err := s.networkManager.CreateVPC(req.Name, req.Namespaces)
	if err != nil {
		return nil, err
	}
	return &pb.CreateVPCResponse{Message: "VPC created successfully."}, nil
}

func (s *server) CreateSubnet(ctx context.Context, req *pb.CreateSubnetRequest) (*pb.CreateSubnetResponse, error) {
	err := s.networkManager.CreateSubnet(req.Name, req.VpcName, req.CidrBlock, req.Gateway, req.Protocol)
	if err != nil {
		return nil, err
	}
	return &pb.CreateSubnetResponse{Message: "Subnet created successfully."}, nil
}

func main() {
	// ... (保留 NewNetworkManager 函数)

	networkManager, err := NewNetworkManager(kubeconfigPath)
	if err != nil {
		log.Fatalf("Error creating Network Manager: %v", err)
	}

	s := &server{networkManager: networkManager}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNetworkServiceServer(grpcServer, s)

	log.Printf("Starting gRPC server on %s", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
