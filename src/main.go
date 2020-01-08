package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"math"
	"os"
)

type Person struct {
	age                    uint8
	gender                 string
	weightKG               float64
	weightLB               float64
	heightM                float64
	heightIn               float64
	bmi                    float64
	originalHarrisBenedict float64
	revisedHarrisBenedict  float64
	mifflinStJeor          float64
	tdee                   float64
}

var (
	app = kingpin.New("TDEE", "A Golang Total Daily Energy Expenditure (TDEE) CLI Calculator."+
		"This program uses the Mifflin-St Jeor Equation to calculate TDEE as it's considered to be more accurate, see  "+
		"https://www.ncbi.nlm.nih.gov/pubmed/15883556.")
	debug = app.Flag("debug", "Enable debug mode.").Bool()

	// TDEE
	unitFlag     = app.Flag("unit,", "unit, use 'metric' or 'imperial'.").Short('u').Required().String()
	weightFlag   = app.Flag("weight", "Your weight.").Short('w').Required().Float64()
	heightFlag   = app.Flag("height", "Your height in meters or inches. ").Short('h').Required().Float64()
	genderFlag   = app.Flag("gender", "Your gender.").Short('g').Required().String()
	ageFlag      = app.Flag("age", "Your age.").Short('a').Required().Uint8()
	activityFlag = app.Flag("activity-level", "Choose one of:"+
		"\n\n1.2 for Sedentary (little or no exercise, desk job)"+
		"\n\n1.375 for Lightly Active (light exercise/activity 1-3 days/week)"+
		"\n\n1.55 for Moderately Active (moderate exercise/activity 6-7 days/week)"+
		"\n\n1.725 for Very Active (2-3 hours of hard exercise every day)"+
		"\n\n1.9 for Extremely Active (hard exercise 2 or more times per day, or training for marathon, or triathlon, etc.)").Required().Float64()
	subtractFlag = app.Flag("subtract", "Subtract this percentage of TDEE and give a new caloric intake. Useful for cutting.").Float64()
	addFlag      = app.Flag("add", "Add this percentage of TDEE and give a new caloric intake. Useful when bulking.").Float64()
)

func (p Person) calcBMI(wKG float64, hM float64) (bmi float64) {
	bmi = wKG / math.Pow(hM, 2)
	bmi = math.Round(bmi*10) / 10
	return
}

/*
	THIS FUNCTION CALCULATES THE ORIGINAL HARRIS-BENEDICT BMR EQUATION
		MEN - BMR = 66.4730 + (13.7516 x weight in kg) + (5.0033 x height in cm) – (6.7550 x age in years)
		WOMEN - BMR = 655.0955 + (9.5634 x weight in kg) + (1.8496 x height in cm) – (4.6756 x age in years)
*/
func (p Person) calcOHarrisBenedict(wKG float64, hM float64, gender string, age uint8) (oHBBMR float64) {
	if gender == "male" {
		oHBBMR = 66.4730 + (13.7515 * wKG) + (5.0033 * (hM * 100)) - (6.755 * float64(age))
		oHBBMR = math.Round((oHBBMR * 10) / 10)
	}
	if gender == "female" {
		oHBBMR = 655.0955 + (9.5634 * wKG) + (1.8496 * (hM * 100)) - (4.6756 * float64(age))
		oHBBMR = math.Round((oHBBMR * 10) / 10)
	}
	return
}

/*
	THIS FUNCTION CALCULATES THE REVISED HARRIS-BENEDICT BMR EQUATION
		MEN - BMR = 88.362 + (13.397 x weight in kg) + (4.799 x height in cm) - (5.677 x age in years)
		WOMEN - BMR = 447.593 + (9.247 x weight in kg) + (3.098 x height in cm) - (4.330 x age in years)
*/
func (p Person) calcRHarrisBenedict(wKG float64, hM float64, gender string, age uint8) (rHBBMR float64) {
	if gender == "male" {
		rHBBMR = math.Round(((88.362 + (13.397 * wKG) + (4.799 * (hM * 100)) - (5.677 * float64(age))) * 10) / 10)
	}
	if gender == "female" {
		rHBBMR = math.Round(((447.593 + (9.247 * wKG) + (3.098 * (hM * 100)) - (4.330 * float64(age))) * 10) / 10)
	}
	return
}

