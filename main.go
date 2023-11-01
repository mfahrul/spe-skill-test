package main

import (
	"fmt"
	"math"
	"strconv"
)

type SpeSkillTest struct{}

func (spe SpeSkillTest) IsNarcistic(n int) bool {

	pjg := len(strconv.Itoa(n))

	temp := n
	var total float64

	for temp > 0 {
		total += math.Pow(float64(temp%10), float64(pjg))
		temp = temp / 10
	}

	return n == int(total)
}

func (spe SpeSkillTest) Outlier(arrNya []int) interface{} {
	var ganjil, genap []int

	for i := 0; i < len(arrNya); i++ {
		if arrNya[i]%2 == 0 {
			genap = append(genap, arrNya[i])
		} else {
			ganjil = append(ganjil, arrNya[i])
		}
	}
	if len(ganjil) == 1 {
		return ganjil[0]
	} else if len(genap) == 1 {
		return genap[0]
	} else {
		return false
	}
}

func (spe SpeSkillTest) FindNeedle(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func (spe SpeSkillTest) BlueOcean(s []int, ins int) []int {

	s = findNRemove(s, ins)

	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }
	return s
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func findNRemove(data []int, element int) []int {
	for k, v := range data {
		if element == v {
			c := remove(data, k)
			data = findNRemove(c, element)
			// return k
		}
	}
	return data //not found.
}

func main() {
	h := 153

	spe := &SpeSkillTest{}
	fmt.Println(spe.IsNarcistic(h))

	arrs := []int{11, 13, 15, 19, 9, 13, -21}
	fmt.Println(spe.Outlier(arrs))

	needle := []string{"red", "blue", "yellow", "black", "grey"}

	fmt.Println(spe.FindNeedle(needle, "blue"))

	blueOc := []int{1, 5, 5, 5, 5, 3}

	fmt.Println(spe.BlueOcean(blueOc, 5))

}
