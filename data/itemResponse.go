package data

type ItemResponse struct {
	Items            []MediaItem `json:"Items"`
	TotalRecordCount int         `json:"TotalRecordCount"`
	StartIndex       int         `json:"StartIndex"`
}
