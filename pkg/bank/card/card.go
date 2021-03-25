package card

import "bank/pkg/bank/types"



func IssueCard(currency types.Currency, color string, name string) types.Card   {
	card := types.Card{
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}
	

	return card
	
}



func Withdraw(card *types.Card, amount types.Money)  {
	
	

	limit := 2_000_000

	if card.Active == false{
		return 
	}
	

	if card.Balance < amount{
		return 
	}
	if amount >= types.Money(limit){
		return 
	}
	card.Balance -= amount



	}


	func Deposit(card *types.Card, amount types.Money)  {
		
		if amount < 0{
			return 
		}

		if !card.Active{
			return
		}

		if amount > 5_000_000{
			return
		}




		card.Balance += amount
		
	}
	


	func Total(cards []types.Card) types.Money  {
		var sum types.Money

	

		for _, card := range cards{
			
			sum +=  card.Balance

		}

		
		return sum
		
	}


