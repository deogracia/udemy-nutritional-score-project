package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	help := flag.Bool("help", false, "Show help")
	energy := flag.Float64("energy", -42, "Energy (Kcal)")
	sugars := flag.Float64("sugars", -42, "Sugars (g)")
	sfa := flag.Float64("sfa", -42, "Saturated Fatty Acids (g)")
	sodium := flag.Float64("sodium", -42, "Sodium (g)")
	fiber := flag.Float64("fiber", -42, "Fiber (g)")
	protein := flag.Float64("protein", -42, "Protein (g)")
	fruitspercent := flag.Float64("fruitspercent", -42, "Fruits Percent")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("vim-go")
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(*energy),
		Sugars:              SugarGram(*sugars),
		SaturatedFattyAcids: SaturatedFattyAcids(*sfa),
		Sodium:              SodiumMilligram(*sodium),
		Fruits:              FruitsPercent(*fruitspercent),
		Fiber:               FiberGram(*fiber),
		Protein:             ProteinGram(*protein),
	}, Food)

	fmt.Printf("Nutritional Score is %v\n", ns.Value)
	fmt.Printf("Nutriscore is %v\n", ns.GetNutriScore())
}
