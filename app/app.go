package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"webhook/config"
)

func DealRequestToExecShell(c *gin.Context) {

	token, _ := c.GetQuery("token")
	shellName, _ := c.GetQuery("shell_name")

	if shellName == "" {
		shellName = "test"
	}

	if config.Token != token {
		c.JSON(500, gin.H{"msg": "token 验证失败"})
		return
	}

	exec.Command("")
	command := `./shell/` + shellName + `.sh`
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))

}
