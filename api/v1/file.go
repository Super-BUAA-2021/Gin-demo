package v1

import (
	"fmt"
	"github.com/Super-BUAA-2021/Gin-demo/global"
	"github.com/Super-BUAA-2021/Gin-demo/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// UploadFile
// @Summary      上传文件
// @Description  上传一个文件测试
// @Tags         其他模块
// @Accept       multipart/form-data
// @Produce      json
// @Param        file     formData  file                   true  "文件"
// @Param        data     body      response.UploadFileQ  true  "name 文件"
// @Success      200      {object}  response.CommonA          "是否成功，返回信息"
// @Router       /api/v1/resource/upload [post]
func UploadFile(c *gin.Context) {
	// 获取请求数据
	var data response.UploadFileQ
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, response.CommonA{Success: false, Message: "请求参数非法" + err.Error(), Code: 400})
		return
	}
	// 保存文件的其余操作
	fmt.Println(data.Name)
	fmt.Println(data.File.Size)
	if err := c.SaveUploadedFile(data.File, filepath.Join(global.VP.GetString("root_path"), "resource", data.Name)); err != nil {
		c.JSON(http.StatusOK, response.CommonA{Success: true, Message: "上传文件失败," + err.Error(), Code: 402})
		return
	}
	c.JSON(http.StatusOK, response.CommonA{Success: true, Message: "上传文件成功", Code: 200})
}
