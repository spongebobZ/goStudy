package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var year string
	var y, m, d int
	for true {
		fmt.Println("请输入出生年份:")
		fmt.Scan(&year)
		_y, error := strconv.Atoi(year)
		if error != nil {
			fmt.Println("转换年份失败")
		} else if _y > 2018 || _y < 1969 {
			fmt.Println("输入的出生年份不合适！")
		} else {
			y = _y
			break
		}
	}
	var month string
	for true {
		fmt.Println("请输入出生月份:")
		fmt.Scan(&month)
		_m, error := strconv.Atoi(month)
		if error != nil {
			fmt.Println("转换月份失败")
		} else if _m > 12 || _m < 1 {
			fmt.Println("输入的出生月份不合适！")
		} else {
			m = _m
			break
		}
	}
	var day string
	for true {
		fmt.Println("请输入出生日:")
		fmt.Scan(&day)
		_d, error := strconv.Atoi(day)
		if error != nil {
			fmt.Println("转换日失败")
		} else if _d > 31 || _d < 1 {
			fmt.Println("输入的出生日期不合适！")
		} else {
			d = _d
			break
		}
	}
	t := time.Now()
	t_birth, _ := time.Parse("2006-01-02 15:04:05", year+"-"+month+"-"+day+" 00:00:00")
	age := t.Sub(t_birth)
	fmt.Printf("你的生日是：%d年%d月%d日\n", y, m, d)
	fmt.Printf("您的年龄是：%d岁\n", int(age.Hours()/24/365))
	time.Sleep(time.Duration(10) * time.Second)
}
