package client

import (
	pkg "github.com/modaniru/grpc-gate-way/pkg/proto"
	"google.golang.org/grpc"
)

type TgfClient struct{
	pkg.TwitchGeneralFollowsClient
}

func NewTgfClient(client *grpc.ClientConn) *TgfClient{
	return &TgfClient{TwitchGeneralFollowsClient: pkg.NewTwitchGeneralFollowsClient(client)}
}