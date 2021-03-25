package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10000
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10_000, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10000

}
func ExampleAddBonus_noMoney() {
	card := types.Card{Balance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10000
}
func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: false,
		},
		{
			Balance: -10_000_00,
			Active: true,
		},
	}
	fmt.Println(Total(cards))
	// Output: 1000000
}
