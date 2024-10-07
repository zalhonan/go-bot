package product

var allProducts = []Product{
	{Title: "Something tasty"},
	{Title: "Something not so tasty"},
}

type Product struct {
	Title string
}
