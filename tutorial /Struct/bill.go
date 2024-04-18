package main

type Bill struct {
	name  string
	items map[string]float64
	tip   int
}

func generateNewBill(name string) Bill {

	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b Bill) format() string {
	return b.name
}
