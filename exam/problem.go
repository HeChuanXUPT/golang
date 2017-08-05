package main

func func1() int {
	m := map[string]Player{
		"tencent": {
			Name: "tencent",
		},
	}
	//m["tencent"].Money = 300
	return m["tencent"].Money
}
