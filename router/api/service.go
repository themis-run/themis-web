package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.themis.run/themis-web/handle"
	"go.themis.run/themis-web/model"
)

type ServiceAPI struct {
	handler *handle.ServiceHandler
}

func (s *ServiceAPI) FindAllService(c *gin.Context) {
	services, err := s.handler.FindAllService()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, model.Success(
		model.ServiceResponse{
			ServiceList:  services,
			ServiceCount: len(services),
		}))
}

func (s *ServiceAPI) FindInstance(c *gin.Context) {
	serviceName := c.Param("serviceName")
	instances, err := s.handler.FindInstance(serviceName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, model.Success(
		model.InstancesResponse{
			InstanceList: instances,
			Count:        len(instances),
		}))
}
