package salary

import (
	"Salary_statements/access2exel"
	"Salary_statements/commend"
	"Salary_statements/utils"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

/*该方法计算住宿人所扣房租*/
func chargeHousing(f *excelize.File, index int) (int, error) {
	var chargeHousing int
	s, err := access2exel.GetRow(f, index)
	if err != nil {
		return 0, err
	}
	isHousing := utils.GetArrayEle(s, 4)
	if isHousing == "" {
		return 0, errors.New("该员工未填写是否住宿")
	}
	if isHousing == "是" {
		chargeHousing = commend.HOUSEING
		//fmt.Println(chargeHousing)
		return chargeHousing, nil
	} else {
		chargeHousing = 0
		return chargeHousing, nil
	}
}

/*该方法用具计算交纳五险一金*/
func fiveInsurancePayment(f *excelize.File, index int) (float64, error) {
	var fivPay float64
	s, err := access2exel.GetRow(f, index)
	if err != nil {
		return 0, err
	}
	getSalary := utils.GetArrayEle(s, 6)
	if getSalary == "" {
		return 0, errors.New("该员工还未计算总工资")
	}
	getSalaryInt, errs := utils.String2Int(getSalary)
	if errs != nil {
		return 0, errs
	}

	fivPay = float64(getSalaryInt) * commend.FIVEINSURANCEPAYMENT
	return fivPay, nil
}

/*该方法用于计算工资税*/
func centralTax(f *excelize.File, index int) (float64, error) {
	var tax float64
	var exemption float64 = 5000
	s, err := access2exel.GetRow(f, index)
	if err != nil {
		return 0, err
	}
	//获取总工资列
	salary := utils.GetArrayEle(s, 6)
	salaryInt, errs := utils.String2Int(salary)
	//fmt.Println("总工资为：", salaryInt)

	if errs != nil {
		return 0, errs
	}

	//计算需要扣除的个人所得税
	if salaryInt <= 5000 {
		tax = 0
		return tax, nil
	} else if salaryInt > 5000 && salaryInt <= 8000 {
		tax = (float64(salaryInt) - exemption) * 0.03
		return tax, nil
	} else if salaryInt > 8000 && salaryInt <= 17000 {
		tax = (float64(salaryInt) - exemption) * 0.1
		return tax, nil
	} else if salaryInt > 17000 && salaryInt <= 30000 {
		tax = (float64(salaryInt) - exemption) * 0.2
		return tax, nil
	} else if salaryInt > 30000 && salaryInt <= 40000 {
		tax = (float64(salaryInt) - exemption) * 0.25
		return tax, nil
	} else if salaryInt > 40000 && salaryInt <= 60000 {
		tax = (float64(salaryInt) - exemption) * 0.3
		return tax, nil
	}
	return 0, nil
}

/*该方法用于计算党员应交党费*/
func payParty(f *excelize.File, index int) (float64, error) {
	var pp float64
	//判断是否为党员
	s, err := access2exel.GetRow(f, index)
	if err != nil {
		return 0, err
	}

	isPP := utils.GetArrayEle(s, 5)
	if isPP == "否" { //不是党员无需交纳
		return 0, nil
	} else if isPP == "是" { //是党员要判断税后工资
		salary := utils.GetArrayEle(s, 6) //获取总工资
		salaryInt, _ := utils.String2Int(salary)
		tax, err := centralTax(f, index)
		if err != nil {
			return 0, err
		}
		AT := float64(salaryInt) - tax //计算税后工资
		if AT <= 3000 {
			pp = float64(salaryInt) * 0.005
			return pp, nil
		} else if AT > 3000 && AT <= 5000 {
			pp = float64(salaryInt) * 0.01
			return pp, nil
		} else if AT > 5000 && AT <= 10000 {
			pp = float64(salaryInt) * 0.015
			return pp, nil
		} else {
			pp = float64(salaryInt) * 0.02
			return pp, nil
		}
	} else {
		return 0, errors.New("获取数据出错")
	}
}

/*计算每个员工每月应扣除多少*/
func DeductSalary(f *excelize.File) {
	var Deduct float64
	for i := 2; i <= 11; i++ {
		//房租
		housing, errHousSing := chargeHousing(f, i)
		if errHousSing != nil {
			fmt.Println(errHousSing.Error())
		}
		//五险一金
		fiv, errFiv := fiveInsurancePayment(f, i)
		if errFiv != nil {
			fmt.Println(errFiv.Error())
		}
		//扣税
		tax, errTax := centralTax(f, i)
		if errTax != nil {
			fmt.Println(errTax.Error())
		}
		//党费
		PP, errPP := payParty(f, i)
		if errPP != nil {
			fmt.Println(errPP.Error())
		}
		Deduct = float64(housing) - fiv - tax - PP
		//保存在表格中
		is := utils.Int2String(i + 1)
		//deduct := utils.Float2String(Deduct)
		axis := "H"+is
		err := f.SetCellValue(commend.SalaryS,axis,Deduct)
		if err != nil {
			fmt.Println(err.Error())
		}
		errs := access2exel.SaveExcel(f)
		if errs != nil {
			fmt.Println(errs.Error())
		}
	}
}
