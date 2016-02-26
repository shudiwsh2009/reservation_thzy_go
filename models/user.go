package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserType int

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

	Fullname         string `bson:"fullname"`
	Gender           string `bson:"gender"`
	College          string `bson:"college"`
	Mobile           string `bson:"mobile"`
	Email            string `bson:"email"`
	CareerConsulting string `bson:"career_consulting"`
	EmergencyPerson  string `bson:"emergency_person"`
	EmergencyMobile  string `bson:"emergency_mobile"`

	Age                int      `bson:"age"`
	Birthday           string   `bson:"birthday"`
	Ethnic             string   `bson:"ethnic`
	EnterYear          string   `bson:"enter_year"`
	SourcePlace        string   `bson:"source_place"`
	OriginalSchool     string   `bson:"original_school"`
	OriginalMajor      string   `bson:"original_major"`
	Marriage           string   `bson:"marrige"`
	Health             string   `bson:"health"`
	FatherJob          string   `bson:"father_job"`
	MotherJob          string   `bson:"mother_job"`
	BrotherOrSister    string   `bson:"brother_or_sister"`
	MentalConsulting   string   `bson:"mental_consulting"`
	OtherConsultingNow string   `bson:"other_consulting_now"`
	WorkingExperience  string   `bson:"working_experience"`
	KnowingMethods     []string `bson:"knowing_methods"`
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
