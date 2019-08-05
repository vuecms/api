package model

//go:generate goqueryset -in sys.go -out sys_query.go

// gen:qs
type Area struct {
	//gorm.Model
	Model
	Parent    *Area  `json:"parent"`
	ParentID  uint   `json:"parent_id"`  // 父级编号
	ParentIds string `json:"parent_ids"` // 所有父级编号
	Children  []Area `json:"children" gorm:"foreignkey:ParentID"`
	Code      string `json:"code"` // 区域编码
	Name      string `json:"name"` // 区域名称
	Type      string `json:"type"` // 区域类型（1：国家；2：省份、直辖市；3：地市；4：区县）
}

//gen:qs
type Organization struct {
	Model
	Parent      *Organization
	ParentId    uint   // 父级编号
	ParentIds   string // 所有父级编号
	Area        Area   // 归属区域
	Code        string // 机构编码
	Name        string // 机构名称
	Type        int8   // 机构类型（1：公司；2：部门；3：小组）
	Grade       int8   // 机构等级（1：一级；2：二级；3：三级；4：四级）
	Address     string // 联系地址
	ZipCode     string // 邮政编码
	Master      string // 负责人
	PhoneNumber string // 电话
	Fax         string // 传真
	Email       string // 邮箱
}

//gen:qs
type Department struct {
	Model
	Name           string
	Organization   Organization `gorm:"foreignkey:OrganizationID"`
	OrganizationID int
}
