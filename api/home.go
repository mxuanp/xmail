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

var Client *client.Client

func init() {
	//加载登录且未注销的用户信息
	f, err := os.Open("conf/user.json")
	if err != nil {
		return
	}
	json.NewDecoder(f).Decode(&conf.Context.Users)
	if len(conf.Context.Users) != 0 {
		conf.Context.CurrentUser = conf.Context.Users[0]
		var err error
		if Client, err = ConnectServer(conf.Context.CurrentUser.Email, conf.Context.CurrentUser.Password, conf.Context.CurrentUser.MailHost, conf.Context.CurrentUser.Port); err != nil {
			log.Println(conf.Context.Locale["loginError"])
			conf.Context.Users = utils.Remove(conf.Context.Users, conf.Context.CurrentUser)
			conf.Context.CurrentUser = model.User{}
			go RestoreUser()
		}
	}
}

//GoHome 跳转到首页
func GoHome(w http.ResponseWriter, r *http.Request) {
	//当前登录用户为空，即未登录
	var user model.User
	r.URL.Path = "/index"
	if conf.Context.CurrentUser == user {
		r.URL.Path = "/login"
	}
	GoPage(w, r)
}

//GoPage 遵照约定，加载指定模板，模板名字为url
func GoPage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	t, err := template.ParseFiles("template" + path + ".html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, conf.Context)
}

//SelectLang 处理转换语言的请求
func SelectLang(w http.ResponseWriter, r *http.Request){
	lang := r.FormValue("lang")
	result := make(map[string]interface{})
	w.Header().Add("Content-Type", "application/json")
	if lang == conf.CurrentLang.Value{
		utils.FormatResult(&result, "0400", conf.Context.Locale["langRepeat"], nil)
		res,_ := json.Marshal(result)
		w.Write(res)
		return
	}
	conf.LoadConfig(lang)
	utils.FormatResult(&result, "0200", conf.Context.Locale["selectSuccess"], nil)
	res, _ := json.Marshal(result)
	w.Write(res)
}

//Shutdown 关闭程序，会保留登录信息
func Shutdown(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})
	utils.FormatResult(&result, "0200", conf.Context.Locale["logoutSuccess"], nil)
	w.Header().Add("Content-Type", "application/json")
	if err := Client.Logout(); err != nil {
		utils.FormatResult(&result, "0300", err.Error(), nil)
		res, _ := json.Marshal(result)
		w.Write(res)
		os.Exit(1)
	}
	res, _ := json.Marshal(result)
	w.Write(res)
	os.Exit(0)
}

//Reload 重新加载国际化文件，该api只在开发时使用
func Reload(w http.ResponseWriter, r *http.Request) {
	conf.LoadConfig(conf.Context.CurrentLang.Value)
	r.URL.Path = "/index"
	GoPage(w, r)
}
