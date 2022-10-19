package salary

import (
	"Salary_statements/access2exel"
	"Salary_statements/commend"
	"Salary_statements/utils"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

/*该方法用于计算每个人总工资*/
func EndSalary(f *excelize.File) {
	for i := 2; i <= 11; i++ {
		is := utils.Int2String(i+1)
		axis := "I" + is
		err := f.SetCellFormula(commend.SalaryS, axis, "SUM(G"+is+":H"+is+")")
		if err != nil {
			fmt.Println(err.Error())
		}
		errs := access2exel.SaveExcel(f)
		if errs != nil {
			fmt.Println(errs.Error())
		}
	}
}
