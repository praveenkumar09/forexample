package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["praveen"] = []string{`something here`, `some more thing here`, `some more more thing here`}

	for i, v := range m {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

	delete(m, "praveen")

	for i, v := range m {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

}
