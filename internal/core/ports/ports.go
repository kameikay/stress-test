package ports

type StressTestRepository interface {
	Save(statusCode int)
	GetTotalRequests() int
	GetStatus200() int
	GetOthersStatus() map[int]int
	GetErrorRequests() int
}
