package go_bmi

import "fmt"

func calcBmi(weight, height float64) (bmi float64, err error) {
	if weight <= 0 {
		err = fmt.Errorf("体重不能小于等于0")
		return
	}

	if height <= 0 {
		err = fmt.Errorf("身高不能小于等于0")
		return
	}
	bmi = weight / (height * height)
	return
}
