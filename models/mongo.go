package models

import (
	"errors"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	Mongo *mgo.Database
)

// 为保证一致性，models中不用冗余存储，采用给字段构建索引的办法加快查询速度，以观后效

/**
Student
*/

func AddSimpleStudent(username string, password string) (*Student, error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("student")
	newStudent := &Student{
		Id:             bson.NewObjectId(),
		CreateTime:     utils.GetNow(),
		UpdateTime:     utils.GetNow(),
		Username:       username,
		Password:       password,
		UserType:       STUDENT,
		KnowingMethods: make([]int, 0),
	}
	if err := collection.Insert(newStudent); err != nil {
		return nil, err
	}
	return newStudent, nil
}

func AddFullStudent(username string, password string, fullname string, gender string, college string, mobile string,
	email string, hasCareerConsulting bool, emergencyPerson string, emergencyMobile string) (*Student, error) {
	if len(username) == 0 || len(password) == 0 || len(fullname) == 0 || len(gender) == 0 || len(college) == 0 ||
		!utils.IsMobile(mobile) || !utils.IsEmail(email) || len(emergencyPerson == 0) || !utils.IsMobile(emergencyMobile) {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("student")
	newStudent := &Student{
		Id:                  bson.NewObjectId(),
		CreateTime:          utils.GetNow(),
		UpdateTime:          utils.GetNow(),
		Username:            username,
		Password:            password,
		UserType:            STUDENT,
		Fullname:            fullname,
		Gender:              gender,
		College:             college,
		Mobile:              mobile,
		Email:               email,
		HasCareerConsulting: hasCareerConsulting,
		EmergencyPerson:     emergencyPerson,
		EmergencyMobile:     emergencyMobile,
		KnowingMethods:      make([]int, 0),
	}
	if err := collection.Insert(newStudent); err != nil {
		return nil, err
	}
	return newStudent, nil
}

func UpsertStudent(student *Student) error {
	if student == nil || !student.Id.Valid() {
		return errors.New("字段不合法")
	}
	collection := Mongo.C("student")
	student.UpdateTime = utils.GetNow()
	_, err := collection.UpsertId(student.Id, student)
	return err
}

func GetStudentById(studentId string) (*Student, error) {
	if len(studentId) == 0 || !bson.IsObjectIdHex(studentId) {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("student")
	student := &Student{}
	if err := collection.FindId(bson.ObjectIdHex(studentId)).One(student); err != nil {
		return nil, err
	}
	return student, nil
}

func GetStudentByUsername(username string) (*Student, error) {
	if len(username) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("student")
	student := &Student{}
	if err := collection.Find(bson.M{"username": username, "user_type": STUDENT}).One(student); err != nil {
		return nil, err
	}
	return student, nil
}

/**
Teacher
*/

func AddTeacher(username string, password string, fullname string, mobile string) (*Teacher, error) {
	if len(username) == 0 || len(password) == 0 || len(fullname) == 0 || len(mobile) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("teacher")
	newTeacher := &Teacher{
		Id:         bson.NewObjectId(),
		CreateTime: utils.GetNow(),
		UpdateTime: utils.GetNow(),
		Username:   username,
		Password:   password,
		Fullname:   fullname,
		Mobile:     mobile,
		UserType:   TEACHER,
	}
	if err := collection.Insert(newTeacher); err != nil {
		return nil, err
	}
	return newTeacher, nil
}

func UpsertTeacher(teacher *Teacher) error {
	if teacher == nil || !teacher.Id.Valid() {
		return errors.New("字段不合法")
	}
	collection := Mongo.C("teacher")
	teacher.UpdateTime = utils.GetNow()
	_, err := collection.UpsertId(teacher.Id, teacher)
	return err
}

func GetTeacherById(teacherId string) (*Teacher, error) {
	if len(teacherId) == 0 || !bson.IsObjectIdHex(teacherId) {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("teacher")
	teacher := &Teacher{}
	if err := collection.FindId(bson.ObjectIdHex(teacherId)).One(teacher); err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetTeacherByUsername(username string) (*Teacher, error) {
	if len(username) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("teacher")
	teacher := &Teacher{}
	if err := collection.Find(bson.M{"username": username}).One(teacher); err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetTeacherByFullname(fullname string) (*Teacher, error) {
	if len(fullname) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("teacher")
	teacher := &Teacher{}
	if err := collection.Find(bson.M{"fullname": fullname}).One(teacher); err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetTeacherByMobile(mobile string) (*Teacher, error) {
	if len(mobile) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("teacher")
	teacher := &Teacher{}
	if err := collection.Find(bson.M{"mobile": mobile}).One(teacher); err != nil {
		return nil, err
	}
	return teacher, nil
}

/**
Admin
*/

func AddAdmin(username string, password string) (*Admin, error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("admin")
	newAdmin := &Admin{
		Id:         bson.NewObjectId(),
		CreateTime: utils.GetNow(),
		UpdateTime: utils.GetNow(),
		Username:   username,
		Password:   password,
		UserType:   ADMIN,
	}
	if err := collection.Insert(newAdmin); err != nil {
		return nil, err
	}
	return newAdmin, nil
}

func UpsertAdmin(admin *Admin) error {
	if admin == nil || !admin.Id.Valid() {
		return errors.New("字段不合法")
	}
	collection := Mongo.C("admin")
	admin.UpdateTime = utils.GetNow()
	_, err := collection.UpsertId(admin.Id, admin)
	return err
}

func GetAdminById(adminId string) (*Admin, error) {
	if len(adminId) == 0 || !bson.IsObjectIdHex(adminId) {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("admin")
	admin := &Admin{}
	if err := collection.FindId(bson.ObjectIdHex(adminId)).One(admin); err != nil {
		return nil, err
	}
	return admin, nil
}

func GetAdminByUsername(username string) (*Admin, error) {
	if len(username) == 0 {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("admin")
	admin := &Admin{}
	if err := collection.Find(bson.M{"username": username}).One(admin); err != nil {
		return nil, err
	}
	return admin, nil
}

/**
Reservation
*/

func AddReservation(startTime time.Time, endTime time.Time, teacherId string) (*Reservation, error) {
	collection := Mongo.C("reservation")
	newReservation := &Reservation{
		Id:                 bson.NewObjectId(),
		CreateTime:         utils.GetNow(),
		UpdateTime:         utils.GetNow(),
		StartTime:          startTime,
		EndTime:            endTime,
		Status:             AVAILABLE,
		TeacherId:          teacherId,
		StudentExpectation: StudentExpectation{},
		StudentFeedback:    StudentFeedback{},
	}
	if err := collection.Insert(newReservation); err != nil {
		return nil, err
	}
	return newReservation, nil
}

func UpsertReservation(reservation *Reservation) error {
	if reservation == nil || !reservation.Id.Valid() {
		return errors.New("字段不合法")
	}
	collection := Mongo.C("reservation")
	reservation.UpdateTime = utils.GetNow()
	_, err := collection.UpsertId(reservation.Id, reservation)
	return err
}

func GetReservationById(id string) (*Reservation, error) {
	if len(id) == 0 || !bson.IsObjectIdHex(id) {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("reservation")
	reservation := &Reservation{}
	if err := collection.FindId(bson.ObjectIdHex(id)).One(reservation); err != nil {
		return nil, err
	}
	return reservation, nil
}

func GetReservationsByStudentId(studentId string) ([]*Reservation, error) {
	if len(studentId) == 0 || !bson.IsObjectIdHex(studentId) {
		return nil, errors.New("字段不合法")
	}
	collection := Mongo.C("reservation")
	var reservations []*Reservation
	if err := collection.Find(bson.M{"student_id": studentId,
		"status": bson.M{"$ne": DELETED}}).Sort("start_time").All(&reservations); err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationsBetweenTime(from time.Time, to time.Time) ([]*Reservation, error) {
	collection := Mongo.C("reservation")
	var reservations []*Reservation
	if err := collection.Find(bson.M{"start_time": bson.M{"$gte": from, "$lte": to},
		"status": bson.M{"$ne": DELETED}}).Sort("start_time").All(&reservations); err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservatedReservationsBetweenTime(from time.Time, to time.Time) ([]*Reservation, error) {
	collection := Mongo.C("reservation")
	var reservations []*Reservation
	if err := collection.Find(bson.M{"start_time": bson.M{"$gte": from, "$lte": to},
		"status": RESERVATED}).Sort("start_time").All(&reservations); err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationsAfterTime(from time.Time) ([]*Reservation, error) {
	collection := Mongo.C("reservation")
	var reservations []*Reservation
	if err := collection.Find(bson.M{"start_time": bson.M{"$gte": from},
		"status": bson.M{"$ne": DELETED}}).Sort("start_time").All(&reservations); err != nil {
		return nil, err
	}
	return reservations, nil
}
