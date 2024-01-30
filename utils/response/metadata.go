package response

type Page struct {
	Limit     int `json:"limit,omitempty"`
	Total     int `json:"total,omitempty"`
	Page      int `json:"current_page,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}

type UserInfo struct {
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Role     string `json:"role,omitempty"`
}
