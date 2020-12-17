package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// bpfmt()
	// bpos()
	// bptime()
	// bpstring()
	// bpmath()
	bpregexp()
}

func bpfmt() {
	s := "hahaha"
	// 默认换行
	fmt.Println(s)
	// 支持格式化，默认不换行
	fmt.Printf("i say %v\n", s)

	fmtStr := fmt.Sprintf("prefix: %v", s)
	fmt.Println(fmtStr)
}

func bpos() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Get workdir error!")
	}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Get hostname error!")
	}
	env := os.Environ()
	// for i := 0; i < len(env); i++ {
	for i := range env {
		fmt.Println(env[i])
	}
	fmt.Println(env)
	// os.Exit(0)

	fmt.Printf(path + "\n" + hostname + "\n")

	fileinfo, err := os.Stat("E:\\work\\Program\\golang\\map.go")
	fmt.Println(fileinfo.Mode())

	// 执行命令
	cmd := exec.Command("ls", ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd fail: %v", err)
	} else {
		fmt.Println(out)
	}
}

func bptime() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("years: %v, month: %v, day: %v,hours: %v, minutes: %v, seconds: %v\n", now.Year(), now.Month(), now.Day(), now.Minute(), now.Second())
	fmt.Println(now.Second())
	y, m, d := now.Date()
	fmt.Println(y, "-", m, "-", d)
	// 格式化时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 时间戳转换为格式化时间
	unixtime := time.Now().Unix()
	fmt.Println(time.Unix(unixtime, 0).Format("2006-01-02 15:04:05"))
	// 时间区间
	duration, _ := time.ParseDuration("-24h1m20s")
	fmt.Println(now.Add(10 * duration).Format("2006-01-02 15:04:05"))
	// 休眠
	sleepduration := time.Second * time.Duration(2)
	time.Sleep(sleepduration)
}

func bpstring() {
	var str1 string = "hello, golang! "
	var str2 string = "i am natzhu"
	// 包含
	fmt.Println(strings.Contains(str1, "go"))
	// 拼接
	fmt.Println(str1 + str2)
	s := []string{"abc", "cde", "xyz"}
	fmt.Println(strings.Join(s, ","))
	// 字符串位置索引,不存在字符串返回-1
	fmt.Println(strings.Index(str1, "go"))
	fmt.Println(strings.Index(str1, "gooo"))
	// 字符串个数
	fmt.Println(strings.Count(str1, "o"))
	// 替换字符串，小于0表示全部替换
	fmt.Println(strings.Replace(str1, "o", "22", 1))
	fmt.Println(strings.Replace(str1, "o", "22", -1))
	// 字符串切割为切片
	fmt.Println(strings.Split(str1, ",")[0])
	// 去除头尾指定字符串
	fmt.Println(strings.Trim(str1, " "))

	// 字符串转换与其他类型转换
	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.ParseInt("123", 10, 64))
}

func bpmath() {
	var flo1 float64 = -0.1
	fmt.Println(math.Abs(flo1))
}

func bpregexp() {
	str1 := "haha, i will beat golang"
	// 带POSIX采用POSIX ERE正则语法，不支持\d、\D、\s、\S、\w、\W语法，返回最长匹配结果
	r, _ := regexp.CompilePOSIX("[a-z]+([a-z]+a+)[a-z]+")
	fmt.Println(r.FindAllString(str1, -1))
	fmt.Println(r.FindAllStringSubmatch(str1, -1))
}
