package conf

import (
	"encoding/json"
	"fmt"
	"github.com/Xuanwo/go-locale"
	"log"
	"os"
	"xmail/model"
)

type context struct {
	//Locale 用于存放xmail本地化的文字
	Locale map[string]string

	//CurrentUser xmail当前登录中的用户
	CurrentUser model.User

	//Users xmail已有登录信息的用户，一般是用户的多个邮箱
	Users []model.User

	//CurrentLang xmail当前选择使用的语言
	CurrentLang model.Language

	//Languages xmail当前支持的语言列表
	Languages []model.Language

	//Emails 当前所在分类的所有邮件,每次最大缓存量50
	Emails []model.Email

	//Count 当前所有的邮件数量
	Count int
}

//全局上下文信息
var Context context = context{Count: 0, Emails: nil}

//xmail启动的配置信息
var Conf map[string]string

//Init 初始化xmail
func init() {
	LoadConfig(getLang())
}

//加载配置文件
func LoadConfig(lang string) {
	//加载xmail启动的配置信息
	confFile, err := os.Open("conf/conf.json")

	if err != nil {
		log.Fatal(err)
	}

	defer confFile.Close()

	json.NewDecoder(confFile).Decode(&Conf)

	//加载默认语言
	if val, ok := Conf["defaultLang"]; ok {
		lang = val
	}

	//加载本app支持的语言列表,只需加载一次
	langFile, err := os.Open("conf/languages.json")
	defer langFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(langFile).Decode(&Context.Languages)

	LoadLocaleFile(lang)
}

//LoadLocaleFile 加载国际化文件
func LoadLocaleFile(lang string) {
	localeFile, err := os.Open("conf/locale/" + lang + ".json")

	defer localeFile.Close()

	if err != nil {
		//app没有完成该语言的版本, 切换成英语
		lang = "en-US"
		if localeFile, err = os.Open("conf/locale/" + lang + ".json"); err != nil {
			log.Fatal(err)
		}
	}

	json.NewDecoder(localeFile).Decode(&Context.Locale)

	//判断当前使用语言
	for _, l := range Context.Languages {
		if l.Value == lang {
			Context.CurrentLang = l
			break
		}
	}

	fmt.Printf("xmail are using: %s\n", Context.CurrentLang.Value)
}

//获取系统语言
func getLang() string {
	tag, err := locale.Detect()
	//默认值
	lang := "en-US"

	if err != nil {
		log.Fatal(err)
		return lang
	}

	fmt.Printf("System's current language: %s\n", tag.String())
	return tag.String();
}
