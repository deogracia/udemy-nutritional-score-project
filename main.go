package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(100),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumMilligram(500),
		Fruits:              FruitsPercent(60),
		Fiber:               FiberGram(4),
		Protein:             ProteinGram(2),
	}, Food)

	fmt.Printf("Nutritional Score is %v\n", ns.Value)
	fmt.Printf("Nutriscore is %v\n", ns.GetNutriScore())
}
