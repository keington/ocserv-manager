package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/start", startOpenConnect)
		api.POST("/stop", stopOpenConnect)
		api.POST("/user/add/:username/:password", addUser)
		api.DELETE("/user/delete/:username", deleteUser)
		api.GET("/user/list", getUserList)
	}
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server")
		return
	}
}

// 启动openConnect程序
func startOpenConnect(c *gin.Context) {
	out, err := exec.Command("bash", "-c", "sudo openconnect --user=user vpn.server.com").Output()
	if err != nil {
		fmt.Println("Failed to start OpenConnect")
		c.AbortWithStatusJSON(500, gin.H{"message": "Failed to start OpenConnect"})
		return
	}
	c.JSON(200, gin.H{"message": string(out)})
}

// 停止openConnect程序
func stopOpenConnect(c *gin.Context) {
	out, err := exec.Command("bash", "-c", "sudo killall openconnect").Output()
	if err != nil {
		fmt.Println("Failed to stop OpenConnect")
		c.AbortWithStatusJSON(500, gin.H{"message": "Failed to stop OpenConnect"})
		return
	}
	c.JSON(200, gin.H{"message": string(out)})
}

// 添加用户
func addUser(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	out, err := exec.Command("bash", "-c", fmt.Sprintf("sudo useradd -m %s && echo %s | sudo passwd --stdin %s", username, password, username)).Output()
	if err != nil {
		fmt.Println("Failed to add user")
		c.AbortWithStatusJSON(500, gin.H{"message": "Failed to add user"})
		return
	}
	c.JSON(200, gin.H{"message": string(out)})
}

// 删除用户
func deleteUser(c *gin.Context) {
	username := c.Param("username")
	out, err := exec.Command("bash", "-c", fmt.Sprintf("sudo userdel -r %s", username)).Output()
	if err != nil {
		fmt.Println("Failed to delete user")
		c.AbortWithStatusJSON(500, gin.H{"message": "Failed to delete user"})
		return
	}
	c.JSON(200, gin.H{"message": string(out)})
}

// 获取用户列表
func getUserList(c *gin.Context) {
	out, err := exec.Command("bash", "-c", "sudo cat /etc/passwd").Output()
	if err != nil {
		fmt.Println("Failed to get user list")
		c.AbortWithStatusJSON(500, gin.H{"message": "Failed to get user list"})
		return
	}
	c.JSON(200, gin.H{"message": string(out)})
}