package types

// Money представляет собой денежную валюту в минимальных единицах
type Money int64

// Category представляет собой катенорию в которой был совершён платёж
type Category string

// Payment представляет информацию о платежах
type Payment struct{
	ID int
	Amount Money
	Category Category
}