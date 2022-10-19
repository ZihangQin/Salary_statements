package access2exel

import (
	"Salary_statements/commend"
	"github.com/360EntSecGroup-Skylar/excelize"
)

//const SalaryS = "Salary_statements" //工作表名设置为常量
/*该方法用于获取指定单元格数据*/
func GetACell(f *excelize.File, axis string) (string, error) {
	//操作单元格
	values, err := f.GetCellValue(commend.SalaryS, axis)
	if err != nil {
		return "", err
	}
	return values, nil
}

/*该方法用于设置某一单元格的值*/
func SetACell(f *excelize.File, axis string, values interface{}) error {
	return f.SetCellValue(commend.SalaryS, axis, values)
}

/*该方法用于获取工作表中全部的行数据*/
//TODO: 数据量不大的情况下可用
func getRows(f *excelize.File) ([][]string, error) {
	return f.GetRows(commend.SalaryS)
}

/*该方法用于获取指定行（使用下表获取）*/
func GetRow(f *excelize.File,index int) ([]string, error) {
	row,err := getRows(f)
	if err != nil {
		return nil,err
	}
	return row[index],nil
}
