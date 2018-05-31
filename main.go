package main

import (
	_ "hdy/shiyanshiv/routers"
	"github.com/astaxie/beego"
	"hdy/shiyanshiv/util"
	"io/ioutil"
	"fmt"
)

func main() {
	data,er:=ioutil.ReadFile("logo.txt")
	if er==nil{
		fmt.Println(string(data))
	}
	util.Info("欢迎使用方圆实验室V版！当前版本号：1")
	util.Info("有任何问题可以加群咨询！群号码：615468609")
	util.Info("正在复制链接到剪贴板...")
	err:=util.WriteAll("127.0.0.1:9876")
	if err != nil {
		util.Info("尝试复制链接到剪贴板失败！")
		util.Info("请使用浏览器访问右边链接---->127.0.0.1:9876")
	}else {
		util.Info("链接已复制到剪贴板！请使用浏览器打开链接---->...")
		util.Info("或者手动用浏览器访问右边的链接---->127.0.0.1:9876")
		util.Info("温馨提示---->关闭本窗口后将关闭本程序！")
		util.Info("方圆实验室V版日志记录开始<----")
	}
	beego.Run()
}
