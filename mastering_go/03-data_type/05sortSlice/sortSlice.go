package main

import (
	"fmt"
	"sort"
)

type userInfo struct {
	name   string
	height int
	weight int
}

func gt(info []userInfo,a, b int) bool {
	return info[a].height > info[b].height
}

func lt(info []userInfo,a, b int) bool {
	return info[a].height < info[b].height
}

func main() {
	info := make([]userInfo, 0)
	info = append(info, userInfo{"Jack", 180, 90})
	info = append(info, userInfo{"Jay", 164, 45})
	info = append(info, userInfo{"Mary", 155, 45})
	info = append(info, userInfo{"Gina", 170, 48})
	fmt.Println("info:", info)

	sort.Slice(info, func(i, j int) bool {
		return info[i].height > info[j].height
	})
	fmt.Println("sort by height gt:", info)

	sort.Slice(info, func(i, j int) bool {
		return info[i].height < info[j].height
	})
	fmt.Println("sort by height lt:", info)
}
