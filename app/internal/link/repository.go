package link

import (
	"kilkenny/purpleschool/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRespository struct {
	Databases *db.Db
}

func NewLinkRepository(databses *db.Db) *LinkRespository {
	return &LinkRespository{
		Databases: databses,
	}
}

func (repo *LinkRespository) Create(link *Link) (*Link, error) {
	res := repo.Databases.DB.Create(link)

	if res.Error != nil {
		return nil, res.Error
	}

	return link, nil
}

func (repo *LinkRespository) Get(hash string) (*Link, error) {
	var link Link
	result := repo.Databases.DB.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRespository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.Databases.DB.First(&link, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRespository) Update(link *Link) (*Link, error) {
	result := repo.Databases.Clauses(clause.Returning{}).Updates(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *LinkRespository) Delete(id uint) error {

	_, err := repo.GetById(id)

	if err == nil {
		return err
	}

	result := repo.Databases.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
