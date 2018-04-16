package vince

import (
"io/ioutil"//IO操作包
"os"//系统工具包
)

func ListDir(dirPath string) (files []string, err error) {
	files = make([]string,0,10)//初始化数组,长度为0，初始容量为10，容量会随实际增加
	dir,err := ioutil.ReadDir(dirPath)//读指定路径
	if err != nil {//如果有异常则返回异常
		return nil,err
	}
	PathSep := string(os.PathSeparator)//系统路径分隔符
	for _,file := range dir {//遍历路径列表（如果把i换成_则可以不用fmt包）
		files = append(files,dirPath + PathSep + file.Name() + "\n")//拼接路径并加入数组中
	}
	return files,nil//返回文件列表
}

func OpenFile(filePath string) (datas string) {
	data,err := ioutil.ReadFile(filePath)//打开指定的文件
	if err != nil {
		return ""
	}
	return string(data)//以文本方式返回文件内容
}
