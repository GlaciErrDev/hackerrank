package main

import (
	"sort"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
func nextInt() uint64 {
        sc.Scan()
        i, _ := strconv.Atoi(sc.Text())
        return uint64(i)
}

type City struct {
	population uint64
	num_of_clinics uint64
	load uint64
}

type ByLoad []City

func (a ByLoad) Len() int           { return len(a) }
func (a ByLoad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLoad) Less(i, j int) bool { return a[i].load < a[j].load }

/*
Sample input:
3 7
2
42
49
Output:
17
*/

func main() {
	var num_of_cities, num_of_clinics, i uint64
	var cities []City

	sc.Split(bufio.ScanWords)
	num_of_cities = nextInt()
	num_of_clinics = nextInt()

	for i = 0; i < num_of_cities; i++ {
		city := City{}
		// scan population
		city.population = nextInt()

		/* assign one clinic per city */
		city.num_of_clinics = 1

		/* Initial load per clinic is total population of city */
	    city.load = city.population
		cities = append(cities, city)
	}

	// sort cities
	sort.Sort(ByLoad(cities))

	for i = 0; i < (num_of_clinics - num_of_cities); i++ {
		cities[num_of_cities - 1].num_of_clinics++
		cities[num_of_cities - 1].load = cities[num_of_cities - 1].population / cities[num_of_cities - 1].num_of_clinics

		if bar := cities[num_of_cities - 1].population % cities[num_of_cities - 1].num_of_clinics; bar != 0{
			cities[num_of_cities - 1].load++
		}
		// sort cities
		sort.Sort(ByLoad(cities))
	}

	fmt.Printf("%d\n",cities[num_of_cities - 1].load)
}