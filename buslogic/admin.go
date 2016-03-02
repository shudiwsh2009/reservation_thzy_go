package buslogic

import (
	"errors"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"strings"
	"time"
"github.com/shudiwsh2009/reservation_thzy_go/workflow"
)

type AdminLogic struct {
}

// 管理员添加咨询
func (al *AdminLogic) AddReservationByAdmin(startTime string, endTime string, teacherUsername string,
	teacherFullname string, teacherMobile string, force bool, userId string, userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	} else if len(startTime) == 0 {
		return nil, errors.New("开始时间为空")
	} else if len(endTime) == 0 {
		return nil, errors.New("结束时间为空")
	} else if len(teacherUsername) == 0 {
		return nil, errors.New("咨询师工号为空")
	} else if len(teacherFullname) == 0 {
		return nil, errors.New("咨询师姓名为空")
	} else if len(teacherMobile) == 0 {
		return nil, errors.New("咨询师手机号为空")
	} else if !utils.IsMobile(teacherMobile) {
		return nil, errors.New("咨询师手机号格式不正确")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错,请联系技术支持")
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
	teacher, err := models.GetTeacherByUsername(teacherUsername)
	if err != nil {
		if teacher, err = models.AddTeacher(teacherUsername, TeacherDefaultPassword, teacherFullname, teacherMobile); err != nil {
			return nil, errors.New("获取数据失败")
		}
	} else if teacher.UserType != models.TEACHER {
		return nil, errors.New("权限不足")
	} else if !strings.EqualFold(teacher.Fullname, teacherFullname) || !strings.EqualFold(teacher.Mobile, teacherMobile) {
		if !force {
			return nil, errors.New(models.CHECK_MESSAGE)
		}
		teacher.Fullname = teacherFullname
		teacher.Mobile = teacherMobile
		if models.UpsertTeacher(teacher) != nil {
			return nil, errors.New("数据获取失败")
		}
	}
	reservation, err := models.AddReservation(start, end, teacher.Id.Hex())
	if err != nil {
		return nil, errors.New("数据获取失败")
	}
	return reservation, nil
}

// 管理员编辑咨询
func (al *AdminLogic) EditReservationByAdmin(reservationId string, startTime string, endTime string,
	teacherUsername string, teacherFullname string, teacherMobile string, force bool, userId string,
	userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	} else if len(reservationId) == 0 {
		return nil, errors.New("咨询已下架")
	} else if len(startTime) == 0 {
		return nil, errors.New("开始时间为空")
	} else if len(endTime) == 0 {
		return nil, errors.New("结束时间为空")
	} else if len(teacherUsername) == 0 {
		return nil, errors.New("咨询师工号为空")
	} else if len(teacherFullname) == 0 {
		return nil, errors.New("咨询师姓名为空")
	} else if len(teacherMobile) == 0 {
		return nil, errors.New("咨询师手机号为空")
	} else if !utils.IsMobile(teacherMobile) {
		return nil, errors.New("咨询师手机号格式不正确")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错,请联系技术支持")
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
	teacher, err := models.GetTeacherByUsername(teacherUsername)
	if err != nil {
		if teacher, err = models.AddTeacher(teacherUsername, TeacherDefaultPassword, teacherFullname, teacherMobile); err != nil {
			return nil, errors.New("获取数据失败")
		}
	} else if teacher.UserType != models.TEACHER {
		return nil, errors.New("权限不足")
	} else if !strings.EqualFold(teacher.Fullname, teacherFullname) || !strings.EqualFold(teacher.Mobile, teacherMobile) {
		if !force {
			return nil, errors.New(models.CHECK_MESSAGE)
		}
		teacher.Fullname = teacherFullname
		teacher.Mobile = teacherMobile
		if models.UpsertTeacher(teacher) != nil {
			return nil, errors.New("数据获取失败")
		}
	}
	reservation.StartTime = start
	reservation.EndTime = end
	reservation.TeacherId = teacher.Id.Hex()
	if err = models.UpsertReservation(reservation); err != nil {
		return nil, errors.New("数据获取失败")
	}
	return reservation, nil
}

