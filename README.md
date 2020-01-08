# Simple GoLang TDEE Calculator
Gives BMI, BMR of most popular equations, TDEE & caloric target based on percent increase/decrease of goals.

## Help
```$xslt
$ .\main --help
usage: TDEE --measurement,=MEASUREMENT, --weight=WEIGHT --height=HEIGHT --gender=GENDER --age=AGE --activity-level=ACTIVITY-LEVEL [<flags>]

A Golang Total Daily Energy Expenditure (TDEE) CLI Calculator. This program uses
the Mifflin-St Jeor Equation to calculate TDEE as it's considered to be more
accurate, see https://www.ncbi.nlm.nih.gov/pubmed/15883556.

Flags:
      --help               Show context-sensitive help (also try --help-long and
                           --help-man).
      --debug              Enable debug mode.
  -m, --measurement,=MEASUREMENT,
                           Measurement, use 'metric' or 'imperial'.
  -w, --weight=WEIGHT      Your weight.
  -h, --height=HEIGHT      Your height in meters or inches.
  -g, --gender=GENDER      Your gender.
  -a, --age=AGE            Your age.
      --activity-level=ACTIVITY-LEVEL
                           Choose one of:

                           S for Sedentary (little or no exercise, desk job)

                           LA for Lightly Active (light exercise/activity 1-3
                           days/week)

                           MA for Moderately Active (moderate exercise/activity
                           6-7 days/week)

                           VA for Very Active (2-3 hours of hard exercise every
                           day)

                           EA for Extremely Active (hard exercise 2 or more
                           times per day, or training for marathon, or
                           triathlon, etc.)
      --subtract=SUBTRACT  Subtract this percentage of TDEE and give a new
                           caloric intake. Useful for cutting.
      --add=ADD            Add this percentage of TDEE and give a new caloric
                           intake. Useful when bulking.

```

## Example
```$xslt
$ .\main -m metric --age=28 --height=1.8 --weight=85.5 --gender=male --activity-level=1.375 --add=10
Your BMI: 26.4
Original Harris-Benedict BMR: 1954 calories
Revised Harris-Benedict BMR: 1939 calories
Mifflin-St Jeor BMR: 1845 calories
Your Total Daily Energy Expenditure (TDEE): 2537 calories
With 10 % Caloric Surplus 2791
```
