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

func ExamplePaymentSources_positive() {
	cards:=[]types.Card{
		{
			PAN:	"1111_xxxx_xxxx_0011",
			Balance:10_000_00,
			Active:	true,
		},
		{
			PAN:	"2222_xxxx_xxxx_0011",
			Balance:10_000_00,
			Active:	false,
		},
		{
			PAN:	"3333_xxxx_xxxx_0011",
			Balance:-10_000_00,
			Active:	false,
		},
	}
	pSource:=[]types.PaymentSource(PaymentSources(cards))
	for _, source := range pSource {
		fmt.Println(source.Number)
	}
	
	//Output:
	//1111_xxxx_xxxx_0011
}