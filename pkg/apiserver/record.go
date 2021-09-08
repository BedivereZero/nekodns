package apiserver

import (
	"net/http"

	"github.com/coredns/coredns/plugin/etcd/msg"
	"github.com/gin-gonic/gin"

	"github.com/BedivereZero/nekodns/api/core/v1alpha1"
)

func (s *Server) CreateRecord(c *gin.Context) {
	record := v1alpha1.Record{}
	if c.Bind(record) != nil {
		return
	}

	key := msg.Path(record.Name, DefaultEtcdPrefix)

	if _, err := s.Client.Put(c, key, record.Content); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (s *Server) DeleteRecord(c *gin.Context) {}

func (s *Server) GetRecord(c *gin.Context) {
	key := msg.Path(c.Param("name"), DefaultEtcdPrefix)

	resp, err := s.Client.Get(c, key)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if resp.Count == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	kv := resp.Kvs[0]
	record := v1alpha1.Record{
		Name:    c.Param("name"),
		Content: string(kv.Value),
	}

	c.JSON(http.StatusOK, record)
}
