package services

type Service struct{}

func (s *Service) GetMessage() string {
	return "Hello, from go Microservices"
}
