package controllers

import (
	"errors"
	"github.com/shudiwsh2009/reservation_thzy_go/buslogic"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"net/http"
	"net/url"
	"strings"
)

func ViewReservationsByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	var result = map[string]interface{}{"state": "SUCCESS"}
	var rl = buslogic.ReservationLogic{}
	var ul = buslogic.UserLogic{}

	reservations, err := rl.GetReservationsByAdmin(userId, userType)
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
			resJson["teacher_username"] = teacher.Username
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

func ViewMonthlyReservationsByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil || len(queryForm["from_date"]) == 0 {
		ErrorHandler(w, r, errors.New("参数错误"))
		return nil
	}
	fromTime := queryForm["from_date"][0]

	var result = map[string]interface{}{"state": "SUCCESS"}
	var rl = buslogic.ReservationLogic{}
	var ul = buslogic.UserLogic{}

	reservations, err := rl.GetReservationsMonthlyByAdmin(fromTime, userId, userType)
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
			resJson["teacher_username"] = teacher.Username
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

func AddReservationByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	startTime := r.PostFormValue("start_time")
	endTime := r.PostFormValue("end_time")
	teacherUsername := r.PostFormValue("teacher_username")
	teacherFullname := r.PostFormValue("teacher_fullname")
	teacherMobile := r.PostFormValue("teacher_mobile")
	force := strings.EqualFold(r.PostFormValue("force"), "FORCE")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}
	var ul = buslogic.UserLogic{}

	var reservationJson = make(map[string]interface{})
	reservation, err := al.AddReservationByAdmin(startTime, endTime, teacherUsername, teacherFullname,
		teacherMobile, force, userId, userType)
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

func EditReservationByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	reservationId := r.PostFormValue("reservation_id")
	startTime := r.PostFormValue("start_time")
	endTime := r.PostFormValue("end_time")
	teacherUsername := r.PostFormValue("teacher_username")
	teacherFullname := r.PostFormValue("teacher_fullname")
	teacherMobile := r.PostFormValue("teacher_mobile")
	force := strings.EqualFold(r.PostFormValue("force"), "FORCE")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}
	var ul = buslogic.UserLogic{}

	var reservationJson = make(map[string]interface{})
	reservation, err := al.EditReservationByAdmin(reservationId, startTime, endTime, teacherUsername,
		teacherFullname, teacherMobile, force, userId, userType)
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

func RemoveReservationsByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	r.ParseForm()
	reservationIds := []string(r.Form["reservation_ids"])

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	removed, err := al.RemoveReservationsByAdmin(reservationIds, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	result["removed_count"] = removed

	return result
}

func CancelReservationsByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	r.ParseForm()
	reservationIds := []string(r.Form["reservation_ids"])

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	canceled, err := al.CancelReservationsByAdmin(reservationIds, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	result["canceled_count"] = canceled

	return result
}

func GetReservatingStudentInfoByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	reservationId := r.PostFormValue("reservation_id")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	var studentJson = make(map[string]interface{})
	var reservationJson = make(map[string]interface{})
	student, reservation, err := al.GetReservatingStudentInfoByAdmin(reservationId, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	studentJson["id"] = student.Id.Hex()
	studentJson["username"] = student.Username
	studentJson["fullname"] = student.Fullname
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
	reservationJson["id"] = reservation.Id.Hex()
	reservationJson["problem"] = reservation.StudentExpectation.Problem
	reservationJson["expectation"] = reservation.StudentExpectation.Expectation
	reservationJson["expected_time"] = reservation.StudentExpectation.ExpectedTime
	reservationJson["time"] = reservation.StudentExpectation.Time.In(utils.Location).Format(utils.DATE_PATTERN)
	result["reservation_info"] = reservationJson
	result["student_feedback"] = reservation.StudentFeedback.ToJson()

	return result
}

func ResetStudentPasswordByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	studentId := r.PostFormValue("student_id")
	password := r.PostFormValue("password")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	_, err := al.ResetStudentPasswordByAdmin(studentId, password, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}

	return result
}

func DeleteStudentAccountByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	studentId := r.PostFormValue("student_id")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	err := al.DeleteStudentAccountByAdmin(studentId, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}

	return result
}

func ExportStudentByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	studentId := r.PostFormValue("student_id")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	url, err := al.ExportStudentByAdmin(studentId, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	result["url"] = url

	return result
}

func SetStudentByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	reservationId := r.PostFormValue("reservation_id")
	studentUsername := r.PostFormValue("student_username")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}
	var ul = buslogic.UserLogic{}

	var reservationJson = make(map[string]interface{})
	reservation, err := al.SetStudentByAdmin(reservationId, studentUsername, userId, userType)
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

func QueryStudentInfoByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	studentUsername := r.PostFormValue("student_username")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	var studentJson = make(map[string]interface{})
	student, err := al.QueryStudentInfoByAdmin(studentUsername, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	studentJson["id"] = student.Id.Hex()
	studentJson["username"] = student.Username
	studentJson["fullname"] = student.Fullname
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

	return result
}

func ExportReservationsByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	r.ParseForm()
	reservationIds := []string(r.Form["reservation_ids"])

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	url, err := al.ExportReservationsByAdmin(reservationIds, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	result["url"] = url

	return result
}

func ExportReservationsPeriodByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	fromDate := r.PostFormValue("from_date")
	toDate := r.PostFormValue("to_date")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	url, err := al.ExportReservationsPeriodByAdmin(fromDate, toDate, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	result["url"] = url

	return result
}

func SearchTeacherByAdmin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	teacherUsername := r.PostFormValue("teacher_username")
	teacherFullname := r.PostFormValue("teacher_fullname")
	teacherMoble := r.PostFormValue("teacher_mobile")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var al = buslogic.AdminLogic{}

	var teacherJson = make(map[string]interface{})
	teacher, err := al.SearchTeacherByAdmin(teacherFullname, teacherUsername, teacherMoble,
		userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	teacherJson["teacher_username"] = teacher.Username
	teacherJson["teacher_fullname"] = teacher.Fullname
	teacherJson["teacher_mobile"] = teacher.Mobile
	result["teacher"] = teacherJson

	return result
}
