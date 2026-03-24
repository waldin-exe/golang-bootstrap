package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Offset int    `json:"offset"`
	Search string `json:"search"`
	SortBy string `json:"sort_by"`
	Order  string `json:"order"`
}

// ParsePagination returns Pagination struct which contains Limit, Page, Offset, Search, SortBy, and Order.
// The offset is calculated automatically from page and limit.
func ParsePagination(c *fiber.Ctx) Pagination {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")
	search := c.Query("search")
	sortBy := c.Query("sort_by")
	order := c.Query("order")

	limit := 10 // default limit
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	page := 1 // default page
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	offset := (page - 1) * limit

	if sortBy == "" {
		sortBy = "created_at"
	}

	if order == "" {
		order = "desc"
	}

	return Pagination{
		Limit:  limit,
		Page:   page,
		Offset: offset,
		Search: search,
		SortBy: sortBy,
		Order:  order,
	}
}
