package services

/*

	This is a service, it is a layer between the controller and anything else. A service is good for code that shouldn't be in the controller and should/could be reusable.

*/

type HelloWorldService struct {
}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

func (s *HelloWorldService) HelloVenus() string {
	return "Hello Venus!"
}
