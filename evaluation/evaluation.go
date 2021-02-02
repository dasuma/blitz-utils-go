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

func ByKind(kind string, value float64, rate float64) float64 {
	var res float64
	switch kind {
	case "percentage":
		if rate > 1 {
			res = (value * rate) / 100
		} else {
			res = (value * rate)
		}
	case "text":
		res = value
	case "value":
		res = value
	}
	return res
}
