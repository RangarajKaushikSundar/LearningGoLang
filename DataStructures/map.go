package main

import "fmt"

func main() {

	// Creating a map using Composite Literals
	ageMap := map[string]int{
		"Kaushik": 29,
		"Fruity":  28,
		"Rajesh":  29,
	}
	fmt.Println("Age Map:", ageMap)
	fmt.Println("Kaushik's age", ageMap["Kaushik"])
	fmt.Println("Loosa's age", ageMap["Loosa"])

	fmt.Println()
	// Accessing the map
	checkifNameExistsInMap(ageMap, "Loosa")
	checkifNameExistsInMap(ageMap, "Fruity")

	fmt.Println()
	// Adding to the Map
	fmt.Println("Adding Loosa to the ageMap ...")
	ageMap["Loosa"] = 24
	checkifNameExistsInMap(ageMap, "Loosa")

	fmt.Println()
	// Accessing the map using for...range
	fmt.Println("Printing all the elements of the ageMap")
	for key, value := range ageMap {
		fmt.Println(key, value)
	}

	fmt.Println()
	// Delete from the map
	fmt.Println("Deleting Loosa from the ageMap")
	delete(ageMap, "Loosa")
	checkifNameExistsInMap(ageMap, "Loosa")

}

func checkifNameExistsInMap(ageMap map[string]int, name string) {
	// Easy Lookup
	fmt.Print("Quick lookup for ", name, ". ")
	if value, ok := ageMap[name]; ok {
		fmt.Println(name, "exists in the map. Age:", value)
	} else {
		fmt.Println(name, "doesn't exist in the map.")
	}
}
