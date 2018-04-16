package main

import (
"fmt"
"log"
"net/http"
"strings"
"./vince"
"os/exec"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var filePath string
	var fileName string
	var cmds string
	for k, v := range r.Form {//遍历请求表单的参数对
		if strings.ContainsAny(k,"ls") {//如果是查看路径
			filePath = strings.Join(v, "")
		}
		if strings.ContainsAny(k,"pt") {//如果是查看文件
			fileName = strings.Join(v, "")
		}
		if strings.ContainsAny(k,"cmd") {//如果是要执行命令
			cmds = strings.Join(v, "")
			if cmds == "reboot" {//是否是重启命令
				fmt.Println("reboot")
				cmd := exec.Command("shutdown","-r","now")//生成重启命令
				cmd.Start()//开始执行
			}
		}
	}
	if filePath != "" {//输出文件列表
		fmt.Fprint(w, "the path is :")
		fmt.Fprint(w, filePath)
		fmt.Fprint(w, "\n")
		files,err := vince.ListDir(filePath)
		if err != nil {
			fmt.Fprint(w, err)
		}
		fmt.Fprint(w, files)
	}
	if fileName != "" {//以文本方式输出文件内容
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, "the file is :")
		fmt.Fprint(w, fileName)
		fmt.Fprint(w, "\n")
		datas := vince.OpenFile(fileName)
		fmt.Fprint(w, datas)
	}

}

//程序入口
func main() {
	var hs Hello
	//监听8080端口
	err := http.ListenAndServe("192.168.1.103:8080", hs)
	if err != nil {
		log.Fatal(err)
	}
}
