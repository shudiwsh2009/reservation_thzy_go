package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserType int

const (
	UNKNOWN UserType = iota
	STUDENT
	TEACHER
	ADMIN
)

var userTypes = [...]string{
	"UNKNOWN",
	"STUDENT",
	"TEACHER",
	"ADMIN",
}

func (ut UserType) String() string {
	return userTypes[ut]
}

func (ut UserType) IntStr() string {
	return fmt.Sprintf("%d", ut)
}

type Student struct {
	Id         bson.ObjectId `bson:"_id"`
	CreateTime time.Time     `bson:"create_time"`
	UpdateTime time.Time     `bson:"update_time"`
	Username   string        `bson:"username"` // Indexed
	Password   string        `bson:"password"`
	UserType   UserType      `bson:"user_type"`

	Fullname            string `bson:"fullname"`
	Gender              string `bson:"gender"`
	College             string `bson:"college"`
	Mobile              string `bson:"mobile"`
	Email               string `bson:"email"`
	HasCareerConsulting string `bson:"has_career_consulting"`
	EmergencyPerson     string `bson:"emergency_person"`
	EmergencyMobile     string `bson:"emergency_mobile"`

	Age                 string `bson:"age"`
	Birthday            string `bson:"birthday"`
	Ethnic              string `bson:"ethnic`
	EnterYear           string `bson:"enter_year"`
	SourcePlace         string `bson:"source_place"`
	OriginalSchool      string `bson:"original_school"`
	OriginalMajor       string `bson:"original_major"`
	Marriage            string `bson:"marrige"`
	Health              string `bson:"health"`
	FatherJob           string `bson:"father_job"`
	MotherJob           string `bson:"mother_job"`
	HasBrotherOrSister  string `bson:"has_brother_or_sister"`
	BrotherAge          string `bson:"brother_age"`
	BrotherJob          string `bson:"brother_job"`
	HasMentalConsulting string `bson:"has_mental_consulting"`
	OtherConsultingNow  string `bson:"other_consulting_now"`
	WorkingExperience   string `bson:"working_experience"`
	WorkingPeriod       string `bson:"working_period"`
	KnowingMethods      []int  `bson:"knowing_methods"`
}

type Teacher struct {
	Id         bson.ObjectId `bson:"_id"`
	CreateTime time.Time     `bson:"create_time"`
	UpdateTime time.Time     `bson:"update_time"`
	Username   string        `bson:"username"` // Indexed
	Password   string        `bson:"password"`
	Fullname   string        `bson:"fullname"`
	Mobile     string        `bson:"mobile"`
	UserType   UserType      `bson:"user_type"`
}

type Admin struct {
	Id         bson.ObjectId `bson:"_id"`
	CreateTime time.Time     `bson:"create_time"`
	UpdateTime time.Time     `bson:"update_time"`
	Username   string        `bson:"username"` // Indexed
	Password   string        `bson:"password"`
	UserType   UserType      `bson:"user_type"`
}

var KnowingMethods = []string{"与导师、辅导员交谈", "与同学交谈", "与家人交谈", "与相关行业在职人员交谈", "参加宣讲会、招聘会",
	"参加相关讲座和工作坊", "阅读报纸、书籍中的就业信息", "浏览求职网站", "选修职业辅导课程", "到就业指导中心寻求帮助",
	"向院系寻求相关就业资料", "参加有关的职业生涯团体", "参加过其他学生团体", "在用人单位实习或者兼职"};