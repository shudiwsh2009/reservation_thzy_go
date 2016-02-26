package models

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type ReservationStatus int

const (
	AVAILABLE ReservationStatus = 1 + iota
	RESERVATED
	FEEDBACK
	DELETED
	CLOSED
)

var reservationStatuses = [...]string{
	"AVAILABLE",
	"RESERVATED",
	"FEEDBACK",
	"DELETED",
	"CLOSED",
}

type StudentExpectation struct {
	Problem      string `bson:"problem"`
	Expectation  string `bson:"expectation"`
	ExpectedTime int    `bson:"expected_time"`
}

type StudentFeedback struct {
	ConsultingCount int       `bson:"consulting_count"`
	Scores          []int     `bson:"scores"`
	Help            string    `bson:"help"`
	Drawback        string    `bson:"drawback"`
	Time            time.Time `bson:"time"`
}

func (sf StudentFeedback) IsEmpty() bool {
	return sf.ConsultingCount == 0 || sf.Scores == nil || len(sf.Scores) == 0
}

func (sf StudentFeedback) ToJson() map[string]interface{} {
	var json = make(map[string]interface{})
	scores := ""
	for _, s := range sf.Scores {
		scores += strconv.Itoa(s) + " "
	}
	json["consulting_count"] = sf.ConsultingCount
	json["scores"] = scores
	json["help"] = sf.Help
	json["drawback"] = sf.Drawback
	return json
}

type Reservation struct {
	Id                 bson.ObjectId     `bson:"_id"`
	CreateTime         time.Time         `bson:"create_time"`
	UpdateTime         time.Time         `bson:"update_time"`
	StartTime          time.Time         `bson:"start_time"` // indexed
	EndTime            time.Time         `bson:"end_time"`
	Status             ReservationStatus `bson:"status"`
	TeacherId          string            `bson:"teacher_id"` // indexed
	StudentId          string            `bson:"student_id"` // indexed
	StudentExpectation StudentFeedback   `bson:"student_expectation"`
	StudentFeedback    StudentFeedback   `bson:"student_feedback"`
}

const CHECK_MESSAGE = "CHECK"
