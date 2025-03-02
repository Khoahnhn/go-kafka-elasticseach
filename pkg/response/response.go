//package response
//
//import (
//	"github.com/gin-gonic/gin"
//)
//
//// JSONResponse chuẩn hóa response format chung
//func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
//	c.JSON(status, gin.H{
//		"status":  status,
//		"message": message,
//		"data":    data,
//	})
//}
