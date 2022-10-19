package commend

/*员工固定增加工资*/
const (
	ALLOWANCE      = 500  //津贴
	ADMINISTRATIVE = 4500 //行政基本工资
	TEACHER        = 5000 //教师基本工资
	CLASSHOURS     = 40   //课时费

	HEADADDALLOWANCE = 300 //教研主任额外增加的津贴
)

/*员工固定删减工资*/
const (
	HOUSEING = -300 //房租
	//Five-insurance payment
	FIVEINSURANCEPAYMENT = 0.15 //五险一金
)

/*工作表名*/
const SalaryS = "Salary_statements" //工作表名设置为常量
