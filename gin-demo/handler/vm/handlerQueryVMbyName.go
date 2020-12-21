package vm

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 通过 id 获取虚拟机summary内容
// @version 1.0
// @Accept application/x-json-stream
// @Param name path string true "name"
// @Success 200 {object} models.Result 成功后返回值
// @Router /api/queryvmbyname/{name} [get]

func QueryVMbyName(context *gin.Context) {
	vmname := context.Param("name")
	vmsummary := &models.VmSummary{}
	mysummary := vmsummary.QueryVM(vmname)
	mysummary.Name = vmname
	context.JSON(http.StatusOK, gin.H{
		"result": models.Result{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    mysummary,
		},
	})
}
