package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	// 0 // TODO: replace this
	var total int
	for _, subordinate := range vp.Subordinate {
		total += subordinate.TotalDivisonSalary()
	}
	total += vp.GetSalary()
	return total
}
