package logic

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInvaidRange(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockBudget := NewMockIBudgetRepository(ctl)
	mockBudget.EXPECT().GetAll().Return([]Bugget{})

	budget := NewBudget(mockBudget)
	startDate := time.Date(2019, 3, 1, 0,0,0,0, time.Local)
	endDate := time.Date(2019, 2, 1, 0,0,0,0, time.Local)
	actual := budget.Query(startDate, endDate)
	assert.Equal(t, float64(0), actual)
}

func TestNoBudgetData(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockBudget := NewMockIBudgetRepository(ctl)
	mockBudget.EXPECT().GetAll().Return([]Bugget{})

	budget := NewBudget(mockBudget)
	startDate := time.Date(2019, 7, 1, 0,0,0,0, time.Local)
	endDate := time.Date(2019, 7, 31, 0,0,0,0, time.Local)
	actual := budget.Query(startDate, endDate)
	assert.Equal(t, float64(0), actual)
}