package protocol

const AREA_TYPE_COUNTRY = 1
const AREA_TYPE_PROVINCE = 2
const AREA_TYPE_CITY = 3
const AREA_TYPE_DISTRICT = 4

type AreaCreationParam struct {
	ParentID uint   `json:"parent_id"` // 父级编号
	Code     string `json:"code"`      // 区域编码
	Name     string `json:"name"`      // 区域名称
	Type     string `json:"type"`      // 区域类型（1：国家；2：省份、直辖市；3：地市；4：区县）
}
