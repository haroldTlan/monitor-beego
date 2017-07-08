package main

import (
	"flag"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "last/routers"
)

var (
	nsq_ip         = beego.AppConfig.String("nsq") + ":" + beego.AppConfig.String("nsq_pub_port")
	nsqdAddr       = flag.String("nsqd", nsq_ip, "nsqd http address")
	maxInFlight    = flag.Int("max-in-flight", 200, "Maximum amount of messages in flight to consume")
	PingExport, _  = beego.AppConfig.Int("pExport")
	PingStorage, _ = beego.AppConfig.Int("pStorage")
	registerIp     = beego.AppConfig.String("registerDB")
	registerPort   = beego.AppConfig.String("registerPort")
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:passwd@/last?charset=utf8&loc=Local")

	logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/facial.log","daily":false,"maxdays":365,"level":3}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetStaticPath("/targets", "/home/targets")
	beego.Run()
}
