package model

// PageInfo 前端用于显示分页数据的请求结构体
type PageInfo struct {
	PageNum  int    `form:"page_num"`  // 当前页码
	PageSize int    `form:"page_size"` // 每一页显示多少数据项
	Key      string `form:"key"`
	Sort     string `form:"sort"` // Sort类型为string，用于在查询返回列表时指定按照什么进行排序(创建时间、主键、更新时间等等) 默认按照创建时间从新到旧排
}

// RemoveRequest 前端需要实现批量删除的请求结构体
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
