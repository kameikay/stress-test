package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/kameikay/stress-test/internal/core/ports"
)

type service struct {
	repository ports.StressTestRepository
}

func New(repository ports.StressTestRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Execute(url string, requests int, concurrency int) {
	initialTime := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go s.makeRequests(url, requests, &wg)
	}

	wg.Wait()

	fmt.Println("Total time: ", time.Since(initialTime))
	fmt.Println("Total requests: ", s.repository.GetTotalRequests())
	fmt.Println("Status 200: ", s.repository.GetStatus200())

	for k, v := range s.repository.GetOthersStatus() {
		fmt.Println("Status ", k, ": ", v)
	}
}

func (s *service) makeRequests(url string, requests int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < requests; i++ {
		resp, err := http.DefaultClient.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		s.repository.Save(resp.StatusCode)

		resp.Body.Close()
	}
}
