package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/coredns/coredns/plugin/etcd/msg"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/BedivereZero/nekodns/api/core/v1alpha1"
)

func (s *Server) CreateRecord(c *gin.Context) {
	record := new(v1alpha1.Record)
	if c.Bind(record) != nil {
		return
	}

	b, err := json.Marshal(msg.Service{Host: record.Content})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	key := msg.Path(record.Name, DefaultEtcdPrefix)

	if _, err := s.Client.Put(c, key, string(b)); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (s *Server) DeleteRecord(c *gin.Context) {
	key := msg.Path(c.Param("name"), DefaultEtcdPrefix)

	if _, err := s.Client.Delete(c, key); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

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
	ms := new(msg.Service)
	if err := json.Unmarshal(kv.Value, ms); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	record := v1alpha1.Record{
		Name:    c.Param("name"),
		Content: ms.Host,
	}

	c.JSON(http.StatusOK, record)
}

func (s *Server) ListRecord(c *gin.Context) {
	rsp, err := s.Client.Get(c, DefaultEtcdPrefix, clientv3.WithPrefix())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	recordList := new(v1alpha1.RecordList)

	for _, kv := range rsp.Kvs {
		record := v1alpha1.Record{
			Name:    msg.Domain(string(kv.Key)),
			Content: string(kv.Value),
		}
		recordList.Items = append(recordList.Items, record)
	}

	c.JSON(http.StatusOK, recordList)
}
