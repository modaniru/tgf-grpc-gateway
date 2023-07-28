package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/modaniru/grpc-gate-way/src/client"
)

type Server struct{
	router *gin.Engine
	tgf *client.TgfClient
	imageService *client.ImageService
}

func NewServer(router *gin.Engine, tgf *client.TgfClient, imageService *client.ImageService) *Server{
	return &Server{router: router, tgf: tgf, imageService: imageService}
}

func (s *Server) Start(port string){
	tgfGroup := s.router.Group("/tgf")
	{
		tgfGroup.GET("/general-follows", s.GetGeneralFollows)
	}
	imageServiceGroup := s.router.Group("/image-service")
	{
		imageServiceGroup.GET("/general-follows", s.GetImageGeneralFollows)
	}
	s.router.Run(":" + port)
}