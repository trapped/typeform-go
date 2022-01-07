package typeform

type WorkspaceListParams struct {
	Page     int `url:"page,omitempty"`
	PageSize int `url:"page_size,omitempty"`
}

type WorkspaceList struct {
	Items      []*Workspace `json:"items,omitempty"`
	PageCount  int          `json:"page_count"`
	TotalItems int          `json:"total_items"`
}
