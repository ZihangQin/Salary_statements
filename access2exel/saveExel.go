package access2exel

import "github.com/360EntSecGroup-Skylar/excelize"

func SaveExcel(f *excelize.File) error {
	errs := f.SaveAs("./Salary_statements.xlsx")
	if errs != nil {
		return errs
	}
	return nil
}
