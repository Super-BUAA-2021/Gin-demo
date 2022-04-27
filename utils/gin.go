package utils

import (
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, model interface{}) interface{} {
	if err := c.ShouldBindJSON(&model); err != nil {
		//_, file, line, _ := runtime.Caller(1)
		//global.LOG.Panic(file + "(line " + strconv.Itoa(line) + "): bind model error")
	}
	return model
}

func ShouldBindAndValid(c *gin.Context, model interface{}) error {
	if err := c.ShouldBind(&model); err != nil {
		return err
	}

	return nil
}
