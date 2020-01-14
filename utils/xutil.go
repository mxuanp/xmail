//utils包包含一些作者封装的用于本项目的工具
//xutil包含一些和前端交互需要的工具
package utils

//FormatResult 将数据格式化到map
func FormatResult(res *map[string]interface{}, code string, message string, data interface{}) {
	(*res)["code"] = code
	(*res)["message"] = message
	(*res)["data"] = data
}
