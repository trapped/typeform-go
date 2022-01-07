package typeform

type ThemeListParams struct {
	Page                 int  `url:"page,omitempty"`
	PageSize             int  `url:"page_size,omitempty"`
	IncludePremiumThemes bool `url:"premium_themes,omitempty"`
}

type ThemeList struct {
	Items      []*Theme `json:"items,omitempty"`
	PageCount  int      `json:"page_count"`
	TotalItems int      `json:"total_items"`
}
