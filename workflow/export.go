package workflow

import (
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"strconv"
)

func ExportStudent(student *models.Student, reservation *models.Reservation, filename string) error {
	data := make([][]string, 0)
	data = append(data, []string{"职业生涯咨询知情同意书"})
	data = append(data, []string{"姓名：", student.Fullname, "性别：", student.Gender})
	data = append(data, []string{"年龄：", student.Age, "出生日期：", student.Birthday})
	data = append(data, []string{"民族：", student.Ethnic, "入学年份：", student.EnterYear})
	data = append(data, []string{"生源地：", student.SourcePlace, "院系：", student.College})
	data = append(data, []string{"学生电子邮件：", student.Email, "原就读学校（本科/硕士）：", student.OriginalSchool})
	data = append(data, []string{"原专业（如有转换）：", student.OriginalMajor, "联系电话：", student.Mobile})
	data = append(data, []string{"婚姻状况：", student.Marriage})
	data = append(data, []string{"健康状况：", student.Health})
	data = append(data, []string{"父亲职业：", student.FatherJob, "母亲职业：", student.MotherJob})
	data = append(data, []string{"是否有兄弟姐妹：", student.HasBrotherOrSister})
	data = append(data, []string{"以前是否接受过职业咨询：", student.HasCareerConsulting})
	data = append(data, []string{"以前是否接受过心理咨询：", student.HasMentalConsulting})
	data = append(data, []string{"目前是否在接受其他咨询：", student.OtherConsultingNow})
	data = append(data, []string{"是否有工作经验：", student.WorkingExperience})
	knowingMethods := make([]string, 0)
	knowingMethods = append(knowingMethods, "我们可以通过很多渠道了解与职业生涯有关的信息，最近一个月，你曾使用一下哪些方法：")
	for _, i := range student.KnowingMethods {
		knowingMethods = append(knowingMethods, models.KnowingMethods[i-1])
	}
	data = append(data, knowingMethods)
	if reservation != nil {
		data = append(data, []string{"此次来最主要想解决的问题是：", reservation.StudentExpectation.Problem})
		data = append(data, []string{"你期望职业生涯咨询帮助达到什么样的效果：", reservation.StudentExpectation.Expectation})
		data = append(data, []string{"期望的咨询次数约为：", reservation.StudentExpectation.ExpectedTime})
	}
	data = append(data, []string{"请写下紧急联系人的姓名：", student.EmergencyPerson, "电话：", student.EmergencyMobile})
	if reservation != nil {
		data = append(data, []string{"填写日期：", reservation.StudentExpectation.Time.In(utils.Location).Format(utils.DATE_PATTERN)})
	}
	if err := utils.WriteToCSV(data, filename); err != nil {
		return err
	}
	return nil
}

