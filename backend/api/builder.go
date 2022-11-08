package api

import (
	"backend/builder"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(router *gin.Engine, services *builder.Services) {
	router.GET("/total-overlap", getTotalIntersection(services))
	router.GET("/common-zipcodes", getCommonZipcodes(services))
	router.GET("/overlap-between-groups", getOverlapBetweenGroups(services))
	router.GET("/all-users", getAllUsersWithNumberOfElements(services))
}
