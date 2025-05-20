package casql

type BaseListInput struct {
	Page uint `json:"page" v:"min:1|integer#页数最少第1页|页数必须是数值" d:"1"`
	Size uint `json:"size" v:"min:1|max:20|integer#每页至少1条内容|每页最多20条记录|每页数量必须是数值" d:"10"`
}

type BaseListOut struct {
	Items     any  `json:"items"`
	ItemCount uint `json:"item_count"`
	PageCount uint `json:"page_count"`
}
