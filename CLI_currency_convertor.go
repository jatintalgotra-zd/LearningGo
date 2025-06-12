package main

import (
	flag2 "flag"
	"os"

	//"errors"
	"fmt"
	//"slices"
	"strconv"
)

var conversions = map[string]float64{
	"USDINR": 85.44,
	"INRUSD": float64(1 / 85.44),
	"USDEUR": float64(1 / 0.93),
	"EURUSD": 0.93,
	"USDJPY": 156.82,
	"JPYUSD": float64(1 / 156.82),
	"INREUR": 0.010,
	"EURINR": float64(1 / 0.010),
	"EURJPY": 166.14,
	"JPYEUR": float64(1 / 166.14),
	"JPYINR": float64(1 / 1.69),
	"INRJPY": 1.69,
}

var currencies = []string{"USD", "INR", "EUR", "JPY"}

func convert(val float64, source string, target string) float64 {
	//fmt.Println(source + target)
	return val * conversions[source+target]
}

func validate(args []string) (float64, string, string, error) {
	var err error

	val, e := strconv.ParseFloat(args[1], 64)
	if e != nil {
		fmt.Println(e)
	}

	var f1, f2 bool
	for i := 0; i < len(currencies); i++ {
		if currencies[i] == args[2] {
			f1 = true
		}
		if currencies[i] == args[3] {
			f2 = true
		}
	}
	if !f1 || !f2 {
		fmt.Println("not a valid currency argument")
	}

	//if !slices.Contains(currencies, args[2]) {
	//	fmt.Println(args[2], " is not valid currency")
	//	err = errors.New("invalid currency")
	//}
	//if !slices.Contains(currencies, args[3]) {
	//	fmt.Println(args[3], " is not valid currency")
	//	err = errors.New("invalid currency")
	//}

	return val, args[2], args[3], err
}
func main() {

	// -list
	var flag = flag2.String("list", "", "Get a list of supported currencies")
	flag2.Parse()

	if *flag == "all" {
		fmt.Println("Supported currencies: ")
		for i := 0; i < len(currencies); i++ {
			fmt.Println(currencies[i])
		}
	} else {
		val, source, target, err := validate(os.Args)
		if err != nil {
			fmt.Println(err)
		} else {
			converted := convert(val, source, target)
			fmt.Printf("%f %s is equivalent to %f %s\n", val, source, converted, target)

		}
	}

	//fmt.Println(converted)

}
