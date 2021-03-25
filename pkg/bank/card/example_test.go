package card

import (
	"bank/pkg/bank/types"	
	"fmt"
)



func ExampleWithdraw_positive() {
	cards := types.Card{Balance:  2_000_000, Active: true}
	 Withdraw(&cards, 1_000_000)
	fmt.Println(cards.Balance)


	// Output:
	// 1000000
}

func ExampleDeposit_possitive()  {
	card := types.Card{Balance: 1_000_000, Active: true}
	Deposit(&card, 1_000_000)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
	
}

func ExampleDeposit_inactive()  {
	card := types.Card{Balance: 1000000, Active: false}
	Deposit(&card, 1000)
	fmt.Println(card.Balance)

	// Output:
	// 1000000
	
}


func ExampleDeposit_negative()  {

	card := types.Card{Balance: -1000, Active: true}
	Deposit(&card, 2000)
	fmt.Println(card.Balance)

	// Output:
	// 1000

	
}




// func ExampleWithdraw_noMoney()  {
// 	card := types.Card{Balance: 0, Active: true}

// 	 Withdraw(&card,1_000_000)
// 	fmt.Println(card.Balance)

// 	// Output:
// 	// 0
	
//}

// func ExampleWithdraw_inactive()  {
// 	result := Withdraw(types.Card{Balance: 2_000_000, Active: false}, 1_000_000)
// 	fmt.Println(result.Balance)

// 	// Output:
// 	// 2000000
	
// }


// func ExampleWithdraw_limit()  {
// 	result := Withdraw(types.Card{Balance: 10_000_000, Active: true}, 2_000_000)
// 	fmt.Println(result.Balance)

// 	// Output:
// 	// 10000000

	
// }


func ExampleDepositeLimit()  {
	card := types.Card{Balance: 1_000_000, Active: true}
	Deposit(&card, 5_000_001)
	fmt.Println(card.Balance)

	// Output: 
	// 1000000
}





