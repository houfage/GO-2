package main

import (
	"fmt"
	"github.com/armstrongli/go-bmi"
)

func main() {
	//录入
	name, weight, tall, age, sex := getInputData()

	//计算bmi
	b, err := gobmi.BMI(weight, tall)
	if err != nil {
		panic(err)
	}

	//计算体脂率
	fatRate, err := calcFatRate(sex, age, b)
	if err != nil {
		panic(err)
	}
	//var bmi float64
	//var fatrate float64
	//bmi = weight / (tall * tall)
	//建议
	s, err := suggest(sex, age, fatRate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("你好%v，你的建议是：%v", name, s)
}

func getInputData() (name string, weight, tall float64, age int, sex string) {

	fmt.Print("姓名：")
	fmt.Scanln(&name)

	fmt.Print("体重：")
	fmt.Scanln(&weight)

	fmt.Print("身高：")
	fmt.Scanln(&tall)

	fmt.Print("年龄：")
	fmt.Scanln(&age)

	fmt.Print("性别：")
	fmt.Scanln(&sex)
	return
}

func calcFatRate(sex string, age int, bmi float64) (fatRate float64, err error) {
	var sexweight int
	if sex == "男" {
		sexweight = 1
	} else {
		sexweight = 0
	}
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexweight)) / 100

	fmt.Println("体脂率是：", fatRate)
	return
}

func suggest(sex string, age int, fatRate float64) (suggest string, err error) {
	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				suggest = "偏瘦。要多吃多锻炼，增强体质。"
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				suggest = "标准。很好，要保持。"
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				suggest = "偏胖。吃完饭多散步，消化消化。"
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				suggest = "肥胖。少吃点，多运动。"
			} else {
				suggest = "非常肥胖。健身游泳跑步。"
			}

		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.1 {
				suggest = "偏瘦。要多吃多锻炼，增强体质。"
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				suggest = "标准。很好，要保持。"
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				suggest = "偏胖。吃完饭多散步，消化消化。"
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				suggest = "肥胖。少吃点，多运动。"
			} else {
				suggest = "非常肥胖。健身游泳跑步。"
			}

		} else if age >= 60 {
			if fatRate <= 0.1 {
				suggest = "偏瘦。要多吃多锻炼，增强体质。"
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				suggest = "标准。很好，要保持。"
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				suggest = "偏胖。吃完饭多散步，消化消化。"
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				suggest = "肥胖。少吃点，多运动。"
			} else {
				suggest = "非常肥胖。健身游泳跑步。"
			}

		} else {
			suggest = "无法评判未成年人。"
		}

	} else {
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				suggest = "偏瘦。要多吃多锻炼，增强体质。"
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				suggest = "标准。很好，要保持。"
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				suggest = "偏胖。吃完饭多散步，消化消化。"
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				suggest = "肥胖。少吃点，多运动。"
			} else {
				suggest = "非常肥胖。健身游泳跑步。"
			}

		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.1 {
				suggest = "偏瘦。要多吃多锻炼，增强体质。"
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				suggest = "标准。很好，要保持。"
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				suggest = "偏胖。吃完饭多散步，消化消化。"
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				suggest = "肥胖。少吃点，多运动。"
			} else {
				suggest = "非常肥胖。健身游泳跑步。"
			}

		} else if age >= 60 {
			if fatRate <= 0.1 {
				suggest = "偏瘦。要多吃多锻炼，增强体质。"
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				suggest = "标准。很好，要保持。"
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				suggest = "偏胖。吃完饭多散步，消化消化。"
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				suggest = "肥胖。少吃点，多运动。"
			} else {
				suggest = "非常肥胖。健身游泳跑步。"
			}

		} else {
			suggest = "无法评判未成年人。"
		}

	}
	return
}
