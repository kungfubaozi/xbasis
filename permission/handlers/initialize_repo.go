package permissionhandlers

type functionsConfig struct {
	Version float32              `json:"version"`
	Desc    string               `json:"desc"`
	Data    []*functionGroupData `json:"api"`
	Roles   []string             `json:"roles"`
}

type functionGroupData struct {
	GroupName string          `json:"group_name"`
	Prefix    string          `json:"prefix"`
	Functions []*functionData `json:"functions"`
}

type functionData struct {
	Api      string   `json:"api"`
	Name     string   `json:"name"`
	AuthType []int64  `json:"auth_type"`
	Roles    []string `json:"roles"`
}
