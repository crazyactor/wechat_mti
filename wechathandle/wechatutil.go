package wechathandle

import (
	"fmt"
	"log"
	"net/http"
)

func UtilHandle(w http.ResponseWriter, r *http.Request) {
	appid := fmt.Sprintf("%s", conf.ConMap["WeChat.APPID"])
	appsercret := fmt.Sprintf("%s", conf.ConMap["WeChat.APPSECRET"])
	access_token := AccessToken{AppId: appid, AppSercret: appsercret}
	s, err := access_token.Fresh()
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		log.Printf("获取access_token出错，错误为%s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
	log.Println("获取access_token成功", s)
	return
}