// 管理员删除咨询
func (al *AdminLogic) RemoveReservationsByAdmin(reservationIds []string, userId string, userType models.UserType) (int, error) {
	if len(userId) == 0 {
		return 0, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return 0, errors.New("权限不足")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return 0, errors.New("管理员账户出错,请联系技术支持")
	}
	removed := 0
	for _, reservationId := range reservationIds {
		if reservation, err := models.GetReservationById(reservationId); err == nil {
			reservation.Status = models.DELETED
			if models.UpsertReservation(reservation) == nil {
				removed++
			}
		}
	}
	return removed, nil
}

// 管理员取消预约
func (al *AdminLogic) CancelReservationsByAdmin(reservationIds []string, userId string, userType models.UserType) (int, error) {
	if len(userId) == 0 {
		return 0, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return 0, errors.New("权限不足")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return 0, errors.New("管理员账户出错,请联系技术支持")
	}
	canceled := 0
	for _, reservationId := range reservationIds {
		reseravtion, err := models.GetReservationById(reservationId)
		if err != nil || reseravtion.Status == models.DELETED {
			continue
		}
		if reseravtion.Status == models.RESERVATED && reseravtion.StartTime.After(time.Now().Local()) {
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

// 管理员查看学生的预约信息
func (al *AdminLogic) GetReservatingStudentInfoByAdmin(reservationId string,
	userId string, userType models.UserType) (*models.Student, *models.Reservation, error) {
	if len(userId) == 0 {
		return nil, nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, nil, errors.New("权限不足")
	} else if len(reservationId) == 0 {
		return nil, nil, errors.New("咨询已下架")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, nil, errors.New("管理员账户出错,请联系技术支持")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil || reservation.Status == models.DELETED {
		return nil, nil, errors.New("咨询已下架")
	} else if reservation.Status == models.AVAILABLE {
		return nil, nil, errors.New("咨询未被预约，无法查看")
	}
	student, err := models.GetStudentById(reservation.StudentId)
	if err != nil {
		return nil, nil, errors.New("咨询未被预约，无法查看")
	}
	return student, reservation, nil
}

// 管理员重置学生密码
func (al *AdminLogic) ResetStudentPasswordByAdmin(studentId string, password string,
	userId string, userType models.UserType) (*models.Student, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	} else if len(studentId) == 0 {
		return nil, errors.New("学生未注册")
	} else if len(password) == 0 {
		return nil, errors.New("密码为空")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错,请联系技术支持")
	}
	student, err := models.GetStudentById(studentId)
	if err != nil {
		return nil, errors.New("学生未注册")
	}
	student.Password = password
	if models.UpsertStudent(student) != nil {
		return nil, errors.New("获取数据失败")
	}
	return student, nil
}

// 管理员删除学生账户
func (al *AdminLogic) DeleteStudentAccountByAdmin(studentId string, userId string, userType models.UserType) error {
	if len(userId) == 0 {
		return errors.New("请先登录")
	} else if userType != models.ADMIN {
		return errors.New("权限不足")
	} else if len(studentId) == 0 {
		return errors.New("学生未注册")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return errors.New("管理员账户出错,请联系技术支持")
	}
	student, err := models.GetStudentById(studentId)
	if err != nil || student.UserType != models.STUDENT {
		return errors.New("学生未注册")
	}
	student.UserType = models.UNKNOWN
	if models.UpsertStudent(student) != nil {
		return errors.New("获取数据失败")
	}
	return nil
}

// 管理员导出知情同意书
func (al *AdminLogic) ExportReservatingStudentInfoByAdmin(reservationId string, userId string, userType models.UserType) (string, error) {
	if len(userId) == 0 {
		return "", errors.New("请先登录")
	} else if userType != models.ADMIN {
		return "", errors.New("权限不足")
	} else if len(reservationId) == 0 {
		return "", errors.New("咨询已下架")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return "", errors.New("管理员账户出错,请联系技术支持")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil {
		return "", errors.New("咨询已下架")
	} else if reservation.Status != models.RESERVATED {
		return "", errors.New("咨询未被预约")
	}
	student, err := models.GetStudentById(reservation.StudentId)
	if err != nil {
		return "", errors.New("学生未注册")
	}
	filename := "student_" + student.Username + "_" +
		reservation.StartTime.In(utils.Location).Format(utils.DATE_PATTERN) + utils.CsvSuffix
	if err = workflow.ExportStudent(student, reservation, filename); err != nil {
		return "", err
	}
	return "/" + utils.ExportFolder + filename, nil
}

// 管理员导出学生信息
func (al *AdminLogic) ExportStudentByAdmin(studentId string, userId string, userType models.UserType) (string, error) {
	if len(userId) == 0 {
		return "", errors.New("请先登录")
	} else if userType != models.ADMIN {
		return "", errors.New("权限不足")
	} else if len(studentId) == 0 {
		return "", errors.New("学生未注册")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return "", errors.New("管理员账户出错,请联系技术支持")
	}
	student, err := models.GetStudentById(studentId)
	if err != nil {
		return "", errors.New("学生未注册")
	}
	filename := "student_" + student.Username + "_" + utils.GetNow().Format(utils.DATE_PATTERN) + utils.CsvSuffix
	if err = workflow.ExportStudent(student, nil, filename); err != nil {
		return "", err
	}
	return "/" + utils.ExportFolder + filename, nil
}

func (al *AdminLogic) SetStudentByAdmin(reservationId string, studentUsername string,
	userId string, userType models.UserType) (*models.Reservation, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	} else if len(reservationId) == 0 {
		return nil, errors.New("咨询已下架")
	} else if len(studentUsername) == 0 {
		return nil, errors.New("学生学号为空")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错，请联系技术支持")
	}
	student, err := models.GetStudentByUsername(studentUsername)
	if err != nil {
		return nil, errors.New("学生未注册")
	}
	reservation, err := models.GetReservationById(reservationId)
	if err != nil || reservation.Status == models.DELETED {
		return nil, errors.New("咨询已下架")
	} else if reservation.Status != models.AVAILABLE {
		return nil, errors.New("咨询已被预约")
	}
	// 更新咨询信息
	reservation.StudentId = student.Id.Hex()
	reservation.Status = models.RESERVATED
	reservation.StudentExpectation = models.StudentExpectation{
		Time: utils.GetNow(),
	}
	if models.UpsertReservation(reservation) != nil {
		return nil, errors.New("获取数据失败")
	}
	return reservation, nil
}

// 管理员查询学生信息
func (al *AdminLogic) QueryStudentInfoByAdmin(studentUsername string,
	userId string, userType models.UserType) (*models.Student, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	} else if len(studentUsername) == 0 {
		return nil, errors.New("学号为空")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错,请联系技术支持")
	}
	student, err := models.GetStudentByUsername(studentUsername)
	if err != nil || student.UserType != models.STUDENT {
		return nil, errors.New("学生未注册")
	}
	return student, nil
}

// 管理员导出咨询
func (al *AdminLogic) ExportReservationsByAdmin(reservationIds []string, userId string, userType models.UserType) (string, error) {
	if len(userId) == 0 {
		return "", errors.New("请先登录")
	} else if userType != models.ADMIN {
		return "", errors.New("权限不足")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return "", errors.New("管理员账户出错,请联系技术支持")
	}
	var reservations []*models.Reservation
	for _, reservationId := range reservationIds {
		reservation, err := models.GetReservationById(reservationId)
		if err != nil || reservation.Status != models.RESERVATED {
			continue
		}
		reservations = append(reservations, reservation)
	}
	filename := "export_" + time.Now().In(utils.Location).Format(utils.DATE_PATTERN) + utils.CsvSuffix
	if len(reservations) == 0 {
		return "", nil
	}
	if err = workflow.ExportReservations(reservations, filename); err != nil {
		return "", err
	}
	return "/" + utils.ExportFolder + filename, nil
}

// 查找咨询师
// 查找顺序:全名 > 工号 > 手机号
func (al *AdminLogic) SearchTeacherByAdmin(teacherFullname string, teacherUsername string, teacherMobile string,
	userId string, userType models.UserType) (*models.Teacher, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	} else if userType != models.ADMIN {
		return nil, errors.New("权限不足")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil || admin.UserType != models.ADMIN {
		return nil, errors.New("管理员账户出错,请联系技术支持")
	}
	if len(teacherFullname) != 0 {
		teacher, err := models.GetTeacherByFullname(teacherFullname)
		if err == nil {
			return teacher, nil
		}
	}
	if len(teacherUsername) != 0 {
		teacher, err := models.GetTeacherByUsername(teacherUsername)
		if err == nil {
			return teacher, nil
		}
	}
	if len(teacherMobile) != 0 {
		teacher, err := models.GetTeacherByMobile(teacherMobile)
		if err == nil {
			return teacher, nil
		}
	}
	return nil, errors.New("用户不存在")
}
