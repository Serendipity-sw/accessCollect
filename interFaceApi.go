package main

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/smtc/glog"
	"encoding/json"
)

const (
	SUCCESS  ="00000"
)

type requestInfo struct {
	ResultCode string `json:"ResiltCode"`//返回码
	Message string `json:"Message"` //内容
}

/**
获取账号
创建人:邵炜
创建时间:2016年7月29日14:55:47
 */
func selectAccess(url string) (text string) {
	text="数据读取失败"
	url=fmt.Sprintf("%s/unitSelectAccess",url)
	request, err := http.NewRequest("POST", url, bytes.NewReader([]byte("")))
	if err != nil {
		glog.Error("selectAccess newRequest error! url: %s err: %s \n",url,err.Error())
		return
	}
	request.Header.Set("authorize", authorize)
	httpClient := http.Client{}

	respone, err := httpClient.Do(request)
	if err != nil {
		glog.Error("selectAccess http send error! url: %s err: %s \n",url,err.Error())
		return
	}
	defer respone.Body.Close()
	responeJsonStr, err := ioutil.ReadAll(respone.Body)
	if err != nil {
		glog.Error("selectAccess data read error! url: %s err: %s \n",url,err.Error())
		return
	}
	var obj requestInfo
	err=json.Unmarshal(responeJsonStr,&obj)
	if err != nil {
		glog.Error("selectAccess json unmarshal error! url: %s jsonStr: %s error: %s \n",url,string(responeJsonStr),err.Error())
		return
	}
	text=obj.Message
	return
}
