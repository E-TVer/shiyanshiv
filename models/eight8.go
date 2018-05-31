package models

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/robertkrimen/otto"
	"hdy/shiyanshiv/util"
)
type PlayerDesc struct {
	Name string
	From string
	Note string
	Urls string
	Server string
}
type PlayerDescS struct {
	Name    string
	Froms   []string
	Notes   []string
	Urls    []string
	Servers []string
}
type Source struct {
	VideoName string   `json:"video_name"`
	Players   []Player `json:"players"`
}
type Player struct {
	PlayerName string    `json:"player_name"`
	Chapters   []Chapter `json:"chapters"`
}
type Chapter struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}
var notifyMessage =make(chan int)
var playSource  Source

func GetMovieFromEight(key string)(movies []Movie) {
	doc,err:=goquery.NewDocument("http://www.88ysw.tv/index.php?m=vod-search&wd="+key)
	if err!=nil{
		util.Info("出现错误："+err.Error())
	}else {
		s:=doc.Find(".main").Find(".index-area").Find("li")
		s.Each(func(i int, selection *goquery.Selection) {
			movie:=Movie{}
			movie.Title,_=selection.Find("a").Attr("title")
			movie.Label="88影视"
			movie.Url,_=selection.Find("a").Attr("href")
			mm:=strings.Split(movie.Url,"/")
			id:=strings.Replace(mm[len(mm)-1],".html","",1)
			movie.Url="http://www.88ysw.tv/vod-play-id-"+id+"-src-1-num-1.html"
			movie.Time=selection.Find("i").Text()
			movies=append(movies, movie)
		})
	}
	return
}
func GetVideoFromEight(url string) (videoes []Video) {
	s:=getSourcesFromUrl(url)
	for _, player := range s.Players {
		for _, chapter := range player.Chapters {
			video:=Video{Title:chapter.Title,Url:chapter.Url,NeedPlayer:false}
			if strings.HasSuffix(chapter.Url,".m3u8"){
				video.NeedPlayer=true
			}else {
				v:=[]string{"sohu","youku","qiyi","qq.com","mgtv","le.com"}
				needAdd:=false
				for _, value := range v {
					if strings.Contains(video.Url,value){
						needAdd=true
						break
					}
				}
				if needAdd {
					video.Url="http://goudidiao.com/?url="+video.Url
				}
			}
			videoes=append(videoes,video)
		}
	}
	return
}
func getSourcesFromUrl(sourceUrl string) Source {
	go func() {
		util.Info("开始从网络获取网页内容...")
		doc,err:=goquery.NewDocument(sourceUrl)
		if err!=nil{
			util.Info("出现错误："+err.Error())
			notifyMessage<-1
		}else {
			doc.Find("script").EachWithBreak(func(i int, selection *goquery.Selection) bool {
				if strings.HasPrefix(selection.Text(),"var mac_flag"){
					util.Info("已经获取到关键信息")
					getPlayUrl(selection.Text())
					return false
				}
				return true
			})
		}
	}()
	<-notifyMessage
	return playSource
}
func getPlayUrl(str string) {
	util.Info("正在解析网页内容...")
	vm := otto.New()
	vm.Run(str)
	playerDesc :=PlayerDesc{}
	if value, err := vm.Get("mac_name"); err == nil {
		if valueStr, err := value.ToString(); err == nil {
			playerDesc.Name=valueStr
		}
	}
	if value, err := vm.Get("mac_from"); err == nil {
		if valueStr, err := value.ToString(); err == nil {
			playerDesc.From=valueStr
		}
	}
	if value, err := vm.Get("mac_server"); err == nil {
		if valueStr, err := value.ToString(); err == nil {
			playerDesc.Server=valueStr
		}
	}
	if value, err := vm.Get("mac_note"); err == nil {
		if valueStr, err := value.ToString(); err == nil {
			playerDesc.Note=valueStr
		}
	}
	if value, err := vm.Get("mac_url"); err == nil {
		if valueStr, err := value.ToString(); err == nil {
			playerDesc.Urls=valueStr
		}
	}
	playerDescS :=PlayerDescS{}
	playerDescS.Name=playerDesc.Name
	playerDescS.Urls=strings.Split(playerDesc.Urls,"$$$")
	playerDescS.Notes=strings.Split(playerDesc.Note,"$$$")
	playerDescS.Froms=strings.Split(playerDesc.From,"$$$")
	playerDescS.Servers=strings.Split(playerDesc.Server,"$$$")
	sources:=Source{VideoName:playerDescS.Name}
	for i, url := range playerDescS.Urls {
		player:=Player{PlayerName:playerDescS.Froms[i]}
		uu:=strings.Split(url,"#")
		for _, value := range uu {
			vs:=strings.Split(value,"$")
			if strings.HasPrefix(vs[1],"http"){
				chapter:=Chapter{}
				chapter.Title=vs[0]
				chapter.Url=vs[1]
				player.Chapters=append(player.Chapters,chapter)
			}
		}
		sources.Players=append(sources.Players,player)
	}
	playSource=sources
	util.Info("已经解析完所有链接")
	notifyMessage<-1
}
