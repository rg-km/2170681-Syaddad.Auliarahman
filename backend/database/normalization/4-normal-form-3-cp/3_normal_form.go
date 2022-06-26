package main

// Normalisasi database dalam bentuk 3NF bertujuan untuk menghilangkan seluruh atribut atau field yang tidak berhubungan dengan primary key.
// Dengan demikian tidak ada ketergantungan transitif pada setiap kandidat key
// fungsi normalisasi 3NF antara lain :
// 1. Memenuhi semua persyaratan dari bentuk normal kedua.
// 2. Menghapus kolom yang tidak tergantung pada primary key.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Rekap struct {
	NoBon     int
	Discount  int
	Total     int
	Bayar     int
	Kembalian int
	KodeKasir string
	Tanggal   string
	Waktu     string
}

type DetailRekap struct {
	NoBon      int
	KodeBarang string
	Harga      int
	Jumlah     int
}

type Barang struct {
	KodeBarang string
	NamaBarang string
	Harga      int
}

type Kasir struct {
	KodeKasir string
	NamaKasir string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
// Buatlah tabel dengan nama rekap, rekap_detail, barang, dan kasir
// Lalu insert data ke masing-masing tabel seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `CREATE TABLE rekap (
		no_bon VARCHAR(20),
		discount INTEGER,
		total INTEGER,
		bayar INTEGER,
		kembalian INTEGER,
		kode_kasir VARCHAR(20),
		tanggal VARCHAR(20),
		waktu DATETIME
	) ` // TODO: replace this
=======
	sqlStmt := `CREATE TABLE rekap ... ` // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`INSERT INTO rekap (no_bon,discount,total,bayar,kembalian,kode_kasir,tanggal,waktu)
	VALUES ("00001",0,13500, 100000, 23000, "K01", "04-05-2022", "12:00:00"), 
	("00001", 0, 36000, 100000, 23000, "K01", "04-05-2022", "12:00:00"), 
	("00001", 0, 42000, 100000, 23000, "K01", "04-05-2022", "12:00:00"), 
	("00001", 0, 77000, 100000, 23000, "K01", "04-05-2022", "12:00:00"), 
	("00002", 0, 4500, 17500, 0, "K02", "04-05-2022", "12:00:00"), 
	("00002", 0, 22000, 117500, 0, "K02", "04-05-2022", "12:00:00"), 
	("00002", 0, 117500, 117500, 0, "K02", "04-05-2022", "12:00:00")`) // TODO: replace this
=======
	_, err = db.Exec(`INSERT INTO rekap ... `) // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt = `CREATE TABLE rekap_detail (
		no_bon VARCHAR(20),
		kode_barang VARCHAR(20),
		harga INTEGER,
		jumlah INTEGER
	) ` // TODO: replace this
=======
	sqlStmt = `CREATE TABLE rekap_detail ... ` // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`INSERT INTO rekap_detail (no_bon,kode_barang,harga,jumlah) VALUES 
	("00001", "B001", 4500, 3), 
	("00001", "B002", 22500, 1), 
	("00001", "B003", 1500, 4), 
	("00001", "B004", 17500, 2), 
	("00002", "B001", 4500, 1),
	("00002", "B004", 17400, 1),
	("00002", "B005", 100000, 1)
	`) // TODO: replace this
=======
	_, err = db.Exec(`INSERT INTO rekap_detail ... `) // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt = `CREATE TABLE barang (
		kode_barang VARCHAR(20),
		nama_barang VARCHAR(50),
		harga INTEGER
	) ` // TODO: replace this
=======
	sqlStmt = `CREATE TABLE barang ... ` // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`INSERT INTO barang (kode_barang,nama_barang,harga) VALUES 
	("B001", "Disket", 4500),
	("B002", "Refil Tinta", 22500),
	("B003", "CD Blank", 1500),
	("B004", "Mouse", 17500),
	("B005", "Flash Disk", 100000)
	`) // TODO: replace this
=======
	_, err = db.Exec(`INSEERT INTO barang ... `) // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt = `CREATE TABLE kasir (
		kode_kasir VARCHAR(50),
		nama_kasir VARCHAR(50)
	);` // TODO: replace this
=======
	sqlStmt = `CREATE TABLE kasir ... ` // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	_, err = db.Exec(`INSERT INTO kasir VALUES 
		("K01", "Rosi"),
		("K02", "Dewi")
	 `) // TODO: replace this
=======
	_, err = db.Exec(`INSERT INTO kasir ... `) // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkNoBonExists:
// checkNoBonExists digunakan untuk menghitung jumlah data yang ada berdasarkan no_bon
func checkNoBonExists(noBon string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT count(no_bon) FROM rekap WHERE no_bon = ?;` // TODO: replace this
=======
	sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi countRekapDetailByNoBon:
// countRekapDetailByNoBon digunakan untuk menghitung jumlah rekap detail yang ada berdasarkan no_bon
func countRekapDetailByNoBon(noBon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT count(no_bon) FROM rekap_detail WHERE no_bon = ?;` // TODO: replace this
=======
	sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return 0, err
	}
	return countBon, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkBarangExists:
func checkBarangExists(kodeBarang string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT kode_barang FROM barang WHERE kode_barang = ?;` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeBarang)
	var latestId string
=======
	sqlStmt := `...` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeBarang)
	var latestId int
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkKasirExists:
func checkKasirExists(kodeKasir string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	sqlStmt := `SELECT kode_kasir FROM kasir WHERE kode_kasir = ?;` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeKasir)
	var latestId string
=======
	sqlStmt := `...` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeKasir)
	var latestId int
>>>>>>> 8eeb6428f7534365f78113449910a139f167243c
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}
