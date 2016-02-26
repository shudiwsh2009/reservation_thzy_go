use reservation_thzy
show collections
db.student.createIndex({"username": 1, "user_type": 1})
db.teacher.createIndex({"username": 1})
db.teacher.createIndex({"fullname": 1})
db.teacher.createIndex({"mobile": 1})
db.admin.createIndex({"username": 1})
db.reservation.createIndex({"student_id": 1, "status": 1, "start_time": 1})
db.reservation.createIndex({"start_time": 1, "status": 1})
