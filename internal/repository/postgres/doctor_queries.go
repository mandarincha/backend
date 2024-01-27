package postgres


const (
	GetAllTypes=`
		select distinct type from doctor 
	`
	GetDoctorByType=`
	select id, name,workplace, work_hours,photo from doctor where type=$1
	`
	GetDoctorById=`
		select id, name,workplace, work_hours,photo ,phone_number from doctor where id=$1 
	`
)