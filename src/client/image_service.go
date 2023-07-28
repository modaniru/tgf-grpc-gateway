package client

import (
	pkg "github.com/modaniru/grpc-gate-way/pkg/proto"
	"google.golang.org/grpc"
)

type ImageService struct{
	pkg.ImageServiceClient
}

func NewImageService(client *grpc.ClientConn) *ImageService{
	return &ImageService{ImageServiceClient: pkg.NewImageServiceClient(client)}
}