func ExportStudentInfo(student *models.Student, filename string) error {
	data := make([][]string, 0)
	data = append(data, []string{"职业生涯咨询知情同意书"})
	// 学生基本信息
	data = append(data, []string{"姓名：", student.Fullname, "性别：", student.Gender})
	data = append(data, []string{"年龄：", student.Age, "出生日期：", student.Birthday})
	data = append(data, []string{"民族：", student.Ethnic, "入学年份：", student.EnterYear})
	data = append(data, []string{"生源地：", student.SourcePlace, "院系：", student.College})
	data = append(data, []string{"学生电子邮件：", student.Email, "原就读学校（本科/硕士）：", student.OriginalSchool})
	data = append(data, []string{"原专业（如有转换）：", student.OriginalMajor, "联系电话：", student.Mobile})
	data = append(data, []string{"婚姻状况：", student.Marriage})
	data = append(data, []string{"健康状况：", student.Health})
	data = append(data, []string{"父亲职业：", student.FatherJob, "母亲职业：", student.MotherJob})
	data = append(data, []string{"是否有兄弟姐妹：", student.HasBrotherOrSister})
	data = append(data, []string{"以前是否接受过职业咨询：", student.HasCareerConsulting})
	data = append(data, []string{"以前是否接受过心理咨询：", student.HasMentalConsulting})
	data = append(data, []string{"目前是否在接受其他咨询：", student.OtherConsultingNow})
	data = append(data, []string{"是否有工作经验：", student.WorkingExperience})
	knowingMethods := make([]string, 0)
	knowingMethods = append(knowingMethods, "我们可以通过很多渠道了解与职业生涯有关的信息，最近一个月，你曾使用一下哪些方法：")
	for _, i := range student.KnowingMethods {
		knowingMethods = append(knowingMethods, models.KnowingMethods[i-1])
	}
	data = append(data, knowingMethods)
	data = append(data, []string{"请写下紧急联系人的姓名：", student.EmergencyPerson, "电话：", student.EmergencyMobile})
	data = append(data, []string{""})
	data = append(data, []string{""})

	// 咨询小结
	if reservations, err := models.GetReservationsByStudentId(student.Id.Hex()); err == nil && len(reservations) != 0 {
		data = append(data, []string{"咨询序号", "日期", "星期", "时间段", "咨询师", "问题", "期望效果", "期望咨询次数",
			"预约填写日期", "反馈问题1", "反馈问题2", "反馈问题3", "反馈问题4", "反馈问题5", "反馈问题6", "反馈问题7",
			"反馈问题8", "反馈问题9", "反馈问题10", "反馈问题11", "反馈问题12", "反馈问题13", "反馈问题14", "帮助",
			"未达到预期", "反馈填写日期"})
		for index, res := range reservations {
			if res.Status != models.RESERVATED {
				continue
			}
			row := make([]string, 0)
			row = append(row, strconv.Itoa(index+1))
			row = append(row, res.StartTime.In(utils.Location).Format(utils.DATE_PATTERN))
			row = append(row, utils.Weekdays[res.StartTime.In(utils.Location).Weekday()])
			row = append(row, res.StartTime.In(utils.Location).Format(utils.CLOCK_PATTERN)+" - "+
				res.EndTime.In(utils.Location).Format(utils.CLOCK_PATTERN))
			if teacher, err := models.GetTeacherById(res.TeacherId); err == nil {
				row = append(row, teacher.Fullname)
			} else {
				row = append(row, "")
			}
			row = append(row, res.StudentExpectation.Problem)
			row = append(row, res.StudentExpectation.Expectation)
			row = append(row, res.StudentExpectation.ExpectedTime)
			row = append(row, res.StudentExpectation.Time.In(utils.Location).Format(utils.DATE_PATTERN))
			if !res.StudentFeedback.IsEmpty() {
				for _, s := range res.StudentFeedback.Scores {
					row = append(row, strconv.Itoa(s))
				}
				row = append(row, res.StudentFeedback.Help)
				row = append(row, res.StudentFeedback.Drawback)
				row = append(row, res.StudentFeedback.Time.In(utils.Location).Format(utils.DATE_PATTERN))
			}
			data = append(data, row)
		}
		data = append(data, []string{""})
	}
	if err := utils.WriteToCSV(data, filename); err != nil {
		return err
	}
	return nil
}

func ExportReservations(reservations []*models.Reservation, filename string) error {
	data := make([][]string, 0)
	data = append(data, []string{"序号", "日期", "星期", "时间段", "咨询师", "来访者姓名", "学号", "性别", "入学年份", "院系", "来访次数"})
	for index, res := range reservations {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(index+1))
		row = append(row, res.StartTime.In(utils.Location).Format(utils.DATE_PATTERN))
		row = append(row, utils.Weekdays[res.StartTime.In(utils.Location).Weekday()])
		row = append(row, res.StartTime.In(utils.Location).Format(utils.CLOCK_PATTERN)+" - "+
			res.EndTime.In(utils.Location).Format(utils.CLOCK_PATTERN))
		if teacher, err := models.GetTeacherById(res.TeacherId); err == nil {
			row = append(row, teacher.Fullname)
		} else {
			row = append(row, "")
		}
		if student, err := models.GetStudentById(res.StudentId); err == nil {
			row = append(row, student.Fullname)
			row = append(row, student.Username)
			row = append(row, student.Gender)
			row = append(row, student.EnterYear)
			row = append(row, student.College)
		}
		data = append(data, row)
	}
	if err := utils.WriteToCSV(data, filename); err != nil {
		return err
	}
	return nil
}

func ExportTodayReservationTimetable(reservations []*models.Reservation, filename string) error {
	data := make([][]string, 0)
	today := utils.GetToday()
	data = append(data, []string{today.Format(utils.DATE_PATTERN)})
	data = append(data, []string{"时间", "咨询师", "学生姓名", "联系方式"})
	for _, r := range reservations {
		teacher, err := models.GetTeacherById(r.TeacherId)
		if err != nil {
			continue
		}
		if student, err := models.GetStudentById(r.StudentId); err == nil {
			data = append(data, []string{r.StartTime.In(utils.Location).Format(utils.CLOCK_PATTERN) + " - " + r.EndTime.In(utils.Location).Format(utils.CLOCK_PATTERN),
				teacher.Fullname, student.Fullname, student.Mobile})
		} else {
			data = append(data, []string{r.StartTime.In(utils.Location).Format(utils.CLOCK_PATTERN) + " - " + r.EndTime.In(utils.Location).Format(utils.CLOCK_PATTERN),
				teacher.Fullname, "", ""})
		}
	}
	if err := utils.WriteToCSV(data, filename); err != nil {
		return err
	}
	return nil
}
