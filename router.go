package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/smtc/glog"
)

/**
查询所有的账号
创建人:邵炜
创建时间:2016年7月29日15:55:06
输入参数::gin对象
输出参数:无
数据传递由gin进行
 */
func selectAccessFun(c *gin.Context) {
	var (
		urlst map[string]string
		dataStr string
	)
	dataObj,_:=json.Marshal(urls)
	json.Unmarshal(dataObj,&urlst)

	for key, value := range urlst {
		dataStr+=fmt.Sprintf("省份: %s \r\n %s \r\n\r\n\r\n\r\n",key,selectAccess(value))
	}
	glog.Info("selectAccessFun success! requestInfo: %s \n",userReqInfo(c.Request))
	c.String(http.StatusOK,dataStr)
}

func userReqInfo(req *http.Request) (info string) {
	info += fmt.Sprintf("ipaddr: %s user-agent: %s referer: %s",
		req.RemoteAddr, req.UserAgent(), req.Referer())
	return info
}