package service

import (
	"github.com/golang/mock/gomock"
	mockRepo "iris_demo/mock"
	"iris_demo/models"
	"testing"
)

func TestGetPeopleName(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mRepo := mockRepo.NewMockMovieRepository(mockCtl)
	mRepo.EXPECT().SelectMany(gomock.Any(), gomock.Any()).Return([]models.Movie{
		{
			ID:   1,
			Name: "mockname",
		},
	})

	movieService := NewMovieService(mRepo)
	all := movieService.GetAll()
	t.Logf("%v", all)
}
