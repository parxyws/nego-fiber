package domain

type Metadata struct {
	Pagination Pagination       `json:"pagination"`
	SortBy     string           `json:"sort_by,omitempty"`
	SortOrder  string           `json:"sort_order,omitempty"`
	Filters    map[string]any   `json:"filters,omitempty"`
	Links      *PaginationLinks `json:"links,omitempty"`
	Cursor     *CursorInfo      `json:"cursor,omitempty"`
}

type Pagination struct {
	PageNumber int   `json:"page_number"`
	PageLimit  int   `json:"page_limit"`
	TotalCount int64 `json:"total_count"`
	TotalSize  int64 `json:"total_size"`
}

type PaginationLinks struct {
	Self string `json:"self,omitempty"`
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

type CursorInfo struct {
	NextCursor string `json:"next_cursor,omitempty"`
	PrevCursor string `json:"prev_cursor,omitempty"`
}

type Cursor map[string]any
