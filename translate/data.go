package main

type MapData struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Category    string `json:"category"`
			Subcat      string `json:"subcat"`
			Description string `json:"description"`
			Map         string `json:"map"`
			Color       string `json:"color"`
			Icon        string `json:"icon"`
			Hash        string `json:"hash"`
			Position    string `json:"position"`
			Elevation   int    `json:"elevation"`
			Completed   bool   `json:"completed"`
		} `json:"properties"`
		Geometry struct {
			Type        string `json:"type"`
			Coordinates []int  `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}
