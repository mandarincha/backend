package postgres

import (
	"context"
	"database/sql"
	"testDeployment/internal/domain"
	"testDeployment/internal/repository"
	"testDeployment/pkg/Bot"
)

type doctor struct{
	db *sql.DB
	bot Bot.Bot
}
func NewDoctorRepository(db *sql.DB,bot Bot.Bot) repository.IDoctorRepository{
	return &doctor{
		db:db,
		bot:bot,
	}
}
func (r *doctor) GetAllDoctor(ctx context.Context)([]*domain.DoctorByType,error){
	var doctors []*domain.DoctorByType
	rows,err:=r.db.QueryContext(ctx,GetAllTypes)
	if err!=nil{
		r.bot.SendErrorNotification(err)
		return nil,err
	}
	for rows.Next(){
		var tip *domain.DoctorByType
		rows.Scan(
			&tip.Type,
		)
		doctors=append(doctors, tip)
	}
	for i,tip:=range doctors{
		rows,err:=r.db.QueryContext(ctx,tip.Type)
		if err!=nil{
			r.bot.SendErrorNotification(err)
			return nil,err
		}
		for rows.Next(){
			var doctor *domain.Doctor
			rows.Scan(
				&doctor.Id,
				&doctor.Name,
				&doctor.Workplace,
				&doctor.WorkHours,
				&doctor.Photo,
			)
			doctors[i].Doctor=append(doctors[i].Doctor,doctor)
		}
	}
return doctors ,nil
}
func (r *doctor) GetById(ctx context.Context,id string) (*domain.DoctorWithType,error){
	var doctor *domain.DoctorWithType
	err:=r.db.QueryRowContext(ctx,GetDoctorById,id).Scan(
		&doctor.Id,
		&doctor.Name,
		&doctor.Workplace,
		&doctor.WorkHours,
		&doctor.Photo,
		&doctor.PhoneNumber,
	)
	if err!=nil{
		r.bot.SendErrorNotification(err)
		return nil,err
	}
	return doctor,nil
}