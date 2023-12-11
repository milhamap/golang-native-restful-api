package simple

type SayHelloInterface interface {
	Hello(name string) string
}

type HelloService struct {
	SayHelloInterface
}

func NewHelloService(sayHelloInterface SayHelloInterface) *HelloService {
	return &HelloService{SayHelloInterface: sayHelloInterface}
}

type SayHello struct {
}

func NewSayHello() *SayHello {
	return &SayHello{}
}

func (s SayHello) Hello(name string) string {
	return "Hello " + name
}
