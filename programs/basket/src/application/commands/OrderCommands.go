package commands

type CreateOrder struct {
}

type AddProduct struct {
	ProductId          string
	ProductDescription string
	Price              float64
}

type RemoveProduct struct {
	ProductId string
}
