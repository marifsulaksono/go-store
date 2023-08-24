package helper

type Page struct {
	Limit     int `json:"limit,omitempty"`
	Total     int `json:"total,omitempty"`
	Page      int `json:"current_page,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}
