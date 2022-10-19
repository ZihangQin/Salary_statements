package commend

import "github.com/360EntSecGroup-Skylar/excelize"

//const SalaryS = "Salary_statements" //工作表名设置为常量

/*设置单元格默认值*/
func SetGenesisCell(f *excelize.File) error {
	//给单元格赋值
	var err error
	err = f.SetCellValue(SalaryS, "A2", "姓名")    //设置默认值
	err = f.SetCellValue(SalaryS, "B2", "职位")    //设置默认值
	err = f.SetCellValue(SalaryS, "C2", "入职时间")  //设置默认值
	err = f.SetCellValue(SalaryS, "D2", "课时")    //设置默认值
	err = f.SetCellValue(SalaryS, "E2", "是否住宿")  //设置默认值
	err = f.SetCellValue(SalaryS, "F2", "是否为党员") //设置默认值
	err = f.SetCellValue(SalaryS, "G2", "本月应得工资")
	err = f.SetCellValue(SalaryS, "H2", "本月应扣除工资")
	err = f.SetCellValue(SalaryS, "I2", "本月到手工资") //最后获取结果

	//设置人
	err = f.SetCellValue(SalaryS, "A3", "严蔚来")
	err = f.SetCellValue(SalaryS, "A4", "李冬梅")
	err = f.SetCellValue(SalaryS, "A5", "谭志虎")
	err = f.SetCellValue(SalaryS, "A6", "谢俊")
	err = f.SetCellValue(SalaryS, "A7", "杨子骏")
	err = f.SetCellValue(SalaryS, "A8", "路人甲")
	err = f.SetCellValue(SalaryS, "A9", "路人乙")
	err = f.SetCellValue(SalaryS, "A10", "演员1")
	err = f.SetCellValue(SalaryS, "A11", "演员2")
	err = f.SetCellValue(SalaryS, "A12", "王庆先")

	//设置职位
	err = f.SetCellValue(SalaryS, "B3", "行政")
	err = f.SetCellValue(SalaryS, "B4", "行政")
	err = f.SetCellValue(SalaryS, "B5", "行政")
	err = f.SetCellValue(SalaryS, "B6", "教研室主任")
	err = f.SetCellValue(SalaryS, "B7", "教师")
	err = f.SetCellValue(SalaryS, "B8", "教师")
	err = f.SetCellValue(SalaryS, "B9", "教师")
	err = f.SetCellValue(SalaryS, "B10", "教师")
	err = f.SetCellValue(SalaryS, "B11", "教师")
	err = f.SetCellValue(SalaryS, "B12", "教师")

	//设置入职时间
	err = f.SetCellValue(SalaryS, "C3", "2020-10-17")
	err = f.SetCellValue(SalaryS, "C4", "2021-10-01")
	err = f.SetCellValue(SalaryS, "C5", "2020-01-17")
	err = f.SetCellValue(SalaryS, "C6", "2018-10-07")
	err = f.SetCellValue(SalaryS, "C7", "2019-08-06")
	err = f.SetCellValue(SalaryS, "C8", "2021-05-16")
	err = f.SetCellValue(SalaryS, "C9", "2019-06-19")
	err = f.SetCellValue(SalaryS, "C10", "2021-09-28")
	err = f.SetCellValue(SalaryS, "C11", "2021-12-12")
	err = f.SetCellValue(SalaryS, "C12", "2022-01-12")

	//设置课时
	err = f.SetCellValue(SalaryS, "D3", "0")
	err = f.SetCellValue(SalaryS, "D4", "0")
	err = f.SetCellValue(SalaryS, "D5", "0")
	err = f.SetCellValue(SalaryS, "D6", "44")
	err = f.SetCellValue(SalaryS, "D7", "35")
	err = f.SetCellValue(SalaryS, "D8", "40")
	err = f.SetCellValue(SalaryS, "D9", "43")
	err = f.SetCellValue(SalaryS, "D10", "33")
	err = f.SetCellValue(SalaryS, "D11", "46")
	err = f.SetCellValue(SalaryS, "D12", "39")

	//设置是否住宿
	err = f.SetCellValue(SalaryS, "E3", "是")
	err = f.SetCellValue(SalaryS, "E4", "是")
	err = f.SetCellValue(SalaryS, "E5", "是")
	err = f.SetCellValue(SalaryS, "E6", "否")
	err = f.SetCellValue(SalaryS, "E7", "是")
	err = f.SetCellValue(SalaryS, "E8", "是")
	err = f.SetCellValue(SalaryS, "E9", "否")
	err = f.SetCellValue(SalaryS, "E10", "否")
	err = f.SetCellValue(SalaryS, "E11", "是")
	err = f.SetCellValue(SalaryS, "E12", "否")

	//设置是否为党员
	err = f.SetCellValue(SalaryS, "F3", "是")
	err = f.SetCellValue(SalaryS, "F4", "是")
	err = f.SetCellValue(SalaryS, "F5", "是")
	err = f.SetCellValue(SalaryS, "F6", "否")
	err = f.SetCellValue(SalaryS, "F7", "是")
	err = f.SetCellValue(SalaryS, "F8", "否")
	err = f.SetCellValue(SalaryS, "F9", "是")
	err = f.SetCellValue(SalaryS, "F10", "否")
	err = f.SetCellValue(SalaryS, "F11", "否")
	err = f.SetCellValue(SalaryS, "F12", "是")
	if err != nil {
		return err
	}
	return nil
}
