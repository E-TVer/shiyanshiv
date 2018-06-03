package main

import (
	_ "hdy/shiyanshiv/routers"
	"github.com/astaxie/beego"
	"hdy/shiyanshiv/util"
	"io/ioutil"
	"fmt"
	"time"
	"strings"
	"net/http"
)

func openBrowserFailed() {
	util.Info("打开浏览器失败！...")
	util.Info("正在尝试复制链接到剪贴板...")
	err:=util.WriteAll("127.0.0.1:9876")
	if err != nil {
		util.Info("复制链接到剪贴板失败！")
		util.Info("请使用浏览器访问右边链接---->127.0.0.1:9876")
	}else {
		util.Info("链接已复制到剪贴板！请使用浏览器打开链接---->...")
		util.Info("或者手动用浏览器访问右边的链接---->127.0.0.1:9876")
		util.Info("温馨提示---->关闭本窗口后将关闭本程序！")
		util.Info("方圆实验室V版日志记录开始<----")
	}
}
func openBrowserSuccess() {
	util.Info("成功使用浏览器打开网站")
	util.Info("也可以手动用浏览器访问右边的链接---->127.0.0.1:9876")
	util.Info("温馨提示---->关闭本窗口后将关闭本程序！")
	util.Info("方圆实验室V版日志记录开始<----")
}
func main() {
	data,er:=ioutil.ReadFile("logo.txt")
	if er==nil{
		fmt.Println(string(data))
	}
	util.Info("欢迎使用方圆实验室V版！当前版本号：2")
	util.Info("有任何问题可以加群咨询！群号码：615468609")
	util.Info("正在尝试使用浏览器访问网站...")
	util.Info("温馨提示---->请自行设置好常用浏览器为默认浏览器")
	var ti =make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(time.Minute):
				util.Info("一分钟了还没启动！请重启程序！")
				return
			case <-ti:
				err:=util.Open("http://127.0.0.1:9876")
				if err!=nil{
					util.Info("出错了--->"+err.Error())
					openBrowserFailed()
				}else {
					openBrowserSuccess()
				}
				return
			}
		}
	}()
	go func() {
		for i:= 0; i < 10; i++ {
			time.Sleep(time.Second*3)
			req,err:=http.NewRequest("GET","http://127.0.0.1:9876",nil)
			if err==nil{
				client:=http.Client{Timeout:time.Second*3}
				resp,err:=client.Do(req)
				if err==nil{
					if strings.HasPrefix(resp.Status,"200"){
						//util.Info("Status is OK")
						ti<-true
						break
					}
				}else {
					util.Info("出错："+err.Error())
					time.Sleep(time.Second*3)
				}
			}else {
				util.Info("出错："+err.Error())
				time.Sleep(time.Second*3)
			}
		}
	}()

	beego.Run()
}
