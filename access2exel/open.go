package access2exel

import (
	"Salary_statements/commend"
	"github.com/360EntSecGroup-Skylar/excelize"
)

/*该方法用于创建excel文件。ps：初始化一个默认表格*/
func CreateANewExcelFile() (*excelize.File,error) {
	//const SalaryS = "Salary_statements"
	f := excelize.NewFile()
	//创建工作表
	index := f.NewSheet(commend.SalaryS)
	//设置表头
	_ = f.MergeCell(commend.SalaryS, "A1", "I1")
	_ = f.SetCellValue(commend.SalaryS, "A1", "工资报表")
	//设置默认值
	err := commend.SetGenesisCell(f)
	if err != nil {
		return nil,err
	}

	f.SetActiveSheet(index)
	errs := f.SaveAs("./Salary_statements.xlsx")
	if errs != nil {
		return nil,errs
	}
	return f,nil
}

/*该方法用于打开目标excel文件*/
func OpenExcelFile() (*excelize.File, error) {
	return excelize.OpenFile("Salary_statements.xlsx")
}
