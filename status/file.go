package status

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// PathExists 判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetCurrentPath 获取当前可执行文件的所在目录
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\" char`)
	}
	return string(path[0 : i+1]), nil
}

// LogPath 日志目录
func LogPath() (string, error) {
	var logPath string
	curr, err := GetCurrentPath()
	if err != nil {
		return logPath, err
	}
	logPath = path.Join(curr, "logs")
	exists, err := PathExists(logPath)
	if err != nil {
		return logPath, err
	}
	if !exists {
		err := os.Mkdir(logPath, os.ModeDir)
		if err != nil {
			return logPath, err
		}
	}
	return logPath, nil
}

// GetLogFileName 获取日志文件
func GetLogFileName() (string, error) {
	var filename string
	logPath, err := LogPath()
	if err != nil {
		return filename, err
	}
	today := time.Now().Format("20060102")
	filename = path.Join(logPath, fmt.Sprintf("logrus_%s.log", today))
	return filename, nil
}
