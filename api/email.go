package api

import "github.com/emersion/go-imap/client"

//GetAllMail 获取当前用户所有邮件
func GetAllMail() {

}

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
