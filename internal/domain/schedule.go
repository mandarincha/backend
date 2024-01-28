package domain


type Schedule struct{
	Id string `json:"id"`
	UserId string 
	Name string `json:"name"`
	PhoneNumber string `json:"phone-number"`
	Date string `json:"date"`
	FromTime string `json:"from-time"`
	ToTime string `json:"to-time"`
	Purpose string `json:"purpose"`
	Agree string `json:"agree"`
	IsAgree bool `json:"is-agree"`
	DoctorID string `json:"doctor-id"`
	Overview string 
}
type ScheduleResponse struct{
	Id string `json:"id"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone-number"`
	Date string `json:"date"`
	FromTime string `json:"from-time"`
	ToTime string `json:"to-time"`
	Purpose string `json:"purpose"`
	IsAgree bool `json:"is-agree"`
	Doctor Doctor
}