package controllers

import (
	"github.com/shudiwsh2009/reservation_thzy_go/buslogic"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"net/http"
	"time"
	"strconv"
)

func StudentRegister(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fullname := r.PostFormValue("fullname")
	gender := r.PostFormValue("gender")
	college := r.PostFormValue("college")
	mobile := r.PostFormValue("mobile")
	email := r.PostFormValue("email")
	hasCareerConsulting := r.PostFormValue("has_career_consulting")
	emergencyPerson := r.PostFormValue("emergency_person")
	emergencyMobile := r.PostFormValue("emergency_mobile")
	age := r.PostFormValue("age")
	birthday := r.PostFormValue("birthday")
	ethnic := r.PostFormValue("ethnic")
	enterYear := r.PostFormValue("enter_year")
	sourcePlace := r.PostFormValue("source_place")
	originalSchool := r.PostFormValue("original_school")
	originalMajor := r.PostFormValue("original_major")
	marriage := r.PostFormValue("marriage")
	health := r.PostFormValue("health")
	fatherJob := r.PostFormValue("father_job")
	motherJob := r.PostFormValue("mother_job")
	hasBrotherOrSister := r.PostFormValue("has_brother_or_sister")
	brotherAge := r.PostFormValue("brother_age")
	brotherJob := r.PostFormValue("brother_job")
	hasMentalConsulting := r.PostFormValue("has_mental_consulting")
	otherConsultingNow := r.PostFormValue("other_consulting_now")
	workingExperience := r.PostFormValue("working_experience")
	workingPeriod := r.PostFormValue("working_period")
	r.ParseForm()
	knowingMethods := []string(r.Form["knowing_methods"])

	var result = map[string]interface{}{"state": "SUCCESS"}
	var ul = buslogic.UserLogic{}

	knowingMethodsInt := []int{}
	for _, m := range knowingMethods {
		if mi, err := strconv.Atoi(m); err == nil {
			knowingMethodsInt = append(knowingMethodsInt, mi)
		}
	}
	student, err := ul.StudentRegister(username, password, fullname, gender, college, mobile, email, hasCareerConsulting,
		emergencyPerson, emergencyMobile, age, birthday, ethnic, enterYear, sourcePlace, originalSchool, originalMajor,
		marriage, health, fatherJob, motherJob, hasBrotherOrSister, brotherAge, brotherJob, hasMentalConsulting,
		otherConsultingNow, workingExperience, workingPeriod, knowingMethodsInt)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    student.Id.Hex(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    student.Username,
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "user_type",
		Value:    student.UserType.IntStr(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	result["url"] = "/reservation/student"

	return result
}

func StudentLogin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var ul = buslogic.UserLogic{}

	student, err := ul.StudentLogin(username, password)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    student.Id.Hex(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    student.Username,
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "user_type",
		Value:    student.UserType.IntStr(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	result["url"] = "/reservation/student"

	return result
}

func TeacherLogin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var ul = buslogic.UserLogic{}

	teacher, err := ul.TeacherLogin(username, password)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    teacher.Id.Hex(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    teacher.Username,
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "user_type",
		Value:    teacher.UserType.IntStr(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	switch teacher.UserType {
	case models.TEACHER:
		result["url"] = "/reservation/teacher"
	default:
		result["url"] = "/reservation/entry"
	}

	return result
}

func AdminLogin(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var ul = buslogic.UserLogic{}

	admin, err := ul.AdminLogin(username, password)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    admin.Id.Hex(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    admin.Username,
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "user_type",
		Value:    admin.UserType.IntStr(),
		Path:     "/",
		Expires:  time.Now().Local().AddDate(1, 0, 0),
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: true,
	})
	switch admin.UserType {
	case models.ADMIN:
		result["url"] = "/reservation/admin"
	default:
		result["url"] = "/reservation/entry"
	}

	return result
}

func Logout(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	var result = map[string]interface{}{"state": "SUCCESS"}

	switch userType {
	case models.ADMIN:
		result["url"] = "/reservation/admin"
	case models.TEACHER:
		result["url"] = "/reservation/teacher"
	case models.STUDENT:
		result["url"] = "/reservation/student"
	default:
		result["url"] = "/reservation/entry"
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "user_id",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "username",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "user_type",
		Path:   "/",
		MaxAge: -1,
	})

	return result
}
