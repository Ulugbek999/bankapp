package payment

import (
	"bank/pkg/bank/types"
)



func Max(payments []types.Payment) types.Payment  {
	payment := types.Payment{}


	for _, value := range payments{
		if value.Amount > payment.Amount{
			payment = value
		}
	}

	return payment
}