package lib

/*
	广点通结构体
 */
type ChanntGdt struct {
	Data struct {
		ClickID string `json:"click_id"`
		URL string `json:"url"`
		AccountID int `json:"account_id"`
		OrderTime string `json:"order_time"`
		Data []struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"data"`
	} `json:"data"`
}

/*
	百度结构体
 */
type ChanntBaiDu struct {
	PageID int `json:"page_id"`
	Data []struct {
		Name string `json:"name"`
		Value string `json:"value"`
	} `json:"data"`
	SiteID int `json:"site_id"`
}