package transfer

func Total(amount int) (bonus int) {
	//bonusPercent := PercentBon(currency)
	bonus = amount + (((amount * 5) / 10) / 100)
	return bonus
}

//func PercentBon(currency string) (bonus int) {
//	if currency == "TJS" {
//		return 5
//	}
//	return 0
//}
