package api

import (
	"encoding/json"
	"github.com/emersion/go-imap/client"
	"html/template"
	"log"
	"net/http"
	"os"
	"xmail/conf"
	"xmail/model"
	"xmail/utils"
)

var (
	//Users用户的登录信息，全局使用
	Users []model.User
	//当前登录用户
	CurrentUser model.User
	// Locale用于放置模板中要替换的信息,全局使用
	Locale *map[string]string
	//邮箱客户端
	Client *client.Client
)

func init() {
	Locale = conf.Init()
	f, err := os.Open("conf/user.json")
	if err != nil {
		return
	}
	json.NewDecoder(f).Decode(&Users)
	if len(Users) != 0{
		CurrentUser = Users[0]
		var err error
		if Client, err = ConnectServer(CurrentUser.Email, CurrentUser.Password, CurrentUser.MailHost, CurrentUser.Port); err != nil{
			log.Println((*Locale)["loginError"])
			Users = utils.Remove(Users, CurrentUser)
			CurrentUser= model.User{}
			go RestoreUser()
		}
	}
}

//GoHome 跳转到首页
func GoHome(w http.ResponseWriter, r *http.Request) {
	//当前登录用户为空，即未登录
	var user model.User
	r.URL.Path = "/index"
	if CurrentUser == user {
		r.URL.Path = "/login"
	}
	GoPage(w, r)
}

//GoPage 遵照约定，加载指定模板，模板名字为url
func GoPage(w http.ResponseWriter, r *http.Request){
	path := r.URL.String()
	t, err := template.ParseFiles("template"+path+".html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, Locale)
}

//Shutdown 关闭程序，会保留登录信息
func Shutdown(w http.ResponseWriter, r *http.Request){
	os.Exit(0)
}
