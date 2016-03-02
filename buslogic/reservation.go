package buslogic

import (
	"errors"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"strings"
	"time"
)

type ReservationLogic struct {
}

// 学生查看前后一周内的所有咨询
func (rl *ReservationLogic) GetReservationsByStudent(userId string, userType models.UserType) ([]*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.STUDENT {
		return nil, errors.New("请重新登录")
	}
	student, err := models.GetStudentById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	} else if student.UserType != models.STUDENT {
		return nil, errors.New("请重新登录")
	}
	from := utils.GetNow().AddDate(0, 0, -7)
	to := utils.GetNow().AddDate(0, 0, 7)
	reservations, err := models.GetReservationsBetweenTime(from, to)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	var result []*models.Reservation
	for _, r := range reservations {
		if r.Status == models.AVAILABLE && r.StartTime.Before(utils.GetNow()) {
			continue
		}
		result = append(result, r)
	}
	return result, nil
}

// 咨询师查看负7天之后的所有咨询
func (rl *ReservationLogic) GetReservationsByTeacher(userId string, userType models.UserType) ([]*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.TEACHER {
		return nil, errors.New("权限不足")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	} else if teacher.UserType != models.TEACHER {
		return nil, errors.New("权限不足")
	}
	from := utils.GetNow().AddDate(0, 0, -7)
	reservations, err := models.GetReservationsAfterTime(from)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	var result []*models.Reservation
	for _, r := range reservations {
		if r.Status == models.AVAILABLE && r.StartTime.Before(utils.GetNow()) {
			continue
		} else if strings.EqualFold(r.TeacherId, teacher.Id.Hex()) {
			result = append(result, r)
		}
	}
	return result, nil
}

// 管理员查看负7天之后的所有咨询
func (rl *ReservationLogic) GetReservationsByAdmin(userId string, userType models.UserType) ([]*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错,请联系技术支持")
	}
	from := utils.GetNow().AddDate(0, 0, -7)
	reservations, err := models.GetReservationsAfterTime(from)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	var result []*models.Reservation
	for _, r := range reservations {
		if r.Status == models.AVAILABLE && r.StartTime.Before(utils.GetNow()) {
			continue
		}
		result = append(result, r)
	}
	return result, nil
}

// 管理员查看指定日期后30天内的所有咨询
func (rl *ReservationLogic) GetReservationsMonthlyByAdmin(from string, userId string, userType models.UserType) ([]*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错,请联系技术支持")
	}
	fromDate, err := time.ParseInLocation(utils.DATE_PATTERN, from, utils.Location)
	if err != nil {
		return nil, errors.New("时间格式错误")
	}
	toDate := fromDate.AddDate(0, 0, 30)
	reservations, err := models.GetReservationsBetweenTime(fromDate, toDate)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	var result []*models.Reservation
	for _, r := range reservations {
		if r.Status == models.AVAILABLE && r.StartTime.Before(utils.GetNow()) {
			continue
		}
		result = append(result, r)
	}
	return result, nil
}
