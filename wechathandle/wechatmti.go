package wechathandle

import (
	"config"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

var conf *config.Config = config.Instance()

func WechatHandle(w http.ResponseWriter, r *http.Request) {

	if !checkSignature(r) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		log.Println("校验消息真实性失败")
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r.FormValue("echostr")))
		log.Println("检验消息真实性成功")
		return
	}
	log.Println("接收到post数据，准备解析")
	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	log.Printf("%s", result)
	return

}

func checkSignature(r *http.Request) bool {

	r.ParseForm()
	token := fmt.Sprintf("%s", conf.ConMap["WeChat.TOKEN"])

	var signature string = r.FormValue("signature")
	var timestamp string = r.FormValue("timestamp")
	var nonce string = r.FormValue("nonce")

	tmps := []string{token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := str2sha1(tmpStr)
	return tmp == signature
}
func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
