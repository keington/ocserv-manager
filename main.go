package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	// 启动 HTTP 服务器
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/control", handleControl)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// 返回前端页面
	http.ServeFile(w, r, "index.html")
}

func handleControl(w http.ResponseWriter, r *http.Request) {
	// 解析 POST 请求参数
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	command := strings.TrimSpace(r.Form.Get("command"))
	parts := strings.Fields(command)
	if len(parts) < 1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	switch parts[0] {
	case "start":
		output, err := exec.Command("sudo", "systemctl", "start", "ocserv").CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to start ocserv: %v\n%s", err, output), http.StatusInternalServerError)
			return
		}
	case "stop":
		output, err := exec.Command("sudo", "systemctl", "stop", "ocserv").CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to stop ocserv: %v\n%s", err, output), http.StatusInternalServerError)
			return
		}
	case "restart":
		output, err := exec.Command("sudo", "systemctl", "restart", "ocserv").CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to restart ocserv: %v\n%s", err, output), http.StatusInternalServerError)
			return
		}
	case "add-user":
		if len(parts) < 3 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		username := parts[1]
		password := parts[2]
		output, err := exec.Command("sudo", "/usr/sbin/ocpasswd", "-c", "/etc/ocserv/ocpasswd", username).CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add user: %v\n%s", err, output), http.StatusInternalServerError)
			return
		}
		stdin, err := exec.Command("sudo", "/usr/sbin/ocpasswd", "-c", "/etc/ocserv/ocpasswd", username).StdinPipe()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to set password: %v", err), http.StatusInternalServerError)
			return
		}
		defer func(stdin io.WriteCloser) {
			err := stdin.Close()
			if err != nil {
				log.Println(err)
			}
		}(stdin)
		_, err = fmt.Fprint(stdin, password)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to set password: %v", err), http.StatusInternalServerError)
			return
		}
	case "delete-user":
		if len(parts) < 2 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		username := parts[1]
		output, err := exec.Command("sudo", "/usr/sbin/ocpasswd", "-c", "/etc/ocserv/ocpasswd", "-d", username).CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete user: %v\n%s", err, output), http.StatusInternalServerError)
			return
		}
	case "disconnect":
		if len(parts) < 2 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		sessionID := parts[1]
		output, err := exec.Command("sudo", "ocserv-control", "disconnect-session", sessionID).CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to disconnect session: %v\n%s", err, output), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
}
