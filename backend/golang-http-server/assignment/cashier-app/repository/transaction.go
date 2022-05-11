package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
	salesRepository    SalesRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository, salesRepository SalesRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository, salesRepository}
}

func (u *TransactionRepository) Pay(cartItems []CartItem, amount int) (int, error) {
	total, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}

	result := amount - total
	if err := u.cartItemRepository.ResetCartItems(); err != nil {
		return 0, err
	}
	return result, nil
}
