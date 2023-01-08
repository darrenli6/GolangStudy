package main

import "fmt"

func main() {
	info := Info{"kube-system", "Pod", "config"}
	visit := func(info *Info, err error) error {
		fmt.Printf("%s,%s,%s\n", info.Namespace, info.Name, info.OtherThings)
		return nil
	}
	info.Visit(visit)
}
