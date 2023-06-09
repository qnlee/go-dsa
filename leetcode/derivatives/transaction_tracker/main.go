package main

import "fmt"

// NOTE: Solution contained in single main file to resemble CoderPad environment

/*
==================================================================================================
Create a class that tracks customer transactions. Who has transacted with who can change over time
Example:
	Have Jim and Pam ever transacted together? -> No
	Jim pays Pam
	Have Jim and Pam ever transacted together? -> Yes
	Pam pays Michael
	Have Michael and Jim ever transacted together? -> No

Given list of tuples representing transactions
e.g. [("alice", "john"), • ("jenny", "zoe"')]

We want to be able to answer the question "has customer A transacted with customer B?"
e.g. Has Alice ever transacted with John? -> Yes
	 Has John ever transacted with Jenny? -> No
==================================================================================================
*/

type TransactionTracker struct {
	customerRecords map[string]*customer
}

// NewTransactionTracker will retain stored transactions
func NewTransactionTracker() *TransactionTracker {
	return &TransactionTracker{
		customerRecords: make(map[string]*customer),
	}
}

func (tt *TransactionTracker) AddTransaction(from, to string) {
	// assumes names can't be empty and customers can't transact with themselves
	if from == "" || to == "" || from == to {
		return
	}
	// if any customer ('from', 'to', or both) not found in customerRecords, add as new customer
	if !tt.knownCustomer(from) {
		tt.customerRecords[from] = newCustomer(from)
	}
	if !tt.knownCustomer(to) {
		tt.customerRecords[to] = newCustomer(to)
	}
	// set transactedWith relationships for both customers (to one another)
	tt.customerRecords[from].transactedWith[to] = true
	tt.customerRecords[to].transactedWith[from] = true
}

func (tt *TransactionTracker) HaveTransacted(customer1, customer2 string) bool {
	// if 'a' has ever transacted with 'b' (or vice versa), both 'a' and 'b' should exist in customerRecords
	// assumes that the only way to add new customers to customerRecords is via AddTransaction
	if !tt.knownCustomer(customer1) || !tt.knownCustomer(customer2) {
		return false
	}
	// from above assumption, customerRecords[customer1].transactedWith[customer2] should always ==
	// customerRecords[customer2].transactedWith[customer1] so it doesn't matter which we use to return here
	return tt.customerRecords[customer1].transactedWith[customer2]
}

// helper function to improve readability & keep duplicated logic minimal
func (tt *TransactionTracker) knownCustomer(name string) bool {
	_, exists := tt.customerRecords[name] // exists will be true if name already exists (as a key) in customerRecords
	return exists                         // false otherwise
}

// private struct not meant to be used outside of this package
type customer struct {
	name           string          // assumes customer names are unique; otherwise need other unique identifier (eg. ID)
	transactedWith map[string]bool // map keys serve as 'set' of customer names that this customer has transacted with
}

func newCustomer(name string) *customer {
	return &customer{name: name, transactedWith: make(map[string]bool)}
}

/*
NOTE: Lots of duplicate print logic here mostly due to time constraints, CoderPad-type IDE limitations, and potential
benefits of being more explicit in cases where interviewer may not be familiar with language used.
The below "tests" could be refactored into table-driven unit tests in a separate test file, allowing this logic to
be more maintainable and "DRY".
*/
func main() {
	tt := NewTransactionTracker()
	fmt.Printf("START\n>>> len(tt.customerRecords): %d\n", len(tt.customerRecords))

	fmt.Println("\nCASE: Switching order of arguments to AddTransaction does not affect result")
	fmt.Printf(">>> Lea, Pam: %v\n", tt.HaveTransacted("Lea", "Pam")) // false
	fmt.Println(">>> Adding transaction... from=Pam to=Lea")
	tt.AddTransaction("Pam", "Lea")
	fmt.Printf(">>> Lea, Pam: %v\n", tt.HaveTransacted("Lea", "Pam")) // true
	fmt.Printf("RES: len(tt.customerRecords): %d\n", len(tt.customerRecords))
	for n, c := range tt.customerRecords {
		fmt.Printf("     %s - %v\n", n, c.transactedWith)
	}

	fmt.Println("\nCASE: Switching order of arguments to HaveTransacted does not affect result")
	fmt.Printf(">>> Jim, Pam: %v\n", tt.HaveTransacted("Jim", "Pam")) // false
	fmt.Println(">>> Adding transaction... from=Jim to=Pam")
	tt.AddTransaction("Jim", "Pam")
	fmt.Printf(">>> Jim, Pam: %v\n", tt.HaveTransacted("Jim", "Pam")) // true
	fmt.Printf(">>> Pam, Jim: %v\n", tt.HaveTransacted("Pam", "Jim")) // true
	fmt.Printf("RES: len(tt.customerRecords): %d\n", len(tt.customerRecords))
	for n, c := range tt.customerRecords {
		fmt.Printf("     %s - %v\n", n, c.transactedWith)
	}

	fmt.Println("\nCASE: Adding new transaction with ex ")
	fmt.Printf(">>> Pam, Michael: %v\n", tt.HaveTransacted("Pam", "Michael")) // false
	fmt.Printf(">>> Jim, Michael: %v\n", tt.HaveTransacted("Jim", "Michael")) // false
	fmt.Println(">>> Adding transaction... from=Pam to=Michael")
	fmt.Printf(">>> Jim, Michael: %v\n", tt.HaveTransacted("Jim", "Michael")) // false
	fmt.Printf(">>> Michael, Jim: %v\n", tt.HaveTransacted("Michael", "Jim")) // false
	fmt.Printf(">>> Michael, Pam: %v\n", tt.HaveTransacted("Michael", "Pam")) // true
	fmt.Printf("RES: len(tt.customerRecords): %d\n", len(tt.customerRecords))
	tt.AddTransaction("Pam", "Michael")
	for n, c := range tt.customerRecords {
		fmt.Printf("     %s - %v\n", n, c.transactedWith)
	}

	fmt.Println("\nCASE: Customer transactions with themselves (assuming unique names) will not be added")
	fmt.Printf(">>> Jim, Jim: %v\n", tt.HaveTransacted("Jim", "Jim")) // false
	fmt.Println(">>> Adding transaction... from=Jim to=Jim")
	tt.AddTransaction("Jim", "Jim")
	fmt.Printf(">>> Jim, Jim: %v\n", tt.HaveTransacted("Jim", "Jim")) // false
	fmt.Printf("RES: len(tt.customerRecords): %d\n", len(tt.customerRecords))
	for n, c := range tt.customerRecords {
		fmt.Printf("     %s - %v\n", n, c.transactedWith)
	}

	fmt.Println("\nCASE: Transactions involving customer with empty '' name will not be added")
	fmt.Printf(">>> Jim, '': %v\n", tt.HaveTransacted("Jim", "")) // false
	fmt.Println(">>> Adding transaction... from=Jim to=")
	tt.AddTransaction("Jim", "")
	fmt.Printf(">>> Jim, '': %v\n", tt.HaveTransacted("Jim", "")) // false
	fmt.Printf("RES: len(tt.customerRecords): %d\n", len(tt.customerRecords))
	for n, c := range tt.customerRecords {
		fmt.Printf("     %s - %v\n", n, c.transactedWith)
	}
}
