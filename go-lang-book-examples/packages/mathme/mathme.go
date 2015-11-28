package mathme

func Average(xs []float64) float64 {
	total := float64(0)
	itens := len(xs)

	for _, value := range xs {
		total += value
	}

	return total / float64(itens)
}