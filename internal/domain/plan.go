package domain

// CityPlan представляет градостроительный план
type CityPlan struct {
	ID          int64
	Name        string
	Description string
	Area        float64
	Population  int
}
