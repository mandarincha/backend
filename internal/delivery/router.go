package delivery

import (
	request "backend/internal/delivery/http"
	"backend/internal/delivery/rest"
	
	"backend/pkg/Bot"

	"github.com/gin-gonic/gin"
)

func SetUp(
	g *gin.Engine,
	uc usecase.IUseCase,
	bot Bot.Bot,
	request request.CustomJSONRequester,
	) {
	SetUpHandlerV1(
			g.Group("/api/v1"),
			uc,
			bot,
			request,
		)

	}
func SetUpHandlerV1(
	group *gin.RouterGroup,
	uc usecase.IUseCase,
	bot Bot.Bot,
	request request.CustomJSONRequester,
	) {
		rest.NewController(
			group,
			uc.IOtherUseCase(),
			bot,
			request,
		
		)
		rest.NewFileController(
			group,
			bot,
		)
		rest.NewNewsController(
			group,
			bot,
			uc.INewsUsecase(),
		)
	}