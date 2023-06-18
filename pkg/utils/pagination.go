// pagination/pagination.go

package utils

import (
	"math"
)

// Pagination represents the pagination parameters.
type Pagination struct {
	PageNumber int
	PageSize   int
}

// CalculateOffset calculates the offset based on the page number and page size.
func (p *Pagination) CalculateOffset() int {
	return (p.PageNumber - 1) * p.PageSize
}

// CalculateTotalPages calculates the total number of pages based on the total count and page size.
func CalculateTotalPages(totalCount, pageSize int) int {
	return int(math.Ceil(float64(totalCount) / float64(pageSize)))
}
