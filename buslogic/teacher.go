package buslogic

import (
	"errors"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"strings"
	"time"
)

type TeacherLogic struct {
}

// 咨询师添加咨询
func (tl *TeacherLogic) AddReservationByTeacher(startTime string, endTime string, teacherFullname string,
	teacherMobile string, force bool, userId string, userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.TEACHER {
		return nil, errors.New("权限不足")
	} else if len(startTime) == 0 {
		return nil, errors.New("开始时间为空")
	} else if len(endTime) == 0 {
		return nil, errors.New("结束时间为空")
	} else if len(teacherFullname) == 0 {
		return nil, errors.New("咨询师姓名为空")
	} else if len(teacherMobile) == 0 {
		return nil, errors.New("咨询师手机号为空")
	} else if !utils.IsMobile(teacherMobile) {
		return nil, errors.New("咨询师手机号格式不正确")
	}
	start, err := time.ParseInLocation(utils.TIME_PATTERN, startTime, utils.Location)
	if err != nil {
		return nil, errors.New("开始时间格式错误")
	}
	end, err := time.ParseInLocation(utils.TIME_PATTERN, endTime, utils.Location)
	if err != nil {
		return nil, errors.New("结束时间格式错误")
	}
	if start.After(end) {
		return nil, errors.New("开始时间不能晚于结束时间")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil {
		return nil, errors.New("咨询师账户失效")
	} else if teacher.UserType != models.TEACHER {
		return nil, errors.New("权限不足")
	}
	if !strings.EqualFold(teacher.Fullname, teacherFullname) || !strings.EqualFold(teacher.Mobile, teacherMobile) {
		if !force {
			return nil, errors.New(models.CHECK_MESSAGE)
		}
		teacher.Fullname = teacherFullname
		teacher.Mobile = teacherMobile
		if models.UpsertTeacher(teacher) != nil {
			return nil, errors.New("获取数据失败")
		}
	}
	reservation, err := models.AddReservation(start, end, teacher.Id.Hex())
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	return reservation, nil
}

// 咨询师编辑咨询
func (tl *TeacherLogic) EditReservationByTeacher(reservationId string, startTime string, endTime string,
	teacherFullname string, teacherMobile string, force bool, userId string, userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.TEACHER {
		return nil, errors.New("权限不足")
	} else if len(reservationId) == 0 {
		return nil, errors.New("咨询已下架")
	} else if len(startTime) == 0 {
		return nil, errors.New("开始时间为空")
	} else if len(endTime) == 0 {
		return nil, errors.New("结束时间为空")
	} else if len(teacherFullname) == 0 {
		return nil, errors.New("咨询师姓名为空")
	} else if len(teacherMobile) == 0 {
		return nil, errors.New("咨询师手机号为空")
	} else if !utils.IsMobile(teacherMobile) {
		return nil, errors.New("咨询师手机号格式不正确")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil || reservation.Status == models.DELETED {
		return nil, errors.New("咨询已下架")
	} else if reservation.Status == models.RESERVATED {
		return nil, errors.New("不能编辑已被预约的咨询")
	}
	start, err := time.ParseInLocation(utils.TIME_PATTERN, startTime, utils.Location)
	if err != nil {
		return nil, errors.New("开始时间格式错误")
	}
	end, err := time.ParseInLocation(utils.TIME_PATTERN, endTime, utils.Location)
	if err != nil {
		return nil, errors.New("结束时间格式错误")
	}
	if start.After(end) {
		return nil, errors.New("开始时间不能晚于结束时间")
	} else if end.Before(time.Now().In(utils.Location)) {
		return nil, errors.New("不能编辑已过期咨询")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil {
		return nil, errors.New("咨询师账户失效")
	} else if teacher.UserType != models.TEACHER {
		return nil, errors.New("权限不足")
	} else if !strings.EqualFold(reservation.TeacherId, teacher.Id.Hex()) {
		return nil, errors.New("只能编辑本人开设的咨询")
	}
	if !strings.EqualFold(teacher.Fullname, teacherFullname) || !strings.EqualFold(teacher.Mobile, teacherMobile) {
		if !force {
			return nil, errors.New(models.CHECK_MESSAGE)
		}
		teacher.Fullname = teacherFullname
		teacher.Mobile = teacherMobile
		if models.UpsertTeacher(teacher) != nil {
			return nil, errors.New("获取数据失败")
		}
	}
	reservation.StartTime = start
	reservation.EndTime = end
	if models.UpsertReservation(reservation) != nil {
		return nil, errors.New("获取数据失败")
	}
	return reservation, nil
}

// 咨询师删除咨询
func (tl *TeacherLogic) RemoveReservationsByTeacher(reservationIds []string, userId string, userType models.UserType) (int, error) {
	if len(userId) == 0 {
		return 0, errors.New("请先登录")
	} else if userType != models.TEACHER {
		return 0, errors.New("权限不足")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil {
		return 0, errors.New("咨询师账户失效")
	} else if teacher.UserType != models.TEACHER {
		return 0, errors.New("权限不足")
	}
	removed := 0
	for _, reservationId := range reservationIds {
		if reservation, err := models.GetReservationById(reservationId); err == nil &&
			strings.EqualFold(reservation.TeacherId, teacher.Id.Hex()) {
			reservation.Status = models.DELETED
			if models.UpsertReservation(reservation) == nil {
				removed++
			}
		}
	}
	return removed, nil
}

// 咨询师取消预约
func (tl *TeacherLogic) CancelReservationsByTeacher(reservationIds []string, userId string, userType models.UserType) (int, error) {
	if len(userId) == 0 {
		return 0, errors.New("请先登录")
	} else if userType != models.TEACHER {
		return 0, errors.New("权限不足")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil {
		return 0, errors.New("咨询师账户失效")
	} else if teacher.UserType != models.TEACHER {
		return 0, errors.New("权限不足")
	}
	canceled := 0
	for _, reservationId := range reservationIds {
		reseravtion, err := models.GetReservationById(reservationId)
		if err != nil || reseravtion.Status == models.DELETED ||
			!strings.EqualFold(reseravtion.TeacherId, teacher.Id.Hex()) {
			continue
		}
		if reseravtion.Status == models.RESERVATED && reseravtion.StartTime.After(utils.GetNow()) {
			reseravtion.Status = models.AVAILABLE
			reseravtion.StudentId = ""
			reseravtion.StudentFeedback = models.StudentFeedback{}
			reseravtion.StudentExpectation = models.StudentExpectation{}
			if models.UpsertReservation(reseravtion) == nil {
				canceled++
			}
		}
	}
	return canceled, nil
}

// 咨询师查看学生的预约信息
func (tl *TeacherLogic) GetReservatingStudentInfoByTeacher(reservationId string,
	userId string, userType models.UserType) (*models.Student, *models.Reservation, error) {
	if len(userId) == 0 {
		return nil, nil, errors.New("请先登录")
	} else if userType != models.TEACHER {
		return nil, nil, errors.New("权限不足")
	} else if len(reservationId) == 0 {
		return nil, nil, errors.New("咨询已下架")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil || teacher.UserType != models.TEACHER {
		return nil, nil, errors.New("咨询师账户失效")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil || reservation.Status == models.DELETED {
		return nil, nil, errors.New("咨询已下架")
	} else if reservation.Status == models.AVAILABLE {
		return nil, nil, errors.New("咨询未被预约，无法查看")
	} else if !strings.EqualFold(reservation.TeacherId, teacher.Id.Hex()) {
		return nil, nil, errors.New("只能查看本人开设的咨询")
	}
	student, err := models.GetStudentById(reservation.StudentId)
	if err != nil {
		return nil, nil, errors.New("咨询未被预约，无法查看")
	}
	return student, reservation, nil
}

// 咨询师查询学生信息
func (tl *TeacherLogic) QueryStudentInfoByTeacher(studentUsername string,
	userId string, userType models.UserType) (*models.Student, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.TEACHER {
		return nil, errors.New("权限不足")
	} else if len(studentUsername) == 0 {
		return nil, errors.New("学号为空")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil {
		return nil, errors.New("咨询师账户失效")
	} else if teacher.UserType != models.TEACHER {
		return nil, errors.New("权限不足")
	}
	student, err := models.GetStudentByUsername(studentUsername)
	if err != nil || student.UserType != models.STUDENT {
		return nil, errors.New("学生未注册")
	}
	return student, nil
}