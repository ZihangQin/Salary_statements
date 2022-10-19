package add

import (
	"Salary_statements/salary"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func AddSalarys(file *excelize.File) {
	var isOk string
	fmt.Print("是否计算员工当月获取工资：")
	_, _ = fmt.Scan(&isOk)
	if isOk == "1" {
		salary.AddSalary(file)
		fmt.Println("已将员工当月获取的工资添加进表格")
	}
}
func AddDeuds(file *excelize.File) {
	var isOk string
	fmt.Print("是否计算员工当月需要扣除工资：")
	_, _ = fmt.Scan(&isOk)
	if isOk == "1" {
		salary.DeductSalary(file)
		fmt.Println("已将员工当月扣除的工资添加进表格")
	}
}
func AddEndSarays(file *excelize.File) {
	var isOk string
	fmt.Print("是否计算员工当月实际获取工资：")
	_, _ = fmt.Scan(&isOk)
	if isOk == "1" {
		salary.EndSalary(file)
		fmt.Println("已将员工当月实际获取的工资添加进表格")
	}
}
