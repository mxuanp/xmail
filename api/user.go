package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"xmail/conf"
	"xmail/model"
	"xmail/utils"
)

//Login 登录imap服务器
func Login(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})
	if r.Method != "POST" {
		utils.FormatResult(&result, "0300", conf.Context.Locale["methodError"], nil)
		res, _ := json.Marshal(result)
		w.Write(res)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	email := r.FormValue("email")
	//禁止重复登录
	var nilUser model.User
	if conf.Context.CurrentUser != nilUser && utils.HasUser(conf.Context.Users, email) {
		utils.FormatResult(&result, "0400", conf.Context.Locale["accountRepeat"], nil)
		res,_ := json.Marshal(result)
		w.Write(res)
		return
	}
	//检查是否已经有登录状态，即已有帐号连接到服务器
	if Client != nil{
        Client.Logout()
	}
	password := r.FormValue("pwd")
	host := r.FormValue("host")
	port := r.FormValue("port")
	var err error
	if Client, err = ConnectServer(email, password, host, port); err != nil {
		utils.FormatResult(&result, "0300", err.Error() + "\n" + conf.Context.Locale["confirm"], nil)
		res, _ := json.Marshal(result)
		w.Write(res)
		return
	}
	conf.Context.CurrentUser = model.User{Password: password, Email: email, MailHost: host, Port: port}
	conf.Context.Users = append(conf.Context.Users, conf.Context.CurrentUser)
	utils.FormatResult(&result, "0200", conf.Context.Locale["loginSuccess"], conf.Context.CurrentUser)
	res, _ := json.Marshal(result)
	go RestoreUser()
	w.Write(res)
}

// 退出登录，注销邮箱服务器登录状态
func Logout(w http.ResponseWriter, r *http.Request) {
	Client.Logout()
	conf.Context.Users = utils.Remove(conf.Context.Users, conf.Context.CurrentUser)
	//暂时直接注销登录用户，跳转到登录界面，之后看情况修改逻辑
	conf.Context.CurrentUser = model.User{}
	go RestoreUser()
	GoHome(w, r)
}

// 将用户登录信息写到配置文件
func RestoreUser() {
	f, err := os.OpenFile("conf/user.json", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(f).Encode(conf.Context.Users)
}
