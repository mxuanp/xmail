package main

import (
	"net/http"
	"os/exec"
	"runtime"
	"xmail/conf"

	"xmail/api"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", api.GoHome)
	http.HandleFunc("/index", api.GoPage)
	http.HandleFunc("/login", api.GoPage)
	http.HandleFunc("/compose", api.GoPage)
	http.HandleFunc("/sent", api.GoPage)
	http.HandleFunc("/drafts", api.GoPage)
	http.HandleFunc("/spam", api.GoPage)
	http.HandleFunc("/trash", api.GoPage)
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/logout", api.Logout)
	http.HandleFunc("/api/shutdown", api.Shutdown)
	http.HandleFunc("/api/selectLang", api.SelectLang)
	http.HandleFunc("/api/selectUser", api.SelectUser)
	http.HandleFunc("/api/loadEmail", api.LoadEmail)

	//这个路径专用于开发时使用
	http.HandleFunc("/api/reload", api.Reload)
	go openBrowser()
	http.ListenAndServe(conf.Conf["host"]+":"+conf.Conf["port"], nil)
}

//程序启动时打开浏览器
func openBrowser() {
	sys := runtime.GOOS
	if sys == "linux" {
		exec.Command("xdg-open", "http://localhost:8080").Start()
		return
	}
	if sys == "windows" {
		exec.Command("cmd", "/c", "start", "http://localhost:8080").Start()
		return
	}
	if sys == "darwin" {
		exec.Command("open", "http://localhost:8080").Start()
	}
}

//捕获 ctrl+c
func catchCtrl() {

}
