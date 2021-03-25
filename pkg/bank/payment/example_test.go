package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExampleMax_possitive()  {
	paymentns := []types.Payment{
		{
			ID: 1,
			Amount: 100,
		},
		{
			ID: 2,
			Amount: 200,
		},
		{
			ID: 3,
			Amount: 500,
		},
		{
			ID: 4,
			Amount: 300,
		},

	}


	payment := Max(paymentns)
	fmt.Println(payment.ID)

	// Output:
	// 3
	
}