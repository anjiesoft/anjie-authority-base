package dao

type RoutesMenu struct {
	Id             int          `json:"id"`
	Name           string       `json:"name"`
	ParentId       int          `json:"parent_id"`
	Level          uint8        `json:"level"`
	Type           uint8        `json:"type"`
	Path           string       `json:"path"`
	Api            string       `json:"api"`
	ViewPath       string       `json:"view_path"`
	Identification string       `json:"identification"`
	Icon           string       `json:"icon"`
	Sort           uint8        `json:"sort"`
	IsShow         uint8        `json:"is_show"`
	Remarks        string       `json:"remarks"`
	Children       []RoutesMenu `json:"children"`
}

type RoutesAuth struct {
	Id             int    `json:"id"`
	Type           uint8  `json:"type"`
	Identification string `json:"identification"`
}

type Routes struct {
	Menu []RoutesMenu `json:"menu"`
	Auth []RoutesAuth `json:"auth"`
}
