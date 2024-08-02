package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"path/filepath"

	pb "github.com/cywang4/ovn-test/kubeovn/pb"

	ovnv1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	ovnv1client "github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type server struct {
	pb.UnimplementedNetworkServiceServer
	Client *ovnv1client.Clientset
}

func NewServer(kubeconfigPath string) (*server, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build config: %w", err)
	}

	clientset, err := ovnv1client.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	return &server{Client: clientset}, nil
}

func (s *server) CreateVPC(ctx context.Context, req *pb.CreateVPCRequest) (*pb.CreateVPCResponse, error) {
	vpc := &ovnv1.Vpc{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec: ovnv1.VpcSpec{
			Namespaces: req.Namespaces,
		},
	}

	_, err := s.Client.KubeovnV1().Vpcs().Create(context.TODO(), vpc, metav1.CreateOptions{})
	if err != nil {
		return &pb.CreateVPCResponse{Status: "Failure", Error: err.Error()}, nil
	}

	return &pb.CreateVPCResponse{Status: "Success"}, nil
}

func (s *server) CreateSubnet(ctx context.Context, req *pb.CreateSubnetRequest) (*pb.CreateSubnetResponse, error) {
	subnet := &ovnv1.Subnet{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec: ovnv1.SubnetSpec{
			Vpc:       req.VpcName,
			CIDRBlock: req.CidrBlock,
			Gateway:   req.Gateway,
			Protocol:  req.Protocol,
		},
	}

	_, err := s.Client.KubeovnV1().Subnets().Create(context.TODO(), subnet, metav1.CreateOptions{})
	if err != nil {
		return &pb.CreateSubnetResponse{Status: "Failure", Error: err.Error()}, nil
	}

	return &pb.CreateSubnetResponse{Status: "Success"}, nil
}

func main() {
	kubeconfigPath := filepath.Join("/Users/walker.wang", ".kube", "ovn-test")

	s, err := NewServer(kubeconfigPath)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNetworkServiceServer(grpcServer, s)

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
