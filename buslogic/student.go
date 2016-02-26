package buslogic

import (
	"errors"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"github.com/shudiwsh2009/reservation_thzy_go/workflow"
	"strings"
)

type StudentLogic struct {
}

// 学生预约咨询
func (sl *StudentLogic) MakeReservationByStudent(reservationId string, fullname string, gender string, college string,
	mobile string, email string, hasCareerConsulting string, emergencyPerson string, emergencyMobile string,
	age string, birthday string, ethnic string, enterYear string, sourcePlace string, originalSchool string,
	originalMajor string, marriage string, health string, fatherJob string, motherJob string, hasBrotherOrSister string,
	brotherAge string, brotherJob string, hasMentalConsulting string, otherConsultingNow string, workingExperience string,
	workingPeriod string, knowingMethods []int, problem string, expectation string, expectedTime string,
	userId string, userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.STUDENT {
		return nil, errors.New("请重新登录")
	} else if len(fullname) == 0 {
		return nil, errors.New("姓名为空")
	} else if len(gender) == 0 {
		return nil, errors.New("性别为空")
	} else if len(college) == 0 {
		return nil, errors.New("院系为空")
	} else if !utils.IsMobile(mobile) {
		return nil, errors.New("手机号格式不正确")
	} else if !utils.IsEmail(email) {
		return nil, errors.New("邮箱格式不正确")
	} else if len(hasCareerConsulting) == 0 {
		return nil, errors.New("是否接受过职业咨询为空")
	} else if len(emergencyPerson) == 0 {
		return nil, errors.New("紧急联系人为空")
	} else if !utils.IsMobile(emergencyMobile) {
		return nil, errors.New("紧急联系人手机号格式不正确")
	} else if len(problem) == 0 {
		return nil, errors.New("咨询问题为空")
	} else if len(expectation) == 0 {
		return nil, errors.New("期望结果为空")
	}
	student, err := models.GetStudentById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	} else if student.UserType != models.STUDENT {
		return nil, errors.New("请重新登录")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil || reservation.Status == models.DELETED {
		return nil, errors.New("咨询已下架")
	} else if reservation.StartTime.Before(utils.GetNow()) {
		return nil, errors.New("咨询已过期")
	} else if reservation.Status != models.AVAILABLE {
		return nil, errors.New("咨询已被预约")
	}
	studentReservations, err := models.GetReservationsByStudentId(student.Id.Hex())
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	for _, r := range studentReservations {
		if r.Status == models.RESERVATED && r.StartTime.After(utils.GetNow()) {
			return nil, errors.New("你好！你已有一个咨询预约，请完成这次咨询后再预约下一次，或致电62782007取消已有预约。")
		}
	}
	// 更新学生信息
	student.Fullname = fullname
	student.Gender = gender
	student.College = college
	student.Mobile = mobile
	student.Email = email
	student.HasCareerConsulting = hasCareerConsulting
	student.EmergencyPerson = emergencyPerson
	student.EmergencyMobile = emergencyMobile
	student.Age = age
	student.Birthday = birthday
	student.Ethnic = ethnic
	student.EnterYear = enterYear
	student.SourcePlace = sourcePlace
	student.OriginalSchool = originalSchool
	student.OriginalMajor = originalMajor
	student.Marriage = marriage
	student.Health = health
	student.FatherJob = fatherJob
	student.MotherJob = motherJob
	student.HasBrotherOrSister = hasBrotherOrSister
	student.BrotherAge = brotherAge
	student.BrotherJob = brotherJob
	student.HasMentalConsulting = hasMentalConsulting
	student.OtherConsultingNow = otherConsultingNow
	student.WorkingExperience = workingExperience
	student.WorkingPeriod = workingPeriod
	student.KnowingMethods = knowingMethods
	if models.UpsertStudent(student) != nil {
		return nil, errors.New("获取数据失败")
	}
	// 更新咨询信息
	reservation.StudentId = student.Id.Hex()
	reservation.Status = models.RESERVATED
	reservation.StudentExpectation = models.StudentExpectation{
		Problem:      problem,
		Expectation:  expectation,
		ExpectedTime: expectedTime,
	}
	if models.UpsertReservation(reservation) != nil {
		return nil, errors.New("获取数据失败")
	}
	// send success sms
	workflow.SendSuccessSMS(reservation)
	return reservation, nil
}

// 学生拉取反馈
func (sl *StudentLogic) GetFeedbackByStudent(reservationId string,
	userId string, userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.STUDENT {
		return nil, errors.New("请重新登录")
	} else if len(reservationId) == 0 {
		return nil, errors.New("咨询已下架")
	}
	student, err := models.GetStudentById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	} else if student.UserType != models.STUDENT {
		return nil, errors.New("请重新登录")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil || reservation.Status == models.DELETED {
		return nil, errors.New("咨询已下架")
	} else if reservation.StartTime.After(utils.GetNow()) {
		return nil, errors.New("咨询未开始,暂不能反馈")
	} else if reservation.Status == models.AVAILABLE {
		return nil, errors.New("咨询未被预约,不能反馈")
	} else if !strings.EqualFold(reservation.StudentId, student.Id.Hex()) {
		return nil, errors.New("只能反馈本人预约的咨询")
	}
	return reservation, nil
}

// 学生反馈
func (sl *StudentLogic) SubmitFeedbackByStudent(reservationId string, consultingCount string, scores []int,
	help string, drawback string, userId string, userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.STUDENT {
		return nil, errors.New("请重新登录")
	} else if len(reservationId) == 0 {
		return nil, errors.New("咨询已下架")
	} else if len(consultingCount) == 0 || len(scores) != 15 || len(help) == 0 || len(drawback) == 0 {
		return nil, errors.New("请完整填写反馈")
	}
	student, err := models.GetStudentById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	} else if student.UserType != models.STUDENT {
		return nil, errors.New("请重新登录")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil || reservation.Status == models.DELETED {
		return nil, errors.New("咨询已下架")
	} else if reservation.StartTime.After(utils.GetNow()) {
		return nil, errors.New("咨询未开始,暂不能反馈")
	} else if reservation.Status == models.AVAILABLE {
		return nil, errors.New("咨询未被预约,不能反馈")
	} else if !strings.EqualFold(reservation.StudentId, student.Id.Hex()) {
		return nil, errors.New("只能反馈本人预约的咨询")
	}
	reservation.StudentFeedback = models.StudentFeedback{
		ConsultingCount: consultingCount,
		Scores:          scores,
		Help:            help,
		Drawback:        drawback,
		Time:            utils.GetNow(),
	}
	if models.UpsertReservation(reservation) != nil {
		return nil, errors.New("获取数据失败")
	}
	return reservation, nil
}
