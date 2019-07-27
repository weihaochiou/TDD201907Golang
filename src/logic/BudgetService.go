package logic

import "time"

type Bugget struct {
	YearMonth string
	Amount int64
}

type IBudgetRepository interface {
	GetAll() []Bugget
}

type BudgetService struct {
	IBudgetService IBudgetRepository
}

func NewBudget(b IBudgetRepository) BudgetService {
	return BudgetService{IBudgetService:b}
}

func(b *BudgetService) Query(startDate time.Time, endDate time.Time) float64 {
	_budgetlist := b.IBudgetService.GetAll()
	if endDate.Before(startDate) {
		return 0
	}

	if startDate.Format("yyyyMM") == endDate.Format("yyyyMM") {
		if len(_budgetlist) == 0 {
			return 0
		}
	}


	return -1
}

