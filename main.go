package main

import (
	"net/http"
	"os/exec"
	"runtime"

	"xmail/api"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", api.GoHome)
	http.HandleFunc("/login", api.GoPage)
	http.HandleFunc("/compose", api.GoPage)
	http.HandleFunc("/sent", api.GoPage)
	http.HandleFunc("/drafts", api.GoPage)
	http.HandleFunc("/spam", api.GoPage)
	http.HandleFunc("/trash", api.GoPage)
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/logout", api.Logout)
	http.HandleFunc("/api/shutdown", api.Shutdown)
	go openBrowser()
	http.ListenAndServe("localhost:8080", nil)
}
func openBrowser() {
	sys := runtime.GOOS
	if sys == "linux" {
		exec.Command("xdg-open", "http://localhost:8080").Start()
		return
	}
	if sys == "windows" {
		exec.Command("cmd", "/c", "start", "http://localhost").Start()
		return
	}
	if sys == "darwin" {
		exec.Command("open", "http://localhost:8080").Start()
	}
}
