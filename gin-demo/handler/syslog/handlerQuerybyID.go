package syslog

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary 通过 id 获取单条记录内容
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 {object} models.Result 成功后返回值
// @Router /api/querybyid/{id} [get]
func QuerybyID(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"result": models.Result{
				Code:    http.StatusBadRequest,
				Message: "id 不是 int 类型, id 转换失败",
				Data:    e.Error(),
			},
		})
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	syslog := &models.SysLogs{}
	row := syslog.QueryById()
	row.Id = int64(i)
	context.JSON(http.StatusOK, gin.H{
		"result": models.Result{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    row,
		},
	})
}
