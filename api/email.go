package api

import (
	"github.com/emersion/go-imap/client"
	"net/http"
	"xmail/conf"
)

// ConnectServer 连接邮箱服务器
func ConnectServer(email, password, host, port string) (*client.Client, error) {
	c, err := client.DialTLS(host+":"+port, nil)

	if err != nil {
		return nil, err
	}

	if err := c.Login(email, password); err != nil {
		return nil, err
	}

	return c, nil
}

//GetEmailCount 获取当前类别的邮件的所有数量
func GetEmailCount(path string){
	conf.Context.Count = 0
}

//LoadEmail 先向服务器请求最新邮件, 再从本地获取旧邮件, 以json格式返回字符串
func LoadEmail(w http.ResponseWriter, r *http.Request) {
	category := []string{"/index", "/drafts", "/sent", "/spam", "/trash"}

	path := r.FormValue("path")
	//page, _ := strconv.Atoi(r.FormValue("page"))

	w.Header().Add("Content-type", "application/json")

	for _, p := range category {
		//加载邮件, 逻辑暂不实现
		if path == p {
			break
		}
	}
}
