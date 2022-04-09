package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.themis.run/themis-web/handle"
	"go.themis.run/themis-web/model"
)

type KeyValueAPI struct {
	handler *handle.KeyValueHandler
}

func (k *KeyValueAPI) ListKV(c *gin.Context) {
	res, err := k.handler.ListKV()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, model.Success(res))
}

func (k *KeyValueAPI) SetKV(c *gin.Context) {
	kv := model.KeyValue{}
	if err := c.ShouldBindJSON(&kv); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := k.handler.SetKV(kv); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, model.Success(nil))
}
