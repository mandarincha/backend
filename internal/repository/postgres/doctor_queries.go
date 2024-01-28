package postgres


const (
	GetAllTypes=`
		select distinct type from doctor 
	`
	GetDoctorByType=`
	select id, name,workplace, work_hours,photo,phone_number,type from doctor where type=$1
	`
	GetDoctorById=`
		select id,name,workplace, work_hours,photo ,phone_number,type from doctor where name like $1 
	`
	GetRating=`
	SELECT 	CAST(SUM(score) AS DECIMAL) / CAST(COUNT(id) AS DECIMAL)
FROM scores
WHERE user_id = $1;
 `

)