package controllers

import (
	"github.com/shudiwsh2009/reservation_thzy_go/buslogic"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"net/http"
	"strings"
	"time"
)

func ViewReservationsByTeacher(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	var result = map[string]interface{}{"state": "SUCCESS"}
	var ul = buslogic.UserLogic{}
	var rl = buslogic.ReservationLogic{}

	teacher, err := ul.GetTeacherById(userId)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	var teacherJson = make(map[string]interface{})
	teacherJson["teacher_fullname"] = teacher.Fullname
	teacherJson["teacher_mobile"] = teacher.Mobile
	result["teacher_info"] = teacherJson

	reservations, err := rl.GetReservationsByTeacher(userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	var array = make([]interface{}, 0)
	for _, res := range reservations {
		resJson := make(map[string]interface{})
		resJson["reservation_id"] = res.Id
		resJson["start_time"] = res.StartTime.In(utils.Location).Format(utils.TIME_PATTERN)
		resJson["end_time"] = res.EndTime.In(utils.Location).Format(utils.TIME_PATTERN)
		resJson["student_id"] = res.StudentId
		resJson["teacher_id"] = res.TeacherId
		if teacher, err := ul.GetTeacherById(res.TeacherId); err == nil {
			resJson["teacher_fullname"] = teacher.Fullname
			resJson["teacher_mobile"] = teacher.Mobile
		}
		if res.Status == models.AVAILABLE {
			resJson["status"] = models.AVAILABLE.String()
		} else if res.Status == models.RESERVATED && res.StartTime.Before(utils.GetNow()) {
			resJson["status"] = models.FEEDBACK.String()
		} else {
			resJson["status"] = models.RESERVATED.String()
		}
		array = append(array, resJson)
	}
	result["reservations"] = array

	return result
}

func AddReservationByTeacher(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	startTime := r.PostFormValue("start_time")
	endTime := r.PostFormValue("end_time")
	teacherFullname := r.PostFormValue("teacher_fullname")
	teacherMobile := r.PostFormValue("teacher_mobile")
	force := strings.EqualFold(r.PostFormValue("force"), "FORCE")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var tl = buslogic.TeacherLogic{}
	var ul = buslogic.UserLogic{}

	var reservationJson = make(map[string]interface{})
	reservation, err := tl.AddReservationByTeacher(startTime, endTime, teacherFullname, teacherMobile, force, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	reservationJson["reservation_id"] = reservation.Id
	reservationJson["start_time"] = reservation.StartTime.In(utils.Location).Format(utils.TIME_PATTERN)
	reservationJson["end_time"] = reservation.EndTime.In(utils.Location).Format(utils.TIME_PATTERN)
	reservationJson["student_id"] = reservation.StudentId
	reservationJson["teacher_id"] = reservation.TeacherId
	if teacher, err := ul.GetTeacherById(reservation.TeacherId); err == nil {
		reservationJson["teacher_username"] = teacher.Username
		reservationJson["teacher_fullname"] = teacher.Fullname
		reservationJson["teacher_mobile"] = teacher.Mobile
	}
	result["reservation"] = reservationJson

	return result
}

func EditReservationByTeacher(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	reservationId := r.PostFormValue("reservation_id")
	startTime := r.PostFormValue("start_time")
	endTime := r.PostFormValue("end_time")
	teacherFullname := r.PostFormValue("teacher_fullname")
	teacherMobile := r.PostFormValue("teacher_mobile")
	force := strings.EqualFold(r.PostFormValue("force"), "FORCE")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var tl = buslogic.TeacherLogic{}
	var ul = buslogic.UserLogic{}

	var reservationJson = make(map[string]interface{})
	reservation, err := tl.EditReservationByTeacher(reservationId, startTime, endTime, teacherFullname, teacherMobile,
		force, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	reservationJson["reservation_id"] = reservation.Id
	reservationJson["start_time"] = reservation.StartTime.In(utils.Location).Format(utils.TIME_PATTERN)
	reservationJson["end_time"] = reservation.EndTime.In(utils.Location).Format(utils.TIME_PATTERN)
	reservationJson["student_id"] = reservation.StudentId
	reservationJson["teacher_id"] = reservation.TeacherId
	if teacher, err := ul.GetTeacherById(reservation.TeacherId); err == nil {
		reservationJson["teacher_username"] = teacher.Username
		reservationJson["teacher_fullname"] = teacher.Fullname
		reservationJson["teacher_mobile"] = teacher.Mobile
	}
	result["reservation"] = reservationJson

	return result
}

func RemoveReservationByTeacher(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	r.ParseForm()
	reservationIds := []string(r.Form["reservation_ids"])

	var result = map[string]interface{}{"state": "SUCCESS"}
	var tl = buslogic.TeacherLogic{}

	removed, err := tl.RemoveReservationsByTeacher(reservationIds, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	result["removed_count"] = removed

	return result
}

func CancelReservationByTeacher(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	r.ParseForm()
	reservationIds := []string(r.Form["reservation_ids"])

	var result = map[string]interface{}{"state": "SUCCESS"}
	var tl = buslogic.TeacherLogic{}

	canceled, err := tl.CancelReservationsByTeacher(reservationIds, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	result["canceled_count"] = canceled

	return result
}

func GetStudentInfoByTeacher(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	studentId := r.PostFormValue("student_id")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var tl = buslogic.TeacherLogic{}
	var ul = buslogic.UserLogic{}

	var studentJson = make(map[string]interface{})
	student, reservations, err := tl.GetStudentInfoByTeacher(studentId, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	studentJson["student_id"] = student.Id.Hex()
	studentJson["student_username"] = student.Username
	studentJson["student_fullname"] = student.Fullname
	studentJson["gender"] = student.Gender
	studentJson["college"] = student.College
	studentJson["mobile"] = student.Mobile
	studentJson["email"] = student.Email
	studentJson["has_career_consulting"] = student.HasCareerConsulting
	studentJson["emergency_person"] = student.EmergencyPerson
	studentJson["emergency_mobile"] = student.EmergencyMobile
	studentJson["age"] = student.Age
	studentJson["birthday"] = student.Birthday
	studentJson["ethnic"] = student.Ethnic
	studentJson["enter_year"] = student.EnterYear
	studentJson["source_place"] = student.SourcePlace
	studentJson["original_school"] = student.OriginalSchool
	studentJson["original_major"] = student.OriginalMajor
	studentJson["marriage"] = student.Marriage
	studentJson["health"] = student.Health
	studentJson["father_job"] = student.FatherJob
	studentJson["mother_job"] = student.MotherJob
	studentJson["has_brother_or_sister"] = student.HasBrotherOrSister
	studentJson["brother_age"] = student.BrotherAge
	studentJson["brother_job"] = student.BrotherJob
	studentJson["has_mental_consulting"] = student.HasMentalConsulting
	studentJson["other_consulting_now"] = student.OtherConsultingNow
	studentJson["working_experience"] = student.WorkingExperience
	studentJson["working_period"] = student.WorkingPeriod
	studentJson["knowing_methods"] = student.KnowingMethods
	result["student_info"] = studentJson

	var reservationJson = make([]interface{}, 0)
	for _, res := range reservations {
		resJson := make(map[string]interface{})
		resJson["start_time"] = res.StartTime.In(utils.Location).Format(utils.TIME_PATTERN)
		resJson["end_time"] = res.EndTime.In(utils.Location).Format(utils.TIME_PATTERN)
		if res.Status == models.AVAILABLE {
			resJson["status"] = models.AVAILABLE.String()
		} else if res.Status == models.RESERVATED && res.StartTime.Before(utils.GetNow()) {
			resJson["status"] = models.FEEDBACK.String()
		} else {
			resJson["status"] = models.RESERVATED.String()
		}
		resJson["student_id"] = res.StudentId
		resJson["teacher_id"] = res.TeacherId
		if teacher, err := ul.GetTeacherById(res.TeacherId); err == nil {
			resJson["teacher_username"] = teacher.Username
			resJson["teacher_fullname"] = teacher.Fullname
			resJson["teacher_mobile"] = teacher.Mobile
		}
		resJson["student_problem"] = res.StudentExpectation.Problem
		resJson["student_expectation"] = res.StudentExpectation.Expectation
		resJson["student_expected_time"] = res.StudentExpectation.ExpectedTime
		resJson["student_feedback"] = res.StudentFeedback.ToJson()
		reservationJson = append(reservationJson, resJson)
	}
	result["reservations"] = reservationJson

	return result
}

func QueryStudentInfoByTeacher(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	studentUsername := r.PostFormValue("student_username")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var tl = buslogic.TeacherLogic{}
	var ul = buslogic.UserLogic{}

	var studentJson = make(map[string]interface{})
	student, reservations, err := tl.QueryStudentInfoByTeacher(studentUsername, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	studentJson["student_id"] = student.Id.Hex()
	studentJson["student_username"] = student.Username
	studentJson["student_fullname"] = student.Fullname
	studentJson["gender"] = student.Gender
	studentJson["college"] = student.College
	studentJson["mobile"] = student.Mobile
	studentJson["email"] = student.Email
	studentJson["has_career_consulting"] = student.HasCareerConsulting
	studentJson["emergency_person"] = student.EmergencyPerson
	studentJson["emergency_mobile"] = student.EmergencyMobile
	studentJson["age"] = student.Age
	studentJson["birthday"] = student.Birthday
	studentJson["ethnic"] = student.Ethnic
	studentJson["enter_year"] = student.EnterYear
	studentJson["source_place"] = student.SourcePlace
	studentJson["original_school"] = student.OriginalSchool
	studentJson["original_major"] = student.OriginalMajor
	studentJson["marriage"] = student.Marriage
	studentJson["health"] = student.Health
	studentJson["father_job"] = student.FatherJob
	studentJson["mother_job"] = student.MotherJob
	studentJson["has_brother_or_sister"] = student.HasBrotherOrSister
	studentJson["brother_age"] = student.BrotherAge
	studentJson["brother_job"] = student.BrotherJob
	studentJson["has_mental_consulting"] = student.HasMentalConsulting
	studentJson["other_consulting_now"] = student.OtherConsultingNow
	studentJson["working_experience"] = student.WorkingExperience
	studentJson["working_period"] = student.WorkingPeriod
	studentJson["knowing_methods"] = student.KnowingMethods
	result["student_info"] = studentJson

	var reservationJson = make([]interface{}, 0)
	for _, res := range reservations {
		resJson := make(map[string]interface{})
		resJson["start_time"] = res.StartTime.In(utils.Location).Format(utils.TIME_PATTERN)
		resJson["end_time"] = res.EndTime.In(utils.Location).Format(utils.TIME_PATTERN)
		if res.Status == models.AVAILABLE {
			resJson["status"] = models.AVAILABLE.String()
		} else if res.Status == models.RESERVATED && res.StartTime.Before(utils.GetNow()) {
			resJson["status"] = models.FEEDBACK.String()
		} else {
			resJson["status"] = models.RESERVATED.String()
		}
		resJson["student_id"] = res.StudentId
		resJson["teacher_id"] = res.TeacherId
		if teacher, err := ul.GetTeacherById(res.TeacherId); err == nil {
			resJson["teacher_username"] = teacher.Username
			resJson["teacher_fullname"] = teacher.Fullname
			resJson["teacher_mobile"] = teacher.Mobile
		}
		resJson["student_problem"] = res.StudentExpectation.Problem
		resJson["student_expectation"] = res.StudentExpectation.Expectation
		resJson["student_expected_time"] = res.StudentExpectation.ExpectedTime
		resJson["student_feedback"] = res.StudentFeedback.ToJson()
		reservationJson = append(reservationJson, resJson)
	}
	result["reservations"] = reservationJson

	return result
}
