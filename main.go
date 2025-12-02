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
	// 定义 "/hello" 路径的处理器
	http.HandleFunc("/hello", HelloHandler)
	// 定义 "/form" 路径的处理器
	http.HandleFunc("/form", FormHandler)

	fmt.Printf("Starting server at port 8080\n")
	// 启动HTTP服务器并监听端口 8080，如果出现错误，则打印错误信息并退出
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//判断请求路径是否正确
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	//判断请求方式是否正确
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	//判断请求路径是否正确
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	//判断请求方式是否正确
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	//解析表单数据是否正确
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name is:%s\n", name)
	fmt.Fprintf(w, "address is: %s", address)
}
