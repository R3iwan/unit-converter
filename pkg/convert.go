package pkg

import "fmt"

var lengthFactors = map[string]float64{
	"millimeter": 1,
	"centimeter": 10,
	"meter":      1000,
	"kilometer":  1000000,
	"inch":       25.4,
	"foot":       304.8,
	"yard":       914.4,
	"mile":       1609344,
}

var weightFactors = map[string]float64{
	"milligram": 1,
	"gram":      1000,
	"kilogram":  1000000,
	"ounce":     28349.5,
	"pound":     453592.37,
}

func ConvertMeasures() {
	var choice int
	var value float64
	var fromUnit, toUnit string

	for {
		fmt.Println("Please Choose:\n1)Length\n2)Weight\n3)Temperature")
		fmt.Scan(&choice)

		fmt.Println("Enter the value to convert:")
		fmt.Scan(&value)

		fmt.Println("Enter the unit to convert from (millimeter, centimeter, meter, kilometer, inch, foot, yard, mile):")
		fmt.Scan(&fromUnit)

		fmt.Println("Enter the unit to convert to (millimeter, centimeter, meter, kilometer, inch, foot, yard, mile):")
		fmt.Scan(&toUnit)

		switch choice {
		case 1:
			result, err := convert(value, fromUnit, toUnit, lengthFactors)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("%.2f %s is equal to %.2f %s\n", value, fromUnit, result, toUnit)
		case 2:
			result, err := convert(value, fromUnit, toUnit, weightFactors)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("%.2f %s is equal to %.2f %s\n", value, fromUnit, result, toUnit)
		case 3:
			result, err := convertTemperature(value, fromUnit, toUnit)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("%.2f %s is equal to %.2f %s\n", value, fromUnit, result, toUnit)
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func convert(value float64, fromUnit, toUnit string, factors map[string]float64) (float64, error) {
	fromFactor, fromExists := factors[fromUnit]
	toFactor, toExists := factors[toUnit]

	if !fromExists || !toExists {
		return 0, fmt.Errorf("invalid unit(s)")
	}

	return (value * fromFactor) / toFactor, nil
}

func convertTemperature(value float64, fromUnit, toUnit string) (float64, error) {
	switch fromUnit {
	case "Celsius":
		switch toUnit {
		case "Fahrenheit":
			return (value * 9 / 5) + 32, nil
		case "Kelvin":
			return value + 273.15, nil
		}
	case "Fahrenheit":
		switch toUnit {
		case "Celsius":
			return (value - 32) * 5 / 9, nil
		case "Kelvin":
			return (value-32)*5/9 + 273.15, nil
		}
	case "Kelvin":
		switch toUnit {
		case "Celsius":
			return value - 273.15, nil
		case "Fahrenheit":
			return (value-273.15)*9/5 + 32, nil
		}
	}
	return 0, fmt.Errorf("invalid unit(s)")
}
