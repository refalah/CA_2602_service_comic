package models

type Comic struct {
	// gorm.Model
	ID           int            `gorm:"primaryKey;autoIncrement"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Publisher    string         `json:"publisher"`
	Cover        string         `json:"cover"`
	CreativeTeam []CreativeTeam `gorm:"foreignKey:ComicId`
}

func (c *Comic) CreateNewComic() *Comic {
	db.Create(&c)
	return c
}

func GetComicById(Id int64) (*Comic, error) {
	comic := &Comic{}

	// Add Preload to fetch CreativeTeam with Creator and Role
	err := db.Preload("CreativeTeam.Creator").
		Preload("CreativeTeam.Role").
		Where("ID = ?", Id).
		First(&comic).Error

	if err != nil {
		return nil, err
	}

	return comic, nil
}

func GetAllComics() []Comic {
	Comics := []Comic{}
	db.Find(&Comics)
	return Comics
}

func (c *Comic) UpdateComicById(Id int64) *Comic {
	db.Where("ID = ?", Id).Updates(Comic{Title: c.Title, Description: c.Description, Publisher: c.Publisher, Cover: c.Cover})
	return c
}

func (c *Comic) DeleteComicById(Id int64) *Comic {
	db.Where("ID = ?", Id).Delete(c)
	return c
}