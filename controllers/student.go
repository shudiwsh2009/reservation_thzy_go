package controllers

import (
	"github.com/shudiwsh2009/reservation_thzy_go/buslogic"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"net/http"
	"strconv"
	"time"
)

func ViewReservationsByStudent(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	var result = map[string]interface{}{"state": "SUCCESS"}
	var rl = buslogic.ReservationLogic{}
	var ul = buslogic.UserLogic{}

	student, err := ul.GetStudentById(userId)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	var studentJson = make(map[string]interface{})
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

	reservations, err := rl.GetReservationsByStudent(userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	var array = make([]interface{}, 0)
	for _, res := range reservations {
		resJson := make(map[string]interface{})
		resJson["reservation_id"] = res.Id.Hex()
		resJson["start_time"] = res.StartTime.In(utils.Location).Format(utils.TIME_PATTERN)
		resJson["end_time"] = res.EndTime.In(utils.Location).Format(utils.TIME_PATTERN)
		if teacher, err := ul.GetTeacherById(res.TeacherId); err == nil {
			resJson["teacher_fullname"] = teacher.Fullname
		}
		if res.Status == models.AVAILABLE {
			resJson["status"] = models.AVAILABLE.String()
		} else if res.Status == models.RESERVATED && res.StartTime.Before(time.Now().In(utils.Location)) {
			resJson["status"] = models.FEEDBACK.String()
		} else {
			resJson["status"] = models.RESERVATED.String()
		}
		array = append(array, resJson)
	}
	result["reservations"] = array

	return result
}

func MakeReservationByStudent(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	reservationId := r.PostFormValue("reservation_id")
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
	problem := r.PostFormValue("problem")
	expectation := r.PostFormValue("expectation")
	expectedTime := r.PostFormValue("expected_time")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var sl = buslogic.StudentLogic{}
	var ul = buslogic.UserLogic{}

	var reservationJson = make(map[string]interface{})
	knowingMethodsInt := []int{}
	for _, m := range knowingMethods {
		if mi, err := strconv.Atoi(m); err == nil {
			knowingMethodsInt = append(knowingMethodsInt, mi)
		}
	}
	reservation, err := sl.MakeReservationByStudent(reservationId, fullname, gender, college, mobile, email, hasCareerConsulting,
		emergencyPerson, emergencyMobile, age, birthday, ethnic, enterYear, sourcePlace, originalSchool, originalMajor,
		marriage, health, fatherJob, motherJob, hasBrotherOrSister, brotherAge, brotherJob, hasMentalConsulting,
		otherConsultingNow, workingExperience, workingPeriod, knowingMethodsInt, problem, expectation, expectedTime,
		userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	reservationJson["start_time"] = reservation.StartTime.In(utils.Location).Format(utils.TIME_PATTERN)
	reservationJson["end_time"] = reservation.EndTime.In(utils.Location).Format(utils.TIME_PATTERN)
	if teacher, err := ul.GetTeacherById(reservation.TeacherId); err == nil {
		reservationJson["teacher_fullname"] = teacher.Fullname
	}
	result["reservation"] = reservationJson

	return result
}

func GetFeedbackByStudent(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	reservationId := r.PostFormValue("reservation_id")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var sl = buslogic.StudentLogic{}

	var feedbackJson = make(map[string]interface{})
	reservation, err := sl.GetFeedbackByStudent(reservationId, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}
	feedbackJson["consulting_count"] = reservation.StudentFeedback.ConsultingCount
	feedbackJson["scores"] = reservation.StudentFeedback.Scores
	feedbackJson["help"] = reservation.StudentFeedback.Help
	feedbackJson["drawback"] = reservation.StudentFeedback.Drawback
	result["feedback"] = feedbackJson

	return result
}

func SubmitFeedbackByStudent(w http.ResponseWriter, r *http.Request, userId string, userType models.UserType) interface{} {
	reservationId := r.PostFormValue("reservation_id")
	consultingCount := r.PostFormValue("consulting_count")
	r.ParseForm()
	scores := []string(r.Form["scores"])
	help := r.PostFormValue("help")
	drawback := r.PostFormValue("drawback")

	var result = map[string]interface{}{"state": "SUCCESS"}
	var sl = buslogic.StudentLogic{}

	scoresInt := []int{}
	for _, p := range scores {
		if pi, err := strconv.Atoi(p); err == nil {
			scoresInt = append(scoresInt, pi)
		}
	}
	_, err := sl.SubmitFeedbackByStudent(reservationId, consultingCount, scores, help, drawback, userId, userType)
	if err != nil {
		ErrorHandler(w, r, err)
		return nil
	}

	return result
}
