package models

type CreativeTeam struct {
	// gorm.Model
	ID        int `gorm:"primaryKey;autoIncrement"`
	ComicId   int `json:"comicId"`
	CreatorId int `json:"creatorId"`
	RoleId    int `json:"roleId"`

	Comic   Comic   `gorm:"foreignKey:ComicId"`
	Creator Creator `gorm:"foreignKey:CreatorId"`
	Role    Role    `gorm:"foreignKey:RoleId"`
}

func (c *CreativeTeam) CreateNewCreativeTeam() *CreativeTeam {
	db.Create(&c)
	return c
}

func GetCreativeTeamByComicId(comicId int64) ([]CreativeTeam, error) {
	var creativeTeam []CreativeTeam

	err := db.Preload("Creator").
		Preload("Role").Preload("Comic").
		Where("comic_id = ?", comicId).
		Find(&creativeTeam).Error

	if err != nil {
		return nil, err
	}

	return creativeTeam, nil
}

func GetAllCreativeTeams() []CreativeTeam {
	CreativeTeams := []CreativeTeam{}
	db.Find(&CreativeTeams)
	return CreativeTeams
}