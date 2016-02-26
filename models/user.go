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
