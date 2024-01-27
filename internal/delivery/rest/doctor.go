package rest

import (
	"log"
	"testDeployment/internal/usecase"
	"testDeployment/pkg/Bot"

	"github.com/gin-gonic/gin"
)

type doctor struct{
	uc usecase.IDoctorUsecase
	bot Bot.Bot
}
func NewDoctorController(g *gin.RouterGroup,uc usecase.IDoctorUsecase,bot Bot.Bot) {
	handler:=doctor{
		uc:uc,
		bot:bot,
	}
	r:=g.Group("/doctors")
	r.GET("/getalldoctors",handler.GetAll)
	r.GET("/getonedoctor",handler.GetOneByID)
}
func (r *doctor) GetAll(c *gin.Context){
	doctors,err:=r.uc.GetAll(c)
	if err!=nil{
		c.JSON(200,gin.H{
			"error":err,
		})
		return
	}
	if doctors==nil{
		c.JSON(200,gin.H{
			"message":"no doctors yet",
		})
		return
	}
	c.JSON(200,doctors)
}
func (r *doctor) GetOneByID(c *gin.Context){
	id:=c.Query("id")
	log.Println(id)
	doctor,err:=r.uc.GetOneByID(c,id)
	if err!=nil{
		c.JSON(200,gin.H{
			"error":err,
		})
		return
	}
	if doctor==nil{
		c.JSON(200,gin.H{
			"message":"no doctor yet",
		})
		return
	}
	c.JSON(200,doctor)
}