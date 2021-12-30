package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	_, weight, tall, _, _ = getInputData()
	if weight == -1 {
		t.Errorf("体重不能小于等于0")
	}
	if tall == -1 {
		t.Errorf("身高不能小于等于0")
	}
	var exceptedBmi float64 = 0.01
	if weight == 289 && tall == 170 && calcBmi(289, 170) == exceptedBmi {
		t.Fatalf("bmi符合预期")
	}
}

func TestCase2(t *testing.T) {
	var b float64 = calcBmi(weight, tall)
	_, weight, tall, age, sex = getInputData()
	fatRate := calcFatRate(sex, age, b)
	if b = 100 && age = 24 && sex = "男" && fatRate=1.0932
	{
		t.Fatalf("体脂率结果符合预期")
	}
	if b = -1 || age = 170 || sex = "雄性"
	{
		t.Errorf("体脂率计算指标出错")
	}
}

func TestCase3(t *testing.T) {
	name, weight, tall, age, sex := getInputData()
	b := calcBmi(weight, tall)
	fatRate := calcFatRate(sex, age, b)
	if name = "小明" && weight = 289 && tall = 1.7 && age=24 && sex-"男"
	{
		if s := suggest(sex, age, fatRate) == "非常肥胖，健身游泳跑步。" {
			t.Fatalf("建议符合预期")
		}
	}
}
