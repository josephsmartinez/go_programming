package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
  m[`joe`]= []string{`hello`,`as`,`as`}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

  in_map := m[`joe`]
  fmt.Println(`Is the value joe in the map? `in_map)
  delete(m,`joe`)

  for i, v := range m{
    fmt.Println("Record No." k)
    for j, v2 := range v{
      fmt.Println("\t",v2)
    }
  }



}
