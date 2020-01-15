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
	Locale      map[string]string
	CurrentUser model.User
	Users       []model.User
	CurrentLang model.Language
	Languages   []model.Language
}

//全局上下文信息
var Context context = context{}

//Init 初始化xmail
func init() {
	LoadConfig(getLang())
}

//加载配置文件
func LoadConfig(lang string) {
	//加载需要插入模板的文本信息
	localeFile, err := os.Open("conf/locale/" + lang + ".json")
	defer localeFile.Close()
	if err != nil {
		//app没有完成该语言的版本, 切换成英语
		lang = "en-US"
		localeFile, err = os.Open("conf/locale/" + lang + ".json")
		if err != nil {
			log.Fatal(err)
		}
	}
	json.NewDecoder(localeFile).Decode(&Context.Locale)
	//加载本app支持的语言列表,只需加载一次
	if len(Context.Languages) == 0 {
		langFile, err := os.Open("conf/languages.json")
		defer langFile.Close()
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(langFile).Decode(&Context.Languages)
	}
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
