package seeds

import "fmt"

func RunSeeds() {
	defer fmt.Println("Seeds completed")
	SeedCourts()
}
