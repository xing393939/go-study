package repo

import (
	"context"
	"gorm.io/gorm"
	"my-clean-rchitecture/models"
	"time"
)

type mysqlArticleRepository struct {
	DB *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(DB *gorm.DB) IArticleRepo {
	return &mysqlArticleRepository{DB}
}

func (m *mysqlArticleRepository) Fetch(ctx context.Context, createdDate time.Time,
	num int) (res []models.Article, err error) {

	// 使用的内存数据库，先初始化表
	tmp := models.Article{ID: 1}
	_ = m.DB.AutoMigrate(&tmp)
	m.DB.Create(&tmp)

	err = m.DB.WithContext(ctx).Model(&models.Article{}).
		Select("id").
		Limit(num).Find(&res).Error
	return
}
