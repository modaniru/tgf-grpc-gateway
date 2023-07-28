package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/modaniru/grpc-gate-way/src/client"
	"github.com/modaniru/grpc-gate-way/src/controller"
	"google.golang.org/grpc"
)

func main(){
	// .env
	_ = godotenv.Load()
	//conns
	tgfConn, err := grpc.Dial(os.Getenv("TGF_HOST"), grpc.WithInsecure())
	if err != nil{
		log.Fatal(err.Error())
	}
	defer tgfConn.Close()
	imageServiceConn, err := grpc.Dial(os.Getenv("IMAGE_HOST"), grpc.WithInsecure())
	if err != nil{
		log.Fatal(err.Error())
	}
	defer imageServiceConn.Close()
	//clients
	tgf := client.NewTgfClient(tgfConn)
	image := client.NewImageService(imageServiceConn)
	server := controller.NewServer(gin.Default(), tgf, image)
	server.Start("8080")
}