package workflow

import (
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
)

func ExportStudentInfo(student *models.Student, filename string) error {
	//	data := make([][]string, 0)
	//	data = append(data, []string{"档案分类", student.ArchiveCategory, "档案编号", student.ArchiveNumber})
	//	// 学生基本信息
	//	data = append(data, []string{"学号", student.Username})
	//	data = append(data, []string{"姓名", student.Fullname})
	//	data = append(data, []string{"性别", student.Gender})
	//	data = append(data, []string{"出生日期", student.Birthday})
	//	data = append(data, []string{"系别", student.School})
	//	data = append(data, []string{"年级", student.Grade})
	//	data = append(data, []string{"现住址", student.CurrentAddress})
	//	data = append(data, []string{"家庭住址", student.FamilyAddress})
	//	data = append(data, []string{"联系电话", student.Mobile})
	//	data = append(data, []string{"Email", student.Email})
	//	if !student.Experience.IsEmpty() {
	//		data = append(data, []string{"咨询经历", "时间", student.Experience.Time, "地点", student.Experience.Location,
	//			"咨询师姓名", student.Experience.Teacher})
	//	} else {
	//		data = append(data, []string{"咨询经历", "无"})
	//	}
	//	data = append(data, []string{"父亲", "年龄", student.FatherAge, "职业", student.FatherJob, "学历", student.FatherEdu})
	//	data = append(data, []string{"母亲", "年龄", student.MotherAge, "职业", student.MotherJob, "学历", student.MotherEdu})
	//	data = append(data, []string{"父母婚姻状况", student.ParentMarriage})
	//	data = append(data, []string{"在近三个月里，是否发生了对你有重大意义的事（如亲友的死亡、法律诉讼、失恋等）？", student.Significant})
	//	data = append(data, []string{"你现在需要接受帮助的主要问题是什么？", student.Problem})
	//	bindedTeacher, err := models.GetTeacherById(student.BindedTeacherId)
	//	if err != nil {
	//		data = append(data, []string{"匹配咨询师", "无"})
	//	} else {
	//		data = append(data, []string{"匹配咨询师", bindedTeacher.Username, bindedTeacher.Fullname})
	//	}
	//	data = append(data, []string{"危机等级", strconv.Itoa(student.CrisisLevel)})
	//	data = append(data, []string{""})
	//	data = append(data, []string{""})
	//
	//	//咨询小结
	//	if reservations, err := models.GetReservationsByStudentId(student.Id.Hex()); err == nil {
	//		for i, r := range reservations {
	//			teacher, err := models.GetTeacherById(r.TeacherId)
	//			if err != nil {
	//				continue
	//			}
	//			data = append(data, []string{"咨询小结" + strconv.Itoa(i+1)})
	//			data = append(data, []string{"咨询师", teacher.Username, teacher.Fullname})
	//			data = append(data, []string{"咨询日期", r.StartTime.In(utils.Location).Format(utils.DATE_PATTERN)})
	//			if !r.TeacherFeedback.IsEmpty() {
	//				data = append(data, []string{"评估分类", models.FeedbackAllCategory[r.TeacherFeedback.Category]})
	//				participants := []string{"出席人员"}
	//				for j := 0; j < len(r.TeacherFeedback.Participants); j++ {
	//					if r.TeacherFeedback.Participants[j] > 0 {
	//						participants = append(participants, models.Reservation_Participants[j])
	//					}
	//				}
	//				data = append(data, participants)
	//				data = append(data, []string{"问题评估", r.TeacherFeedback.Problem})
	//				data = append(data, []string{"咨询记录", r.TeacherFeedback.Record})
	//			}
	//			if !r.StudentFeedback.IsEmpty() {
	//				scores := []string{"来访者反馈"}
	//				for _, s := range r.StudentFeedback.Scores {
	//					scores = append(scores, strconv.Itoa(s))
	//				}
	//				data = append(data, scores)
	//			}
	//		}
	//		data = append(data, []string{""})
	//	}
	//	if err := utils.WriteToCSV(data, filename); err != nil {
	//		return err
	//	}
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
