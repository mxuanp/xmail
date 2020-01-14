package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"xmail/model"
	"xmail/utils"
)

//Login 登录imap服务器
func Login(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})
	if r.Method != "POST" {
		utils.FormatResult(&result, "0300", "just support post method", nil)
		res, _ := json.Marshal(result)
		w.Write(res)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("pwd")
	host := r.FormValue("host")
	port := r.FormValue("port")
	w.Header().Add("Content-Type", "application/json")
	var err error
	if Client, err = ConnectServer(email, password, host, port); err != nil {
		utils.FormatResult(&result, "0300", err.Error(), nil)
		res, _ := json.Marshal(result)
		w.Write(res)
		return
	}
	CurrentUser = model.User{Password: password, Email: email, MailHost: host, Port: port}
	Users = append(Users, CurrentUser)
	utils.FormatResult(&result, "0200", (*Locale)["loginSuccess"], CurrentUser)
	res, _ := json.Marshal(result)
	go RestoreUser()
	w.Write(res)
}

// 退出登录，注销邮箱服务器登录状态
func Logout(w http.ResponseWriter, r *http.Request) {
	Client.Logout()
	Users = utils.Remove(Users, CurrentUser)
	CurrentUser = model.User{}
	go RestoreUser()
	GoHome(w, r)
}

// 将用户登录信息写到配置文件
func RestoreUser() {
	f, err := os.OpenFile("conf/user.json", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(f).Encode(Users)
}
