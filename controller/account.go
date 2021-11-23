package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_gorm_rest/service"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	var s service.Service
	p, err := s.GetAllAccount()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) Create(c *gin.Context) {
	var s service.Service
	p, err := s.CreateAccount(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.GetAccount(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.UpdateAccount(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	if err := s.DeleteAccount(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
