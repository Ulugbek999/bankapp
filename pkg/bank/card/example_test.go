package card

import (
	
	"bank/pkg/bank/types"	
	
	"fmt"
)



func ExamplePaymentSources_positive() { cards:=[]types.Card{  { 
	  PAN: "1111_xxxx_xxxx_0011",   Balance:10_000_00,   Active: true,  },  {
      PAN: "2222_xxxx_xxxx_0011",   Balance:10_000_00,   Active: false,  },  { 
	  PAN: "3333_xxxx_xxxx_0011",   Balance:-10_000_00,   Active: false,  },}
	  
	  PaymentSource:=PaymentSources(cards)
	   for _, source := range PaymentSource {  fmt.Println(source.Number) }  

	   //Output: 
	   //1111_xxxx_xxxx_0011








	  }