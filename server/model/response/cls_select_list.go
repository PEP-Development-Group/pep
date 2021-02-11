package response

import "time"

// 学生选课列表返回
type ClassList struct {
	ID        int    `json:"id"`
	ClassName string `json:"class_name"`
	Hours     int    `json:"hours"`
	l         []List
}

type List struct {
	ID          int       `json:"id"`
	TeacherName string    `json:"teacher_name"`
	Time        time.Time `json:"time"`
	Desc        string    `json:"desc"`
	ClassRoom   string    `json:"class_room"`
	Max         int       `json:"max"`
	Now         int       `json:"now"`
	Selected    bool      `json:"selected"`
}

/*
{
    "code": 0,
    "data": {
        "courseList": [
            {
                "id": 1,
                "name": "分光计",
                "hours": 2,
                "lists": [
                    {
                        "id": 990,
                        "name": "栾玉国",
                        "time": "2021-03-21 08:20:00",
                        "desc": "3周 周4 12节",
                        "classroom": "A201",
                        "max": 40,
                        "now": 20,
                        "selected": true
                    }
                ]
            }
        ]
    },
    "msg": "获取成功"
}
*/
