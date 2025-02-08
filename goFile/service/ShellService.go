package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"literm/goFile/types"
	"os"
)

type ShellService struct {
}

func init() {
	// 设置日志格式为JSON（可选）
	log.SetFormatter(&log.JSONFormatter{})

	// 设置日志级别（例如：debug, info, warn, error, fatal, panic）
	log.SetLevel(log.InfoLevel)

	// 输出到标准输出（默认），也可以设置为文件等其他输出
	log.SetOutput(os.Stdout)
}

func (s *ShellService) TestConnection(info types.ConnInsertInfo) *types.Response {
	ip := info.IP
	port := info.Port
	userName := ""
	password := ""
	if info.AuthType == "1" {
		userName = info.UserName
		password = info.Password
	}
	log.Println(info)
	log.Println("用户名：{}" + userName + "密码：{}" + password)
	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 注意：仅用于测试环境，在生产环境中应验证主机密钥
	}

	addr := fmt.Sprintf("%s:%s", ip, port)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return &types.Response{
			ResponseCode: "999",
			ResponseMsg:  "连接失败",
			ResponseInfo: err.Error(),
		}
	}
	defer conn.Close()

	return &types.Response{
		ResponseCode: "000",
		ResponseMsg:  "连接成功",
	}
}
