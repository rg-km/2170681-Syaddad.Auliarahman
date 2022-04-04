package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 1 * BaseSalary

// TODO: answer here
type Employee interface {
	GetBonus() int
}

type Manager struct {
	BaseSalary int
}

func (mn Manager) GetBonus() int {
	return mn.BaseSalary * 3
}

type SeniorEngineer struct {
	BaseSalary int
}

func (se SeniorEngineer) GetBonus() int {
	return se.BaseSalary * 2
}

type JuniorEngineer struct {
	BaseSalary int
}

func (je JuniorEngineer) GetBonus() int {
	return je.BaseSalary
}

func TotalEmployeeBonus(employees []Employee) int {
	// Hitunglah total bonus yang dikeluarkan dari list of Employee
	// TODO: answer here
	var eachBonus int
	for _, em := range employees {
		eachBonus += em.GetBonus()
	}
	return eachBonus
}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	// TODO: answer here
	manager := Manager{BaseSalary: 200}
	senior := SeniorEngineer{BaseSalary: 150}
	junior := JuniorEngineer{BaseSalary: 100}
	fmt.Println(TotalEmployeeBonus([]Employee{manager, senior, junior}))
	fmt.Println(manager.GetBonus())
	fmt.Println(senior.GetBonus())
	fmt.Println(junior.GetBonus())

}
