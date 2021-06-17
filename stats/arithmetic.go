package stats

func Mean(population []float64)(mean float64){
	populationLength := float64(len(population))
	if populationLength == 0{
		return 0
	}

	var populationSum float64
	for _, sampleValue := range population{
		populationSum += sampleValue
	}
	return populationSum/populationLength
}
