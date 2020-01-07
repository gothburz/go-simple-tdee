package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"math"
	"os"
)

type Person struct {
	age uint8
	gender string
	weightKG float64
	weightLB float64
	heightM float64
	heightIn float64
	bmi float64
	originalHarrisBenedict float64
	revisedHarrisBenedict float64
	mifflinStJeor float64
	//bodyFat float32
	//bmi uint16
	//tdee uint16
}

var (
	app 	 = kingpin.New("TDEE", "A Golang Total Daily Energy Expenditure (TDEE) CLI Calculator.")
	debug    = app.Flag("debug", "Enable debug mode.").Bool()

	// TDEE
	mFlag = app.Flag("measurement,", "Measurement, use 'metric' or 'imperial'.").Short('m').Required().String()
	weightFlag = app.Flag("weight", "Your weight.").Short('w').Required().Float64()
	heightFlag = app.Flag("height", "Your height in meters or inches. ").Short('h').Required().Float64()
	genderFlag = app.Flag("Gender", "Your gender.").Short('g').Required().String()
	ageFlag = app.Flag("Age", "Your age.").Short('a').Required().Uint8()

)

func (p Person) calcBMI(wKG float64, hM float64) (bmi float64){
	bmi = wKG / math.Pow(hM, 2)
	bmi = math.Round(bmi * 10) / 10
	return
}

// THIS FUNCTION CALCULATES THE ORIGINAL HARRIS-BENEDICT BMR EQUATION
func (p Person) calcOHarrisBenedict(wKG float64, hM float64, gender string, age uint8) (oHBBMR float64) {
	if gender == "male" {
		oHBBMR = 66.4730 + (13.7515 * wKG) + (5.0033 * (hM * 100)) - (6.755 * float64(age))
		oHBBMR = math.Round((oHBBMR * 100) / 100)
	}
	if gender == "female" {
		oHBBMR = 655.0955 + (9.5634 * wKG) + (1.8496 * (hM * 100)) - (4.6756 * float64(age))
		oHBBMR = math.Round((oHBBMR * 100) / 100)
	}
	return
}

// THIS FUNCTION CALCULATES THE REVISED HARRIS-BENEDICT BMR EQUATION
func (p Person) calcRHarrisBenedict(wKG float64, hM float64, gender string, age uint8)  (rHBBMR float64){
	if gender == "male" {
		rHBBMR = 88.362 + (13.397 * wKG) + (4.799 * (hM * 100)) - (5.677 * float64(age))
		rHBBMR = math.Round((rHBBMR * 100) / 100)
	}
	if gender == "female" {
		rHBBMR = 447.593 + (9.247 * wKG) + (3.098 * (hM * 100)) - (4.330 * float64(age))
		rHBBMR = math.Round((rHBBMR * 100) / 100)
	}
	return
}

// THIS FUNCTION CALCULATES THE MIFFLIN-ST JEOR BMR EQUATION
func (p Person) calcMifflinJeor(wKG float64, hM float64, gender string, age uint8) (mjBMR float64){
	if gender == "male" {
		mjBMR = (10 * wKG) + (6.25 * (hM * 100)) - (5 * float64(age)) + 5
		mjBMR = math.Round((mjBMR * 100) / 100)
	}
	if gender == "female" {
		mjBMR = (10 * wKG) + (6.25 * (hM * 100)) - (5 * float64(age)) - 161
		mjBMR = math.Round((mjBMR * 100) / 100)
	}
	return
}



func main(){
	kingpin.MustParse(app.Parse(os.Args[1:]))
	p := Person{}
	p.age = *ageFlag
	p.gender = *genderFlag
	if *mFlag != "metric" {
		p.weightKG = *weightFlag / 2.205
		p.heightM = *heightFlag / 39.37
		p.bmi = p.calcBMI(p.weightKG, p.heightM)
		fmt.Println("Your BMI:", p.bmi)
	}
	if *mFlag == "metric" {
		p.weightKG = *weightFlag
		p.heightM = *heightFlag
		p.bmi = p.calcBMI(p.weightKG, p.heightM)
		fmt.Println("Your BMI:", p.bmi)
		p.originalHarrisBenedict = p.calcOHarrisBenedict(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Original Harris-Benedict BMR:", p.originalHarrisBenedict, "calories")
		p.revisedHarrisBenedict = p.calcRHarrisBenedict(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Revised Harris-Benedict BMR:", p.revisedHarrisBenedict, "calories")
		p.mifflinStJeor = p.calcMifflinJeor(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Mifflin-St Jeor BMR:", p.mifflinStJeor, "calories.")
	}
}
