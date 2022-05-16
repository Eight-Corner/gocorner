package controller

import (
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "fail"})
		return
	}
	defer func() {
		if err := recover(); err != nil {
			c.JSON(200, gin.H{"code": 500, "msg": "500 Runtime Error"})
			return
		}
	}()

	response := "response ....."

	c.JSON(200, gin.H{"code": 200, "msg": "success", "result": response})

}

/*	if err := c.Bind(&req); err != nil {
		panic(err)
	}
	if err := recover(); err != nil {
		str := fmt.Sprintf("%v", err)
		if validatorError := strings.Contains(str, `User`); validationError {
			c.JSON(200, gin.H{"code": 400, "message": "Not Found Body(JSON)"})
			return
		}
		c.JSON(200, gin.H{"code": 500, "message": "500 Run time Error"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "message": "success"})*/
