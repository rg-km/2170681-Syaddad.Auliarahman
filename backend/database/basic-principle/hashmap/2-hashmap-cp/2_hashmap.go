package main

import "fmt"

type PrimaryKey int

type InvoiceRow struct {
	ID               PrimaryKey
	SubscriptionCode string
	CustomerName     string
	Address          string
	Phone            string
}

type InvoiceDB struct {
	m map[PrimaryKey]InvoiceRow
}

func NewInvoice() *InvoiceDB {
	return &InvoiceDB{
		m: make(map[PrimaryKey]InvoiceRow),
	}
}

func (db *InvoiceDB) Insert(code string, name string, address string, phone string) {
	db.m[PrimaryKey(len(db.m)+1)] = InvoiceRow{
		ID:               PrimaryKey(len(db.m)) + 1,
		SubscriptionCode: code,
		CustomerName:     name,
		Address:          address,
		Phone:            phone,
	}
}

func (db *InvoiceDB) Where(id PrimaryKey) *InvoiceRow {
	// for i := 0; i < len(db.m); i++ {
	// 	if db.m[i].ID == id {
	// 		return &db.m[i]
	// 	}
	// }
	// return nil
	// return &db.m[id]
	row := (*db).m[id]
	return &InvoiceRow{
		ID:               id,
		SubscriptionCode: row.SubscriptionCode,
		CustomerName:     row.CustomerName,
		Address:          row.Address,
		Phone:            row.Phone,
	} // TODO: replace this
}

func (db *InvoiceDB) Update(id PrimaryKey, code string, name string, address string, phone string) (*InvoiceRow, error) {
	// (*db).m[id] = InvoiceRow{
	// 	ID:               id,
	// 	SubscriptionCode: code,
	// 	CustomerName:     name,
	// 	Address:          address,
	// 	Phone:            phone,
	// }
	// return &(*db).m[id], nil
	row := db.Where(id)
	if row == nil {
		return nil, fmt.Errorf("not found")
	}

	return &InvoiceRow{
		SubscriptionCode: code,
		CustomerName:     name,
		Address:          address,
		Phone:            phone,
	}, nil
	// return nil, nil // TODO: replace this
}
