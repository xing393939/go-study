// file: services/movie_service.go

package service

import (
	"iris_demo/models"
	"iris_demo/repo"
)

// `MovieService` 会处理一些 `movie` 数据模型层的 CRUID 操作
// 这取决于 `movie` 存储库 的一些行为.
//这里将数据源和高级组件进行解耦
// 所以，我们可以在不做任何修改的情况下，轻松的切换使用不同的储库类型
// 这个是一个通用的接口
//因为我们可能需要在不的地方修改和尝试不同的逻辑
type MovieService interface {
	GetAll() []models.Movie
	GetByID(id int64) (models.Movie, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndGenreByID(id int64, poster string, genre string) (models.Movie, error)
}

// NewMovieService 返回默认的 movie 服务层.
func NewMovieService(repo repo.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

type movieService struct {
	repo repo.MovieRepository
}

// GetAll 返回所有的 movies.
func (s *movieService) GetAll() []models.Movie {
	return s.repo.SelectMany(func(_ models.Movie) bool {
		return true
	}, -1)
}

// GetByID 根据 id 返回一个 movie .
func (s *movieService) GetByID(id int64) (models.Movie, bool) {
	return s.repo.Select(func(m models.Movie) bool {
		return m.ID == id
	})
}

// UpdatePosterAndGenreByID 更新 一个 movie 的 poster 和 genre 字段.
func (s *movieService) UpdatePosterAndGenreByID(id int64, poster string, genre string) (models.Movie, error) {
	// update the movie and return it.
	return s.repo.InsertOrUpdate(models.Movie{
		ID:     id,
		Poster: poster,
		Genre:  genre,
	})
}

// DeleteByID 根据 id 删除一个 movie
//
// Returns true if deleted otherwise false.
func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m models.Movie) bool {
		return m.ID == id
	}, 1)
}