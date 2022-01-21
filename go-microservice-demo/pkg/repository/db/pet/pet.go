package pet

import (
	"github.com/win5do/golang-microservice-demo/pkg/model"
	"gorm.io/gorm"

	"github.com/win5do/go-lib/errx"

	"github.com/win5do/golang-microservice-demo/pkg/repository/db/dbcore"
)

func init() {
	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &model.Pet{})
	})
}

type petDb struct {
	db *gorm.DB
}

func (s *petDb) List(query *model.Pet, offset, limit int) ([]*model.Pet, error) {
	var r []*model.Pet

	db := dbcore.WithOffsetLimit(s.db, offset, limit)

	err := db.Where(query).Find(&r).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return r, nil
}

func (s *petDb) Get(id string) (*model.Pet, error) {
	var r model.Pet
	err := s.db.Where("id = ?", id).First(&r).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return &r, nil
}

func (s *petDb) Create(in *model.Pet) (*model.Pet, error) {
	err := s.db.Create(in).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return in, nil
}

func (s *petDb) Update(in *model.Pet) (*model.Pet, error) {
	err := s.db.Updates(in).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return in, nil
}

func (s *petDb) Delete(in *model.Pet) error {
	err := s.db.Where(in).Delete(&model.Pet{}).Error
	if err != nil {
		return errx.WithStackOnce(err)
	}

	return nil
}
