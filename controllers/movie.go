package controllers

import (
	"github.com/astaxie/beego"
	"hdy/shiyanshiv/models"
	"strings"
)

type MovieController struct {
	beego.Controller
}
var searchResults = make(chan []models.Movie,2)

func (c *MovieController)GetHtml()  {
	c.Data["KeyWord"]=c.GetString("wd")
	c.TplName="searchResult_v2.html"
}
func (c *MovieController)Player()  {
	c.Data["MovieUrl"]=c.GetString("url")
	c.TplName="player.html"
}
func (c *MovieController)Search()  {
	ss:=c.GetString("source")
	if ss!=""&&ss!="zuida"{
		c.Data["json"]=models.GetMovieFromEight(c.GetString("wd"))
	}else {
		c.Data["json"]=models.GetMovies(c.GetString("wd"))
	}
	c.ServeJSON()
}
func (c *MovieController)Video()  {
	url:=c.GetString("url")
	ss:=c.GetString("source")
	if ss!=""&&ss!="zuida"{
		c.Data["json"]=models.GetVideoFromEight(url)
	}else {
		c.Data["json"]=models.GetVideoes(url)
	}
	c.ServeJSON()
}
func (c *MovieController)Search2()  {
	go func() {
		searchResults<-models.GetMovieFromEight(c.GetString("wd"))
	}()
	go func() {
		searchResults<-models.GetMovies(c.GetString("wd"))
	}()
	sss:=<-searchResults
	sss=append(sss,<-searchResults...)
	c.Data["json"]=sss
	c.ServeJSON()
}
func (c *MovieController)Video2()  {
	url:=c.GetString("url")
	if strings.HasPrefix(url,"http://www.88ysw"){
		c.Data["json"]=models.GetVideoFromEight(url)
	}else {
		c.Data["json"]=models.GetVideoes(url)
	}
	c.ServeJSON()
}
func (c *MovieController)CC()  {
	updateMessage:=models.UpdateMessage{ThisVersion:1,NeedUpdate:false}
	u:=models.GetUpdate()
	//fmt.Println(u)
	if u.V>updateMessage.ThisVersion{
		updateMessage.UpdateUrl=u.Url
		updateMessage.NeedUpdate=true
	}
	c.Data["json"]=updateMessage
	c.ServeJSON()
}
func(c *MovieController) GetHot()  {
	c.Data["json"]=models.GetHot()
	c.ServeJSON()
}