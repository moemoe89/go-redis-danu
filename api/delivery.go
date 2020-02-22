//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"github.com/moemoe89/practicing-redis-golang/api/api_struct/form"
	"github.com/moemoe89/practicing-redis-golang/api/api_struct/model"
	cons "github.com/moemoe89/practicing-redis-golang/constant"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
	log *logrus.Entry
	svc Service
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(log *logrus.Entry, svc Service) *ctrl {
	return &ctrl{log, svc}
}

func (ct *ctrl) Set(c *gin.Context) {
	req := &form.KeyValueForm{}
	if err := c.BindJSON(&req); err != nil {
		ct.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{"Oops! Something went wrong with your request"}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	expiry := req.Expiry * time.Second
	status, err := ct.svc.Set(req.Key, req.Value, expiry)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusCreated, model.NewGenericResponse(http.StatusCreated, cons.OK, []string{"Data has been saved"}, nil))
}

func (ct *ctrl) SAdd(c *gin.Context) {
	req := &form.KeyValuesForm{}
	if err := c.BindJSON(&req); err != nil {
		ct.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{"Oops! Something went wrong with your request"}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	status, err := ct.svc.SAdd(req.Key, req.Values)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusCreated, model.NewGenericResponse(http.StatusCreated, cons.OK, []string{"Data has been saved"}, nil))
}

func (ct *ctrl) HSet(c *gin.Context) {
	req := &form.KeyFieldValueForm{}
	if err := c.BindJSON(&req); err != nil {
		ct.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{"Oops! Something went wrong with your request"}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	isCreate, status, err := ct.svc.HSet(req.Key, req.Field, req.Value)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	status = http.StatusOK
	message := "Data has been updated"
	if isCreate {
		status = http.StatusCreated
		message = "Data has been created"
	}

	c.JSON(status, model.NewGenericResponse(status, cons.OK, []string{message}, nil))
}

func (ct *ctrl) Get(c *gin.Context) {
	key := c.Param("key")

	data, status, err := ct.svc.Get(key)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, data))
}

func (ct *ctrl) GetHSet(c *gin.Context) {
	key := c.Param("key")
	field := c.Param("field")

	data, status, err := ct.svc.GetHSet(key, field)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, data))
}

func (ct *ctrl) GetSAdd(c *gin.Context) {
	key := c.Param("key")

	datas, status, err := ct.svc.GetSAdd(key)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, datas))
}

func (ct *ctrl) Delete(c *gin.Context) {
	key := c.Param("key")

	status, err := ct.svc.Delete(key)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been deleted"}, nil))
}