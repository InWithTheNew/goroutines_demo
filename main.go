package main

import "fmt"

func main() {

	serialQuery, serialReturnedVal := DoBigQuery(20)

	goRoutineQuery, concurrentReturnedVal := DoBigQueryInARoutine(50000)

	fmt.Printf("Time for serial query: %v\n", serialQuery)
	fmt.Printf("Time for concurrent query: %v\n", goRoutineQuery)

	fmt.Printf("sQ val: %v\n", serialReturnedVal)
	fmt.Printf("cQ val: %v\n", concurrentReturnedVal)

}
