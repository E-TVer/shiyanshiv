function addHtml(data){
    var html2="<li><a href=\"javascript:void(0);\" onclick=\"urlClick("+"'"+data.url+"')\">" + data.title+
        "<span class=\"am-badge am-badge-success am-radius\" style='margin-left: 10px'>"+data.label+"</span>" +
        "<span class=\"am-badge am-badge-warning am-radius\" style='margin-left: 10px'>"+data.time+"</span>" +
        "</a></li>";
    $("#fy-data").append(html2);
}
function addHtmlFromHot(count,idx,data){
    var all1="<div class=\"am-u-sm-"+12/(theType*3);
    var all2="\" onclick=\"urlClickFromHot("+"'"+data.title+"')\">" +
        "    <div class=\"am-thumbnail\">" +
        "      <img src=\""+data.cover+"\" alt=\"\"/>" +
        "      <p class=\"am-thumbnail-caption\" style='padding-bottom: 0;margin-bottom: 0;'>"+data.title+"</p>" +
        "      <p class=\"am-thumbnail-caption\" style='color: lightcoral;padding-top: 0;padding-bottom: 0'>评分："+data.rate+"</p>" +
        "    </div>" +
        "  </div>";
    if(count===idx+1){
        all1=all1+" am-u-end";
    }
    if(idx%(theType*3)===0){
        console.log(idx);
        var idd="<div class=\"am-g\" id=\"fy-";
        var html3=">" +all1+all2+
            "</div>";
        html3=idd+idx+"\""+html3;
        $("#dy-list-hot").append(html3);
    }else {
        var lastId="fy-"+(idx-idx%(theType*3)).toString();
        console.log(lastId);
        $("#"+lastId).append(all1+all2);
    }
}
function addHtml2(data){
    var ss=data.needPlayer?"<span class=\"am-badge am-badge-warning am-radius\" style='margin-left: 10px'>推荐</span>":"";
    if(!data.needPlayer && data.url.indexOf("=") != -1){
        ss="<span class=\"am-badge am-badge-warning am-radius\" style='margin-left: 10px'>不推荐</span>";
    }
    var html="<li><a href=\"javascript:void(0);\" onclick=\"urlClick2("+data.needPlayer+",'"+data.url+"')\">" + data.title+
        "<span class=\"am-badge am-radius\" style='margin-left: 10px'>"+keyWord+"</span>" +
        ss +
        "</a></li>";
    $("#fy-data").append(html);
}
function urlClickFromHot(key) {
    mui("#progressbar").progressbar().show();
    keyWord=key;
    $("#search_text").attr("placeholder", keyWord);
    getJsonData('/movie/sou',keyWord);
}
function urlClick2(needPlayer, url) {
    console.log(needPlayer);
    console.log(url);
    if(needPlayer){
        window.open("/movie/player?url="+url);
    }else {
        window.open(url);
    }
}
function urlClick(url) {
    $("#fy-data").empty();
    mui("#progressbar").progressbar().show();
    mui.ajax("/movie/video",{
        dataType:'json',
        type:'post',
        timeout:100000,
        data:{url:url,source:source},
        success:function(data){
            mui("#progressbar").progressbar().hide();
            $.each(data,function(idx, obj){
                addHtml2(obj);
            });
        },
        error:function(xhr,type,errorThrown){
            mui("#progressbar").progressbar().hide();
            mui.toast("请刷新网页，加载出错:"+errorThrown,{ duration:'short', type:'div' });
        }
    });
}
function getJsonData(url,wd) {
    if($("#dy-list-hot").css("display")!=='none'){
        $("#dy-list-hot").hide();
    }
    mui.ajax(url,{
        dataType:'json',
        type:'post',
        timeout:100000,
        data:{
            source:source,
            wd:wd
        },
        success:function(data){
            //console.log(data);
            if($("#dy-list").css("display")=='none'){
                $("#dy-list").show();
            }
            mui("#progressbar").progressbar().hide();
            $.each(data,function(idx, obj){
                addHtml(obj);
            });
        },
        error:function(xhr,type,errorThrown){
            mui("#progressbar").progressbar().hide();
            mui.toast("请刷新网页，加载出错:"+errorThrown,{ duration:'short', type:'div' });
        }
    });
}
function getHot() {
    $.getJSON("/movie/hot", function(data) {
        var count=Object.keys(data.subjects).length;
        $.each(data.subjects,function(idx, obj){
            addHtmlFromHot(count,idx,obj);
        });
    });

}
function searchClick(which){
    var str = $("#search_text").val();
    if(str!==""){
        if(hasWhich!=which){//避免重复点击
            setTimeout(function () {
                hasWhich="111";
            },5000);
        }else {
            return;
        }
        hasWhich=which;
        str=encodeURI(str);
        var url="no";
        switch(which){
            case "fy":
                url="http://118.24.88.92:8080/v2/search?q="+str+"&p=1";
                break;
            case "88":
                url=(theType===1)?"http://m.88ys.tv/index.php?m=vod-search&wd="+str:"http://www.88ys.tv/index.php?m=vod-search&wd="+str;
                break;
            case "bd":
                $("#bdyKey").val(str);
                document.getElementById("fm").submit();
                break;
            case "fk":
                url="http://ifkdy.com/?q="+str+"&p=1";
                break;
            case "owl":
                url="https://www.owllook.net/search?wd="+str;
                break;
            case "bt":
                url="https://www.btmule.org/q/"+str+".html";
                break;
            case "zh":
                url="https://www.zhihu.com/search?type=content&q="+str;
                break;
            case "db":
                url=(theType===1)?"https://m.douban.com/search/?query="+str:"https://www.douban.com/search?q="+str;
                break;
            case "sina":
                url=(theType===1)?"https://m.weibo.cn/p/100103type%3D1%26q%3D"+str:"http://s.weibo.com/weibo/"+str+"&Refer=STopic_box";
                break;
            case "wx":
                url=(theType===1)?"http://weixin.sogou.com/weixinwap?type=2&query="+str:"http://weixin.sogou.com/weixin?type=2&query="+str;
                break;
            case "mu":
                url="http://music.2333.me/?name="+str+"&type=netease";
                break;
            case "yk":
                url="http://www.soku.com/m/y/video?q="+str+"&tpa=dW5pb25faWQ9MTAzOTQyXzEwMDAwMV8wMV8wMQ";
                break;
            case "bj":
                url=(theType===1)?"https://m.baidu.com/pu=sz%401321_480/s?pn=10&usm=4&word="+str+"&sa=np&ant_ct=RK7ymxDKg7%2FqgQBBBy1f%2FT2FpqsTn9ilyJZoGzSXgFIzavWjEdEHFZVKCoRuP5rp&rqid=7015555724176550070":"https://www.baidu.com/baidu?wd=" + str;
        }
        if(url!=="no"){
            window.open(url);
        }
    } else {
        $("#search_text").attr("placeholder", "请先在这里输入关键词").focus()
    }
}
function choose(){
    mui('.mui-popover').popover('toggle',document.getElementById("popover"));
}
function a() {
    var f = $("#search_text").val();
    var e = $.trim(f);
    if (e !== "") {
        if(hasWhich!="baidu"){//避免重复点击
            setTimeout(function () {
                hasWhich="111";
            },3000);
        }else {
            return;
        }
        hasWhich="baidu";
        $("#fy-data").empty();
        mui("#progressbar").progressbar().show();
        $("#search_text").attr("placeholder", e);
        keyWord=e;
        getJsonData('/movie/sou',keyWord);
    } else {
        $("#search_text").attr("placeholder", "请先在这里输入关键词").focus()
    }
}
function searchTab(ss){
    source=ss;
    if(keyWord!==""){
        $("#fy-data").empty();
        mui("#progressbar").progressbar().show();
        $("#search_text").attr("placeholder", keyWord);
        getJsonData('/movie/sou',keyWord);
    }else {
        $("#search_text").attr("placeholder", "请先在这里输入关键词").focus()
    }
}
function getSource(){
    window.open("https://github.com/qiusunshine/shiyanshiv");
}
mui("#progressbar").progressbar().show();
var keyWord=$("#keyWord").text();
var hasWhich="111";
var source="zuida";
var theType=(mui.os.ios||mui.os.android||mui.os.wechat)?1:2;
if(theType===1)$(".am-topbar-inverse").css("margin-top","10px");
mui('#quickSearch').on('tap','a',function(){
    this.click();
});
mui("#search_go").on("tap", function () {
    this.click();
});
$(document).on("keydown", function (e) {
    if (e.keyCode == "13") {
        a();
    }
});
mui('#moreSearch').on('tap','img',function(){
    searchClick(this.alt);
});
if(keyWord!==""){
    $("#search_text").attr("placeholder", keyWord);
    getJsonData('/movie/sou',keyWord);
}else {
    getHot();
    mui("#progressbar").progressbar().hide();
}
