package model

type Tag struct {
	*Model        // 将Model结构体引入
	Name   string `json:"name"`
	State  uint8  `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}
