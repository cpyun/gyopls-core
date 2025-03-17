package api

import (
	"fmt"
	"net/http"

	"github.com/cpyun/gyopls-core/logger"
	"github.com/cpyun/gyopls-core/sdk/pkg"
	"github.com/cpyun/gyopls-core/sdk/pkg/response"
	"github.com/cpyun/gyopls-core/sdk/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct {
	Context *gin.Context
	Logger  *logger.Logger
	Orm     *gorm.DB
	Errors  error
	Msg     string
}

// AddError 添加错误
func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Errors = fmt.Errorf("%v; %w", e.Error, err)
	}
	//return e.Errors
}

// MakeContext 设置context
func (e *Api) MakeContext(ctx *gin.Context) *Api {
	e.Context = ctx
	e.Logger = GetRequestLogger(ctx)
	return e
}

// MakeOrm 设置Orm DB
func (e *Api) MakeOrm() *Api {
	var err error
	db, err := pkg.GetOrm(e.Context)
	if err != nil {
		e.Logger.Error("数据库连接获取失败", http.StatusInternalServerError, err.Error())
		e.AddError(err)
	}
	e.Orm = db
	return e
}

// GetOrm 获取Orm DB
func (e Api) GetOrm() *gorm.DB {
	db, err := pkg.GetOrm(e.Context)
	if err != nil {
		e.Logger.Error("数据库连接获取失败", http.StatusInternalServerError, err)
	}
	return db
}

// MakeService 设置服务Log、Orm
func (e *Api) MakeService(c *service.Service) *Api {
	c.Log = e.Logger
	c.Orm = e.Orm
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, msg string) {
	response.Error(e.Context, code, msg)
}

// Success 成功数据处理
func (e Api) Success(msg string, data any) {
	response.Success(e.Context, msg, data)
}

// Custom 自定义内容响应
func (e Api) Custom(httpCode int, jsonObj any) {
	response.Custom(e.Context, httpCode, jsonObj)
}
