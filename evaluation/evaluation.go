package evaluation

func Start(tax float64, min float64, kind string) float64 {
	res := tax
	if tax < min {
		switch kind {
		case "value":
			res = min
		default:
			res = (tax * min) / 100
		}
	}
	return res
}
