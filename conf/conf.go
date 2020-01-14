package conf

import (
	"encoding/json"
	"github.com/Xuanwo/go-locale"
	"log"
	"os"
)

//Init 初始化xmail
func Init() *map[string]string {
	return loadConfig(getLang())
}

//加载配置文件
func loadConfig(lang string) *map[string]string {
	localeFile, e := os.Open("conf/locale/" + lang + ".json")
	if e != nil {
		log.Fatal(e)
		os.Exit(-1)
	}
	locale := make(map[string]string)
	json.NewDecoder(localeFile).Decode(&locale)
	return &locale
}

//获取系统语言
func getLang() string {
	tag, err := locale.Detect()
	//默认值
	lang := "en_US"
	if err != nil {
		log.Fatal(err)
		return lang
	}
	return tag.String();
}
