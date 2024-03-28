package response

type Page struct {
	Limit     int `json:"limit,omitempty"`
	Total     int `json:"total,omitempty"`
	Page      int `json:"current_page,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}

type LoginInfo struct {
	Username     string `json:"username,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
