package res

// PageParam 分页查询条件
type PageParam struct {
	PageSign string `query:"pageSign"`            // 请求参数, total | list | both
	PageNo   uint   `query:"pageNo,default=1"`    // 当前页
	PageSize uint   `query:"pageSize,default=20"` // 页大小 binding:"max=100"
	Total    uint   `query:"total"`               // 上次统计的数据条数
}

// PageResult 分页数据
type PageResult struct {
	PageSign string      `json:"pageSign,default=both,omitempty"` // 请求参数, total | list | both
	PageNo   int         `json:"pageNo,omitempty"`                // 页索引
	PageSize int         `json:"pageSize,omitempty"`              // 页条数
	Total    int         `json:"total,omitempty"`                 // 总条数
	List     interface{} `json:"list,omitempty"`                  // 数据
}