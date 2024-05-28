package service

import (
	"fmt"

	"github.com/ivofulco/goexpert-stress-test/internal/domain"
)

type reportService struct{}

func NewReportService() domain.ReportGenerator {
	return &reportService{}
}

func (s *reportService) GenerateReport(result domain.TestResult) {
	fmt.Printf("Tempo total gasto: %v\n", result.Duration)
	fmt.Printf("Total de requests: %d\n", result.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", result.StatusCounts[200])
	fmt.Printf("Requests com status 300: %d\n", result.StatusCounts[300])
	fmt.Printf("Requests com status 404: %d\n", result.StatusCounts[404])
	fmt.Printf("Requests com status 429: %d\n", result.StatusCounts[429])
	fmt.Printf("Requests com status 500: %d\n", result.StatusCounts[500])

	for status, count := range result.StatusCounts {
		if status != 200 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
}
