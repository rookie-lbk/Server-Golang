package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	//创建一个文件服务器，会去static目录下找index.html
	fileServer := http.FileServer(http.Dir("./static"))
	// 将 "/" 路径映射到文件服务器
	http.Handle("/", fileServer)
	// 定义 "/privacy-policy" 路径的处理器（隐私协议）
	http.HandleFunc("/privacy-policy", PrivacyPolicyHandler)
	// 定义 "/terms-of-service" 路径的处理器（使用条款）
	http.HandleFunc("/terms-of-service", TermsOfServiceHandler)

	fmt.Printf("Starting server at port 8080\n")
	// 启动HTTP服务器并监听端口 8080，如果出现错误，则打印错误信息并退出
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}

// PrivacyPolicyHandler 隐私协议处理器
func PrivacyPolicyHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/privacy-policy.html")
}

// TermsOfServiceHandler 使用条款处理器
func TermsOfServiceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/terms-of-service.html")
}
