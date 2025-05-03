package render

import "fmt"

func Logo() {
	logo := []string{
		`    _     _____   _______  _____  `,
		`   / \   |  -- \ |__   __||  -- \ `,
		`  / _ \  | |__) |   | |   | |__) |`,
		` / ___ \ |  _  /    | |   |  _  / `,
		`/_/   \_\|_| \_\    |_|   |_| \_\`,
		`                                 `,
		`   Aten Remote Task Runner       `,
		`   https://www.atentec.com       `,
	}

	for _, l := range logo {
		fmt.Printf("\033[1;33m%v\033[0m\n", l) // Yellow bold output
	}
}
