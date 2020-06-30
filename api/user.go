package api

import "github.com/gin-gonic/gin"

// @Tags 用户
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept  json
// @Produce json
// @Param   num  query int false "分页参数"
// @Param   page query int false "分页参数"
// @Success 200 {object} ResOK
// @Router /users [get]
func UserGetList(c *gin.Context) {
	app := Gin{c}

	type user struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}

	users := []user{
		{Name: "王花花", Age: 16, Gender: "女"},
		{Name: "李栓蛋", Age: 17, Gender: "男"},
	}

	ret := gin.H{
		"list":  users,
		"count": len(users),
	}

	app.ResOK(ret)

	return
}
