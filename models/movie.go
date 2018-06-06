package models

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"hdy/shiyanshiv/util"
)

type Movie struct {
	Title string `json:"title"`
	Label string `json:"label"`
	Url   string `json:"url"`
	Time  string `json:"time"`
}
type Video struct {
	Title      string `json:"title"`
	Url        string `json:"url"`
	NeedPlayer bool   `json:"needPlayer"`
}
type UpdateMessage struct {
	ThisVersion int    `json:"now_version"`
	UpdateUrl   string `json:"update_url"`
	NeedUpdate  bool   `json:"need_update"`
}
type UpdateDetail struct {
	V   int    `json:"v"`
	Url string `json:"url"`
}
func GetMovies(key string)(movies []Movie) {
	doc,err:=goquery.NewDocument("http://www.zuixinzy.com/index.php?m=vod-search&wd="+key)
	if err!=nil{
		util.Info("出现错误："+err.Error())
	}else {
		s:=doc.Find(".xing_vb").Find("ul")
		s.Each(func(i int, selection *goquery.Selection) {
			if i!=0&&i!=-1&&i!=s.Length()-1{
				movie:=Movie{}
				movie.Title=selection.Find("a").Text()
				movie.Label=selection.Find("span").Eq(2).Text()
				movie.Url,_=selection.Find("a").Attr("href")
				movie.Url="http://www.zuixinzy.com"+movie.Url
				movie.Time=selection.Find("span").Eq(3).Text()
				movies=append(movies, movie)
			}
		})
	}
	return
}

func GetVideoes(url string)(videoes []Video) {
	doc,err:=goquery.NewDocument(url)
	if err!=nil{
		fmt.Println(err)
	}else {
		s:=doc.Find(".vodplayinfo").Eq(2).Find("li")
		s.Each(func(i int, selection *goquery.Selection) {
			video:=Video{}
			ss:=strings.Split(selection.Text(),"$")
			video.Title=ss[0]
			video.Url=ss[1]
			video.NeedPlayer=strings.Contains(ss[1],".m3u8")
			videoes=append(videoes, video)
		})
	}
	return
}
func GetUpdate() (u UpdateDetail) {
	resp,err:=http.Get("https://coding.net/u/Gold2River/p/NetRaw/git/raw/master/shiyanshiv_u.txt")
	if err!=nil{
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err=json.Unmarshal(body,&u)
	if err!=nil{
		return
	}
	return
}
func GetHot() (u interface{}) {
	resp,err:=http.Get("https://movie.douban.com/j/search_subjects?type=movie&tag=%E7%83%AD%E9%97%A8&sort=recommend&page_limit=20&page_start=0")
	if err!=nil{
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err=json.Unmarshal(body,&u)
	if err!=nil{
		return
	}
	return
}