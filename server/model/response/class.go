package response

import "time"

// 学生选课列表返回（树状api）
type ClassListResponse struct {
	G map[string]*Group `json:"classes"`
}

type Group struct {
	ID      int  `json:"id"`
	Hours   int  `json:"hours"`
	Learned bool `json:"learned"`
	List    []Course
}

type Course struct {
	ID          uint      `json:"id"`
	TeacherName string    `json:"teacher_name"`
	Time        time.Time `json:"time"`
	Desc        string    `json:"desc"`
	ClassRoom   string    `json:"class_room"`
	Max         int       `json:"max"`
	Now         int       `json:"now"`
	Selected    bool      `json:"selected"`
}

type PersonalClassResponse struct {
	Crs []ClassRecord `json:"crs"`
}

type ClassRecord struct {
	ID        uint   `json:"id"`
	Cname     string `json:"cname"`
	Hours     int    `json:"hours"`
	Tname     string `json:"tname"`
	Desc      string `json:"desc"`
	Classroom string `json:"classroom"`
	Grade     uint   `json:"grade"`
}

type TeacherClassResponse struct {
	Tcrs []TeacherClassRecord `json:"classes"`
}

type TeacherClassRecord struct {
	Cid       uint      `json:"cid"`
	Cname     string    `json:"cname"`
	Ccredit   int       `json:"ccredit"`
	Desc      string    `json:"desc"`
	Time      time.Time `json:"time"`
	Selected  int       `json:"selected"`
	Classroom string    `json:"classroom"`
}

type TeacherClassStuResponse struct {
	Tcsrs []TeacherClassStuRecord `json:"students"`
}

type TeacherClassStuRecord struct {
	Username int64  `json:"username"`
	Name     string `json:"name"`
	Class    string `json:"class"`
	Grade    uint   `json:"grade"`
}

type TeacherRecord struct {
	Name string `json:"value"`
	TID  int64  `json:"tid"`
}

/*
{
    "code": 0,
    "data": {
        "classes": [
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