/*
	THIS FUNCTION CALCULATES THE MIFFLIN-ST JEOR BMR EQUATION
		MEN - BMR (metric) = (10 × weight in kg) + (6.25 × height in cm) - (5 × age in years) + 5
		WOMEN - BMR (metric) = (10 × weight in kg) + (6.25 × height in cm) - (5 × age in years) - 161
*/
func (p Person) calcMifflinStJeor(wKG float64, hM float64, gender string, age uint8) (mjBMR float64) {
	if gender == "male" {
		mjBMR = math.Round((((10 * wKG) + (6.25 * (hM * 100)) - (5 * float64(age)) + 5) * 10) / 10)
	}
	if gender == "female" {
		mjBMR = math.Round((((10 * wKG) + (6.25 * (hM * 100)) - (5 * float64(age)) - 161) * 10) / 10)
	}
	return
}

func (p Person) calcTDEE(a float64) (tdee float64) {
	tdee = math.Round(((p.mifflinStJeor * a) * 10) / 10)
	return
}

func (p Person) subPFromTDEE(tdee float64, percent float64) (calTar float64) {
	pCals := ((percent * tdee) / 100)
	calTar = math.Round(((tdee - pCals) * 10) / 10)
	return
}

func (p Person) addPFromTDEE(tdee float64, percent float64) (calTar float64) {
	pCals := ((percent * tdee) / 100)
	calTar = math.Round(((tdee + pCals) * 10) / 10)
	return
}

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	p := Person{}
	p.age = *ageFlag
	p.gender = *genderFlag
	if *unitFlag != "metric" {
		p.weightKG = *weightFlag / 2.205
		p.heightM = *heightFlag / 39.37
		p.bmi = p.calcBMI(p.weightKG, p.heightM)
		fmt.Println("Your BMI:", p.bmi)

		// PRINT ORIGINAL HARRIS BENEDICT EQUATION
		p.originalHarrisBenedict = p.calcOHarrisBenedict(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Original Harris-Benedict BMR:", p.originalHarrisBenedict, "cals")

		// PRINT REVISED HARRIS BENEDICT EQUATION
		p.revisedHarrisBenedict = p.calcRHarrisBenedict(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Revised Harris-Benedict BMR:", p.revisedHarrisBenedict, "cals")

		// PRINT MIFFLIN-ST. JEOR EQUATION
		p.mifflinStJeor = p.calcMifflinStJeor(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Mifflin-St Jeor BMR:", p.mifflinStJeor, "cals")

		// PRINT TDEE
		p.tdee = p.calcTDEE(*activityFlag)
		fmt.Println("Your Total Daily Energy Expenditure (TDEE):", p.tdee, "cals")

		// SUBTRACT PERCENT FROM TDEE IF SUBTRACT FLAG IS PASSED
		if *subtractFlag != 0 && (*addFlag == 0) {
			calTar := p.subPFromTDEE(p.tdee, *subtractFlag)
			fmt.Println("With", *subtractFlag, "% Caloric Reduction:", calTar, "cals")
		}
		if (*addFlag != 0) && (*subtractFlag == 0) {
			calTar := p.addPFromTDEE(p.tdee, *addFlag)
			fmt.Println("With", *addFlag, "% Caloric Surplus", calTar, "cals")
		}
	}
	if *unitFlag == "metric" {
		p.weightKG = *weightFlag
		p.heightM = *heightFlag

		// PRINT BMI
		p.bmi = p.calcBMI(p.weightKG, p.heightM)
		fmt.Println("Your BMI:", p.bmi)

		// PRINT ORIGINAL HARRIS BENEDICT EQUATION
		p.originalHarrisBenedict = p.calcOHarrisBenedict(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Original Harris-Benedict BMR:", p.originalHarrisBenedict, "cals")

		// PRINT REVISED HARRIS BENEDICT EQUATION
		p.revisedHarrisBenedict = p.calcRHarrisBenedict(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Revised Harris-Benedict BMR:", p.revisedHarrisBenedict, "cals")

		// PRINT MIFFLIN-ST. JEOR EQUATION
		p.mifflinStJeor = p.calcMifflinStJeor(p.weightKG, p.heightM, p.gender, p.age)
		fmt.Println("Mifflin-St Jeor BMR:", p.mifflinStJeor, "cals")

		// PRINT TDEE
		p.tdee = p.calcTDEE(*activityFlag)
		fmt.Println("Your Total Daily Energy Expenditure (TDEE):", p.tdee, "cals")

		// SUBTRACT PERCENT FROM TDEE IF SUBTRACT FLAG IS PASSED
		if *subtractFlag != 0 && (*addFlag == 0) {
			calTar := p.subPFromTDEE(p.tdee, *subtractFlag)
			fmt.Println("With", *subtractFlag, "% Caloric Reduction:", calTar, "cals")
		}
		if (*addFlag != 0) && (*subtractFlag == 0) {
			calTar := p.addPFromTDEE(p.tdee, *addFlag)
			fmt.Println("With", *addFlag, "% Caloric Surplus", calTar, "cals")
		}
	}
}
