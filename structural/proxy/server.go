package main

import "fmt"

/* https://golangbyexample.com/proxy-design-pattern-in-golang/
代理设计模式是一种结构设计模式。这种模式建议为控制和智能访问主要对象提供额外的间接层。
在此模式中，创建了一个新的代理类，该类实现了与主对象相同的接口。这允许您在主对象的实际逻辑之前执行一些行为。让我们用一个例子来更多地理解它
1.借记卡是您银行账户的代理。它遵循与银行账户相同的接口，并且更容易使用。
2.像Nginx这样的Web服务器可以作为应用程序服务器的代理。它提供
	对应用程序服务器的受控访问-例如，它可以进行速率限制
	附加行为-例如，它可以进行一些缓存
在下面的UML图中
	主题：它代表了代理和真实主题应该遵循的接口
	代理：代理类嵌入realSubject，并在完成处理后将请求传递给realSubject
	实语：这是持有主要业务逻辑的类，并保持在代理后面
	客户端：可以与代理和realSubject进行交互，因为它们都遵循相同的界面。
*/

type server interface {
	handleRequest(string, string) (int, string)
}

type application struct{}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}

	return 404, "Not Ok"
}

type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginx() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}

	return n.application.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] += 1

	return true
}

func main() {
	n := newNginx()
	appStatusUrl := "/app/status"
	createUseUrl := "/create/user"
	httpCode, body := n.handleRequest(appStatusUrl, "GET")
	fmt.Printf("url: %v, httpCode: %v, body: %v\n", appStatusUrl, httpCode, body)
	httpCode, body = n.handleRequest(appStatusUrl, "GET")
	fmt.Printf("url: %v, httpCode: %v, body: %v\n", appStatusUrl, httpCode, body)
	httpCode, body = n.handleRequest(appStatusUrl, "GET")
	fmt.Printf("url: %v, httpCode: %v, body: %v\n", appStatusUrl, httpCode, body)
	httpCode, body = n.handleRequest(createUseUrl, "POST")
	fmt.Printf("url: %v, httpCode: %v, body: %v\n", createUseUrl, httpCode, body)
	httpCode, body = n.handleRequest(createUseUrl, "GET")
	fmt.Printf("url: %v, httpCode: %v, body: %v\n", createUseUrl, httpCode, body)
}
