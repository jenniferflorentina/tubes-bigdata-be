package app

import (
	"github.com/tubes-bigdata/domain"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapUrls()
	domain.ConnDB()
	router.Run(":8083")
}
