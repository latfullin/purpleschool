package link

import (
	"kilkenny/purpleschool/pkg/db"
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

func (repo *LinkRespository) Update(hash string) (*Link, error) {
	return nil, nil
}

func (repo *LinkRespository) Delete(id string) bool {
	return false
}
