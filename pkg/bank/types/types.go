package types

//Money представляет собой минимальных денежных единицах например(дирамы, копейки и.т)
type Money int64

//Currency представляет  код валюты
type Currency string

//код валют
const (
	TJS Currency="TJS"
	RUB Currency="RUB"
	USD Currency="USD"
)

//PAN представляет номер карты
type PAN string

//Card представляет информацию о платёжном карте
type Card struct{
	ID int
	PAN PAN
	Balance Money
	Color string
	Name string 
	Active bool
}