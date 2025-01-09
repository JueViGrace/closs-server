package types

type CustomerQueries struct{}

type ManagerCustomerQueries struct {
	CustomerQueries
	Manager string `query:"manager"`
}

type SalesmanCustomerQueries struct {
	CustomerQueries
	Salesman string `query:"salesman"`
}

type DocumentQueries struct {
	WithLines bool `query:"withLines"`
}

type ManagerDocumentQueries struct {
	DocumentQueries
	Manager string `query:"manager"`
}

type SalesmanDocumentQueries struct {
	DocumentQueries
	Salesman string `query:"salesman"`
}

type CustomerDocumentQueries struct {
	DocumentQueries
	Customer string `query:"customer"`
}
