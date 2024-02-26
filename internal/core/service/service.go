package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/kameikay/stress-test/internal/core/ports"
	"github.com/kameikay/stress-test/pkg/utils"
)

type service struct {
	repository ports.StressTestRepository
	mu         sync.Mutex
}

func New(repository ports.StressTestRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Execute(url string, requests int, concurrency int) {
	initialTime := time.Now()

	statusChannel := make(chan int)

	for i := 0; i < concurrency; i++ {
		go s.request(url, requests, statusChannel)
	}

	for i := 0; i < requests; i++ {
		status := <-statusChannel
		s.mu.Lock()
		s.repository.Save(status)
		s.mu.Unlock()
	}

	s.PrintResponse(initialTime)
}

func (s *service) request(url string, times int, statusChannel chan int) {
	for i := 0; i < times; i++ {
		resp, err := utils.MakeRequest(url)
		if err != nil {
			fmt.Println("Error: ", err)
			statusChannel <- 0
			continue
		}
		defer resp.Body.Close()
		statusChannel <- resp.StatusCode
	}
}

func (s *service) PrintResponse(initialTime time.Time) {
	fmt.Println("Total time: ", time.Since(initialTime))
	fmt.Println("Total requests: ", s.repository.GetTotalRequests())
	fmt.Println("Status 200: ", s.repository.GetStatus200())
	fmt.Println("Error requests: ", s.repository.GetErrorRequests())

	for k, v := range s.repository.GetOthersStatus() {
		fmt.Println("Status ", k, ": ", v)
	}
}
