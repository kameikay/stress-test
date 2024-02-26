package repository

type RequestsMemory struct {
	Requests ReportRequest
}

type ReportRequest map[int]int

func NewRequestsMemory() *RequestsMemory {
	return &RequestsMemory{
		Requests: make(ReportRequest),
	}
}

func (s *RequestsMemory) Save(statusCode int) {
	_, ok := s.Requests[statusCode]
	if ok {
		s.Requests[statusCode]++
	} else {
		s.Requests[statusCode] = 1
	}
}

func (s *RequestsMemory) GetTotalRequests() int {
	total := 0

	for _, v := range s.Requests {
		total += v
	}

	return total
}

func (s *RequestsMemory) GetStatus200() int {
	return s.Requests[200]
}

func (s *RequestsMemory) GetOthersStatus() map[int]int {
	others := make(map[int]int)

	for k, v := range s.Requests {
		if k != 200 {
			others[k] = v
		}
	}

	return others
}

func (s *RequestsMemory) GetErrorRequests() int {
	return s.Requests[0]
}
