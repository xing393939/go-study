package pet

import (
	"github.com/win5do/golang-microservice-demo/pkg/model"
	"gorm.io/gorm"

	"github.com/win5do/go-lib/errx"

	"github.com/win5do/golang-microservice-demo/pkg/repository/db/dbcore"
)

func init() {
	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &model.OwnerPet{})
	})
}

type ownerPetDb struct {
	db *gorm.DB
}

func (s *ownerPetDb) Query(in *model.OwnerPet) ([]*model.OwnerPet, error) {
	var r []*model.OwnerPet
	err := s.db.Where(in).Find(&r).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return r, nil
}

func (s *ownerPetDb) Create(in *model.OwnerPet) (*model.OwnerPet, error) {
	err := s.db.Create(in).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return in, nil
}

func (s *ownerPetDb) Delete(in *model.OwnerPet) error {
	err := s.db.Where(in).Delete(&model.OwnerPet{}).Error
	if err != nil {
		return errx.WithStackOnce(err)
	}

	return nil
}
