package http

import (
	httpCommon "github.com/aryahmph/kumparan-assesment/common/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPArticleDelivery) create(c *gin.Context) {
	var body httpCommon.AddArticle
	if err := c.BindJSON(&body); err != nil {
		return
	}

	id, err := d.articleUsecase.Create(c.Request.Context(), d.mapArticleBodyToModel(body))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}
