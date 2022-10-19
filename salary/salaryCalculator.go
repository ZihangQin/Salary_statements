package salary

import (
	"Salary_statements/access2exel"
	"Salary_statements/commend"
	"Salary_statements/utils"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

//const SalaryS = "Salary_statements" //工作表名设置为常量

/*该方法用于计算员工津贴*/
func allCalc(f *excelize.File, index int) (int, error) {
	var calc int //总津贴
	//获取行
	s, err := access2exel.GetRow(f, index)
	if err != nil {
		return 0, err
	}
	s2 := utils.GetArrayEle(s, 1)
	if s2 == "教研室主任" { //教研主任津贴加300
		s1 := utils.GetArrayEle(s, 2)
		in := utils.AddAllCalc(s1)
		calc = int(in * (commend.ALLOWANCE + commend.HEADADDALLOWANCE))
		return calc, nil
	} else {
		//获取单元格(按入职时间计算津贴)
		s1 := utils.GetArrayEle(s, 2)
		in := utils.AddAllCalc(s1)
		calc = int(in * commend.ALLOWANCE)
		return calc, nil
	}
}

/*该方法用于计算课时费*/
func addClassHours(f *excelize.File, index int) (int, error) {
	var classHoursAdd int
	//每月课时量*课时费
	s, _ := access2exel.GetRow(f, index)
	//查询职位，如果是行政则无课时
	s1 := utils.GetArrayEle(s, 1)
	if s1 == "行政" {
		return 0, errors.New("行政无课时")
	} else {
		s2 := utils.GetArrayEle(s, 3)
		i, err := utils.String2Int(s2)
		if err != nil {
			return 0, err
		}
		classHoursAdd = i * commend.CLASSHOURS

		return classHoursAdd, nil
	}
}

/*计算出每个员工获得多少工资*/
func AddSalary(f *excelize.File) {
	var BaseSalary int
	var AllSalary int
	for i := 2; i <= 11; i++ {
		s, err := access2exel.GetRow(f, i)
		if err != nil {
			fmt.Println(err.Error())
		}
		s1 := utils.GetArrayEle(s, 1)
		if s1 == "行政" { //行政基本工资4500，且没有课时费
			BaseSalary = commend.ADMINISTRATIVE
			calc, err := allCalc(f, i)
			if err != nil {
				fmt.Println(err.Error())
			}
			AllSalary = BaseSalary + calc
			//fmt.Println(AllSalary) //基本工资加津贴
			is := utils.Int2String(i + 1)
			//allsalary := utils.Int2String(AllSalary)
			axis := "G" + is
			//fmt.Println(axis, allsalary)
			//TODO：打开文件后设置单元格不成功 2022-10-17(已解决)
			err = f.SetCellValue(commend.SalaryS, axis, AllSalary)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = access2exel.SaveExcel(f)
			if err != nil {
				fmt.Println(err.Error())
			}
			//fmt.Println("已经将数据添加到目标文件里")
		} else { //教师基本工资5000，有课时费
			//计算 基本工资+津贴
			BaseSalary = commend.TEACHER
			calc, err := allCalc(f, i)
			if err != nil {
				fmt.Println(err.Error())
			}
			//AllSalary = BaseSalary + calc
			//fmt.Println(AllSalary) //基本工资加津贴
			//计算课时费
			ClassFees, err := addClassHours(f, i)
			if err == errors.New("行政无课时") { //如果有行政计算课时直接跳出本次循环
				continue
			}
			AllSalary = BaseSalary + calc + ClassFees
			//fmt.Println(AllSalary)//教师最后所得工资
			is := utils.Int2String(i + 1)
			//allsalary := utils.Int2String(AllSalary)
			axis := "G" + is
			//fmt.Println(axis, allsalary)
			err = f.SetCellValue(commend.SalaryS, axis, AllSalary)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = access2exel.SaveExcel(f)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
