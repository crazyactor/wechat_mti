package wechathandle

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {

	//错误字段
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`

	//access_token
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func get(url string) (*response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rtn response
	if err := json.Unmarshal(data, &rtn); err != nil {
		return nil, errors.New(fmt.Sprintf("%d %s", rtn.ErrCode, rtn.ErrMsg))
	}
	return &rtn, nil
}
