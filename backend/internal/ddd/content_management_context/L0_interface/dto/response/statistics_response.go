package response

type BlogOverviewStatsResp struct {
	TotalArticles   int `json:"totalArticles"`
	TotalLikes      int `json:"totalLikes"`
	TotalViews      int `json:"totalViews"`
	TotalCategories int `json:"totalCategories"`
	TotalTags       int `json:"totalTags"`
	TotalComments   int `json:"totalComments"`
}

type ViewCountTrendItemResp struct {
	Date  string `json:"date"`
	Views int    `json:"views"`
}

type ViewCountTrendResp struct {
	Trend []ViewCountTrendItemResp `json:"trend"`
	Days  int                      `json:"days"`
}
