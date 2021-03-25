package card

import (
	"bank/pkg/bank/types"
)

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	bonus := (int(card.MinBalance) * percent / 100 * daysInMonth) / daysInYear
	if !(*card).Active {
		return
	}
	if card.Balance < 0 {
		return
	}
	if bonus > 5_000_00 {
		bonus = 0
	}
	card.Balance += types.Money(bonus)
}
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > 50_000_00 {
		return
	}
	if amount < 0 {

		return
	}
	card.Balance += amount
}
func Total(cards []types.Card) int {
	sum :=0
	for _, card := range cards {
		if (card.Balance<0&&card.Active==false){
			return sum
		}
		sum += int(card.Balance)
	}
	return sum
}