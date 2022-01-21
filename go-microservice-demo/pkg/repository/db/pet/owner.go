package pet

import (
	"github.com/win5do/golang-microservice-demo/pkg/model"
	"gorm.io/gorm"

	"github.com/win5do/go-lib/errx"

	"github.com/win5do/golang-microservice-demo/pkg/repository/db/dbcore"
)

func init() {
	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &model.Owner{})
	})
}

type ownerDb struct {
	db *gorm.DB
}

func (s *ownerDb) List(query *model.Owner, offset, limit int) ([]*model.Owner, error) {
	var r []*model.Owner

	db := dbcore.WithOffsetLimit(s.db, offset, limit)

	err := db.Where(query).Find(&r).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return r, nil
}

func (s *ownerDb) Get(id string) (*model.Owner, error) {
	var r model.Owner
	err := s.db.Where("id = ?", id).First(&r).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return &r, nil
}

func (s *ownerDb) Create(in *model.Owner) (*model.Owner, error) {
	err := s.db.Create(in).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return in, nil
}

func (s *ownerDb) Update(in *model.Owner) (*model.Owner, error) {
	err := s.db.Updates(in).Error
	if err != nil {
		return nil, errx.WithStackOnce(err)
	}

	return in, nil
}

func (s *ownerDb) Delete(in *model.Owner) error {
	err := s.db.Where(in).Delete(&model.Owner{}).Error
	if err != nil {
		return errx.WithStackOnce(err)
	}

	return nil
}
