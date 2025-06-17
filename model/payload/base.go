// Package payload ...
package payload

// BaseResponse ...
type BaseResponse struct {
	Success bool          `json:"success"`
	Data    interface{}   `json:"data,omitempty"`
	Meta    *MetaResponse `json:"meta,omitempty"`
	Message *string       `json:"message,omitempty" default:"berhasil"`
	Error   interface{}   `json:"error,omitempty"`
}

// MetaResponse ...
type MetaResponse struct {
	CurrentPage  uint64 `json:"current_page" default:"1"`
	NextPage     *int64 `json:"next_page" default:"2"`
	PreviousPage *int64 `json:"previous_page,omitempty" default:"0"`
	Limit        uint64 `json:"limit" default:"100"`
	TotalData    uint64 `json:"total_data" default:"100"`
	LastPage     uint64 `json:"last_page,omitempty" default:"1"`
}

// OrderingFilter ...
type OrderingFilter struct {
	Order     string `query:"order" json:"order" validate:"" example:"created_at"`
	Direction string `query:"direction" json:"direction" validate:"oneof=ASC DESC" example:"ASC" enums:"ASC,DESC"`
}

// PaginationFilter ...
type PaginationFilter struct {
	Page    uint64 `query:"page" json:"page" validate:"gte=1" example:"1" default:"1"`
	PerPage uint64 `query:"per_page" json:"per_page" validate:"gte=1,lte=100" example:"50" default:"50"`
}

// BaseQueryFilter ...
type BaseQueryFilter struct {
	PaginationFilter
	OrderingFilter
}
