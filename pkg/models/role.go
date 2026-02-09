package models

type Role struct {
	// gorm.Model
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}

func GetAllRoles() []Role {
	Roles := []Role{}
	db.Find(&Roles)
	return Roles
}

func GetRoleById(Id int64) *Role  {
	role := Role{}
	db.Where("ID = ?", Id).Find(&role)
	return &role
}

func (r *Role) CreateNewRole() *Role {
	db.Create(&r)
	return r
}

func (r *Role) UpdateRole(Id int64) *Role {
	db.Where("ID = ?", Id).Update("name", r.Name)
	return  r
}

func DeleteRoleById(Id int64) *Role  {
	role := Role{}
	db.Where("ID = ?", Id).Delete(&role)
	return &role
}