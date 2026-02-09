package models

type Creator struct {
	// gorm.Model
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	CreativeTeam []CreativeTeam `gorm:"foreignKey:CreatorId"`
}

func (c *Creator) CreateNewCreator() *Creator {
	db.Create(&c)
	return c
}

func GetCreatorById(Id int64) *Creator  {
	Creators := &Creator{}
	db.Where("ID = ?", Id).Find(&Creators)
	return Creators
}

func GetAllCreators() []Creator  {
	Creators := []Creator{}
	db.Find(&Creators)
	return Creators
}

func (c *Creator) UpdateCreatorById(Id int64) *Creator{
	db.Where("ID = ?", Id).Updates(Creator{Name: c.Name, Description: c.Description, Image: c.Image})
	return c
}

func (c *Creator) DeleteCreatorById(Id int64) *Creator{
	db.Where("ID = ?", Id).Delete(c)
	return c
}