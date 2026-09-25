package models

type UserTable struct {
	ID          int                 `json:"id"`
	Username    string              `json:"username"`
	Password    string              `json:"password"`
	Nickname    string              `json:"nickname"`
	Group       int                 `json:"group"`
	Enable      int                 `json:"enable"`
	TokenID     string              `json:"token_id"`
	GroupInfo   GroupTable          `json:"group_info"`
	Permissions []map[string]string `json:"permissions"`
}

type GroupTable struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Admin      int    `json:"admin"`
	Permission string `json:"permission"`
}

type PermissionTable struct {
	ID       int               `json:"id"`
	Key      string            `json:"key"`
	Name     string            `json:"name"`
	Parent   int               `json:"parent"`
	Children []PermissionTable `json:"children"`
}

type FrpRuleTable struct {
	ID        int           `json:"id"`
	Min       int           `json:"min"`
	Max       int           `json:"max"`
	Token     int           `json:"token"`
	User      int           `json:"user"`
	TokenInfo FrpTokenTable `json:"token_info"`
	UserInfo  UserTable     `json:"user_info"`
}

type FrpTokenTable struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Token    string    `json:"token"`
	User     int       `json:"user"`
	Enable   int       `json:"enable"`
	UserInfo UserTable `json:"user_info"`
}
