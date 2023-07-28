package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pkg "github.com/modaniru/grpc-gate-way/pkg/proto"
)

func (s *Server) GetGeneralFollows(c *gin.Context){
	body := new(Nicknames)
	err := c.ShouldBindJSON(body)
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := s.tgf.GetGeneralFollows(context.Background(), &pkg.GetTGFRequest{Usernames: body.Nicknames})
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, response)
}

func (s *Server) GetImageGeneralFollows(c *gin.Context){
	body := new(Nicknames)
	err := c.ShouldBindJSON(body)
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := s.imageService.GetGeneralFollowsImage(context.Background(), &pkg.GeneralFollowsImageRequest{Nicknames: body.Nicknames})
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, response)
}