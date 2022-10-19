package main

import (
	"Salary_statements/access2exel"
	"Salary_statements/add"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"sync"
)

var file *excelize.File
var waitGroup sync.WaitGroup

func initOpenOrCreateFile() {
	//fmt.Println("这里是init函数")
	var err error
	file, err = access2exel.OpenExcelFile()
	if err != nil {
		fmt.Println(err.Error())
	}
	if file == nil {
		//fmt.Println("没有工作标创建")
		//如果没有就创建excel工作表
		file, err = access2exel.CreateANewExcelFile()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(1)
	go initOpenOrCreateFile() //开启创建或打开文件的协程
	waitGroup.Wait()          //停止等待组完成协程
	//TODO: 不能异步进行 2022-10-18(用同步等待组 goroutine解决)
	fmt.Println("在本程序中，需要输入是不是的请输入1或者2")
	add.AddSalarys(file)   //添加工资
	add.AddDeuds(file)     //添加应扣除工资
	add.AddEndSarays(file) //计算总工资
	//TODO：打开文件后设置单元格不成功 2022-10-17(已解决，对单元格更改后未进行保存)
	fmt.Println("程序执行完成")
}
