package validator

// LogLoginItemsValidator 查看列表
type LogLoginItemsValidator struct {
	Page     int      `form:"page" json:"page"`
	Size     int      `form:"size" json:"size"`
	Name     string   `form:"name" json:"name"`
	AdminId  int      `form:"admin_id" json:"admin_id"`
	Ip       string   `form:"ip" json:"ip"`
	Username string   `form:"username" json:"username"`
	Status   uint8    `form:"status" json:"status"`
	Time     []string `form:"time" json:"time"`
}

// GetMessage 查看列表 - 提示消息
func (p LogLoginItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{}
}

// LogActionValidator 创建
type LogActionValidator struct {
	ProjectId      int    `form:"project_id" json:"project_id" binding:"required"`
	Content        string `form:"content" json:"content" binding:"required"`
	Type           uint8  `form:"type" json:"type" binding:"required"`
	ModuleName     string `form:"module_name" json:"module_name" binding:"required"`
	Ip             string `form:"ip" json:"ip" binding:"required"`
	BrowserInfo    string `form:"browser_info" json:"browser_info" binding:"required"`
	BrowserName    string `form:"browser_name" json:"browser_name" binding:"required"`
	BrowserVersion string `form:"browser_version" json:"browser_version" binding:"required"`
	Reason         string `form:"reason" json:"reason"`
}

// GetMessage 创建 - 提示消息
func (p LogActionValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"ProjectId.required":      "项目不能为空",
		"Content.required":        "内容不能为空",
		"Type.required":           "类型不能为空",
		"TypeName.required":       "类型名称不能为空",
		"Ip.required":             "IP不能为空",
		"BrowserInfo.required":    "浏览器信息不能为空",
		"BrowserName.required":    "浏览器名称不能为空",
		"BrowserVersion.required": "浏览器版本不能为空",
	}
}

// LogActionItemsValidator 查看列表
type LogActionItemsValidator struct {
	Page       int      `form:"page" json:"page"`
	Size       int      `form:"size" json:"size"`
	Type       uint8    `form:"type" json:"type"`
	Ip         string   `form:"ip" json:"ip"`
	AdminId    int      `form:"admin_id" json:"admin_id"`
	AdminName  string   `form:"admin_name" json:"admin_name"`
	ProjectId  int      `form:"project_id" json:"project_id"`
	ModuleName string   `form:"module_name" json:"module_name"`
	Time       []string `form:"time" json:"time"`
}

// GetMessage 查看列表 - 提示消息
func (p LogActionItemsValidator) GetMessage() ValidatorMessages {
	return ValidatorMessages{}
}
