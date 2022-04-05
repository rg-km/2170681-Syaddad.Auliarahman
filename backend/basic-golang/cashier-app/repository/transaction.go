package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {
	// TODO: replace this

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
