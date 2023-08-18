package main

import (
	"github.com/gogf/gf/encoding/gyaml"
	"hdnslog/Core"
	"hdnslog/Dns"
	"hdnslog/Http"
	"log"
	"os"
)

func main() {
	ConfigBody, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Println("读取文件失败：", err)
		return
	}
	err = gyaml.DecodeTo(ConfigBody, &Core.Config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	go Dns.ListingDnsServer()
	Http.ListingHttpManagementServer()
}
