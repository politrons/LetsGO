package commands

type CreateOrder struct {
}

type AddProduct struct {
	orderId            string
	productId          string
	productDescription string
	price              float64
}

type RemoveProduct struct {
	orderId   string
	productId string
}
