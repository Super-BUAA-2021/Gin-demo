package v1

import (
	"fmt"
	"github.com/Super-BUAA-2021/Gin-demo/global"
	"github.com/Super-BUAA-2021/Gin-demo/model/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"path/filepath"
)

// UploadFile
// @Summary      上传文件
// @Description  上传一个文件测试
// @Tags         其他模块
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file                  true  "文件"
// @Param        name  formData  string                true  "name"
// @Success      200   {object}  response.UploadFileA  "是否成功，返回信息"
// @Router       /resource/upload [post]
func UploadFile(c *gin.Context) {
	// 获取请求数据
	var data response.UploadFileQ
	if err := c.ShouldBindWith(&data, binding.FormMultipart); err != nil {
		c.JSON(http.StatusOK, response.UploadFileA{Success: false, Message: "请求参数非法 : " + err.Error(), Code: 400})
		return
	}
	// 保存文件的其余操作
	fmt.Println(data.File.Size)

	if err := c.SaveUploadedFile(data.File, filepath.Join(global.VP.GetString("root_path"), "resource", data.Name)); err != nil {
		c.JSON(http.StatusOK, response.UploadFileA{Success: true, Message: "上传文件失败," + err.Error(), Code: 402})
		return
	}

	c.JSON(http.StatusOK, response.UploadFileA{Success: true, Message: "上传文件成功", Code: 200, Path: global.VP.GetString("server.backend_url") + "/resource/file/" + data.Name})
}
