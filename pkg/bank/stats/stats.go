package stats

import (
	"github.com/coursar/bank/pkg/bank/types"
)




//Avg расчитовает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	countPayments := types.Money(len(payments))
	sumPayments := types.Money(0)
	for _, payment := range payments {
		moneyPayments := payment.Amount
		sumPayments += moneyPayments
	}
	return sumPayments / countPayments
}









// TotalInCategory находит 	сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		moneyPayments := payment.Amount
		sumPayments += moneyPayments
	}
	return sumPayments
}










//AVG расчитывает среднюю сумму платежа
/*func Avg(payments []types.Payment) types.Money  {
	
	payment := []types.Category{}

	for _, Payment := range payments {

		if Amount > 0 {
			payment = append(payment, types.Payment{

				ID: ,
				Amount: ,
				Category: "Payment" ,
	
	
	
	
	
			})
			
		}


		
	}
	
	
	
	



}


/*
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
*/