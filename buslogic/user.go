package buslogic

import (
	"errors"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"strings"
)

const (
	AdminDefaultPassword   = "THZYFZZX"
	TeacherDefaultPassword = "thzyfzzx"
)

type UserLogic struct {
}

// 学生登录
func (ul *UserLogic) StudentLogin(username string, password string) (*models.Student, error) {
	if len(username) == 0 {
		return nil, errors.New("用户名为空")
	} else if len(password) == 0 {
		return nil, errors.New("密码为空")
	}
	student, err := models.GetStudentByUsername(username)
	if err == nil && strings.EqualFold(student.Password, password) {
		return student, nil
	}
	return nil, errors.New("用户名或密码不正确")
}

// 咨询师登录
func (ul *UserLogic) TeacherLogin(username string, password string) (*models.Teacher, error) {
	if len(username) == 0 {
		return nil, errors.New("用户名为空")
	} else if len(password) == 0 {
		return nil, errors.New("密码为空")
	}
	teacher, err := models.GetTeacherByUsername(username)
	if err == nil && (strings.EqualFold(password, teacher.Password) ||
		(teacher.UserType == models.TEACHER && strings.EqualFold(password, TeacherDefaultPassword))) {
		return teacher, nil
	}
	return nil, errors.New("用户名或密码不正确")
}

// 管理员登录
func (ul *UserLogic) AdminLogin(username string, password string) (*models.Admin, error) {
	if len(username) == 0 {
		return nil, errors.New("用户名为空")
	} else if len(password) == 0 {
		return nil, errors.New("密码为空")
	}
	admin, err := models.GetAdminByUsername(username)
	if err == nil && (strings.EqualFold(password, admin.Password) ||
		(admin.UserType == models.ADMIN && strings.EqualFold(password, AdminDefaultPassword))) {
		return admin, nil
	}
	return nil, errors.New("用户名或密码不正确")
}

// 学生注册
func (ul *UserLogic) StudentRegister(username string, password string, fullname string, gender string, college string,
	mobile string, email string, hasCareerConsulting bool, emergencyPerson string, emergencyMobile string,
	age string, birthday string, ethnic string, enterYear string, sourcePlace string, originalSchool string,
	originalMajor string, marriage string, health string, fatherJob string, motherJob string, hasBrotherOrSister bool,
	brotherAge string, brotherJob string, hasMentalConsulting bool, otherConsultingNow string, workingExperience int,
	workingPeriod string, knowingMethods []int) (*models.Student, error) {
	if len(username) == 0 {
		return nil, errors.New("用户名为空")
	} else if len(password) == 0 {
		return nil, errors.New("密码为空")
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
	} else if len(emergencyPerson) == 0 {
		return nil, errors.New("紧急联系人为空")
	} else if !utils.IsMobile(emergencyMobile) {
		return nil, errors.New("紧急联系人手机号格式不正确")
	}
	if !utils.IsStudentId(username) {
		return nil, errors.New("请用学号注册")
	}
	if _, err := models.GetStudentByUsername(username); err == nil {
		return nil, errors.New("该学号已被注册")
	}
	newStudent, err := models.AddFullStudent(username, password, fullname, gender, college, mobile, email,
		hasCareerConsulting, emergencyPerson, emergencyMobile)
	if err != nil {
		return nil, errors.New("注册失败，请联系管理员")
	}
	newStudent.Age = age
	newStudent.Birthday = birthday
	newStudent.Ethnic = ethnic
	newStudent.EnterYear = enterYear
	newStudent.SourcePlace = sourcePlace
	newStudent.OriginalSchool = originalSchool
	newStudent.OriginalMajor = originalMajor
	newStudent.Marriage = marriage
	newStudent.Health = health
	newStudent.FatherJob = fatherJob
	newStudent.MotherJob = motherJob
	newStudent.HasBrotherOrSister = hasBrotherOrSister
	newStudent.BrotherAge = brotherAge
	newStudent.BrotherJob = brotherJob
	newStudent.HasMentalConsulting = hasMentalConsulting
	newStudent.OtherConsultingNow = otherConsultingNow
	newStudent.WorkingExperience = workingExperience
	newStudent.WorkingPeriod = workingPeriod
	newStudent.KnowingMethods = knowingMethods
	if models.UpsertStudent(newStudent) != nil {
		return nil, errors.New("获取数据失败")
	}
	return newStudent, nil
}

// 获取学生
func (ul *UserLogic) GetStudentById(userId string) (*models.Student, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	}
	student, err := models.GetStudentById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	}
	return student, nil
}

// 获取咨询师
func (ul *UserLogic) GetTeacherByUsername(username string) (*models.Teacher, error) {
	if len(username) == 0 {
		return nil, errors.New("请先登录")
	}
	teacher, err := models.GetTeacherByUsername(username)
	if err != nil {
		return nil, errors.New("请先登录")
	}
	return teacher, nil
}

// 获取咨询师
func (ul *UserLogic) GetTeacherById(userId string) (*models.Teacher, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	}
	teacher, err := models.GetTeacherById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	}
	return teacher, nil
}

// 获取管理员
func (ul *UserLogic) GetAdminById(userId string) (*models.Admin, error) {
	if len(userId) == 0 {
		return nil, errors.New("请先登录")
	}
	admin, err := models.GetAdminById(userId)
	if err != nil {
		return nil, errors.New("请先登录")
	}
	return admin, nil
}
