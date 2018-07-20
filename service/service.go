package service

type Service interface {
	sayHello(content string) (string, error)
}

type ChinaHelloService struct{}

func (helloService *ChinaHelloService) sayHello(content string) (string, error) {
	return content, nil
}
