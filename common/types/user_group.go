package types

type UserGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UserGroup) TableName() string {
	return "user_group"
}
