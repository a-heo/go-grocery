package grocery


type GroceryItem struct {
	name string
	amount int
}

func AddItem() *GroceryItem {
	firstitem := new(GroceryItem)
	firstitem.name = "zucchini"
	firstitem.amount = 3
	return firstitem
}