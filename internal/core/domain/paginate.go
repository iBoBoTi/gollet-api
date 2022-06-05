package domain

type Paginate struct {
	Limit      int         `form:"limit,default=5"`
	Page       int         `form:"page,default=1"`
	Offset     int         `form:"offset,default=0"`
	TotalRows  int         `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func NewPaginate(limit, page int) *Paginate {
	paginator := &Paginate{
		Limit: limit,
		Page:  page,
	}
	return paginator
}
