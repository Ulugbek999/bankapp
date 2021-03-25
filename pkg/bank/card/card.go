package card

import (
	"bank/pkg/bank/types"
	
)


func PaymentSources(cards []types.Card) []types.PaymentSource {

		payment := []types.PaymentSource{}
		
			
			for _, card := range cards {
			if !card.Active == true{
				continue
			}
				
			if card.Balance <= 0 {
				continue
			}
			
			payment = append(payment, types.PaymentSource{
				Type: "card",
				Number: string(card.PAN),
				Balance: card.Balance,
			  })

			  
		}

		
		return payment

	}

		
	


















	



