package http

import (
	articleUc "github.com/aryahmph/kumparan-assesment/internal/usecase/article"
	"github.com/gin-gonic/gin"
)

type HTTPArticleDelivery struct {
	articleUsecase articleUc.Usecase
}

func NewHTTPArticleDelivery(router *gin.RouterGroup, articleUsecase articleUc.Usecase) HTTPArticleDelivery {
	h := HTTPArticleDelivery{articleUsecase: articleUsecase}
	router.GET("", h.list)
	router.POST("", h.create)
	return h
}
