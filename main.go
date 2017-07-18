package main

import (
	"log"
	"net/http"
	"wechathandle"
)

func main() {
	// test := &testpb.MyMsg{
	// 	Name:  "小三",
	// 	Age:   14,
	// 	Phone: "IOS",
	// }
	// data, err := proto.Marshal(test)
	// if err != nil {
	// 	log.Fatal("marshaling error: ", err)
	// }
	// newTest := &testpb.MyMsg{}
	// err = proto.Unmarshal(data, newTest)
	// if err != nil {
	// 	log.Fatal("unmarshaling error: ", err)
	// }
	// log.Printf("NAME:%s;AGE:%d;PHONE:%s;", newTest.Name, newTest.Age, newTest.Phone)

	http.HandleFunc("/wechat_mti", wechathandle.WechatHandle)
	http.HandleFunc("/wechat_util", wechathandle.UtilHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
