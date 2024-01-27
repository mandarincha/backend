package domain
type Doctor struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Workplace string `json:"workplace"`
	WorkHours string `json:"work_hours"`
	Photo string `json:"photo"`
}
type DoctorWithType struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Workplace string `json:"workplace"`
	WorkHours string `json:"work_hours"`
	Type string `json:"type"`
	PhoneNumber string `json:"phone_number"`
	Photo string `json:"photo"`
}
type DoctorByType struct{
	ID string `json:"id"`
	Type string `json:"type"`
	Doctor []*Doctor
}
type DoctorList struct{
	TotalCount int `json:"totat_count"`
	TotalPages int `json:"total_pages"`
	Page int `json:"page"`
	Size int `json:"size"`
	HasMore bool `json:"has_more"`
	Doctor []*Doctor
}