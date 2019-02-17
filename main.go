package main

import (
	"encoding/json"
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/rubinus/zgo/zgo"
)

func main() {
	//测试工具函数
	ip := zgo.Utils.GetIntranetIp()
	fmt.Println(ip)
	ts := zgo.Utils.GetTimestamp(13)
	fmt.Printf("timestamp: %d\n", ts)
	fmt.Println(zgo.Utils.GetUUIDV4())

	//测试zoneinfo信息
	//b, _ := zgo.ZoneInfo.GetTZData("Asia/Shanghai")
	//fmt.Printf("%v", b)

	//测试读取mongodb数据
	args := make(map[string]interface{})
	args["db"] = "local"
	args["collection"] = "startup_log"
	args["query"] = bson.M{"hostname": "154f6a1b60b4"}
	result, err := zgo.Mongo.Get(args)
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes), err, "---from mongo---")

}
