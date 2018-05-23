package main

import "la7lo7/types"

func main() {
	bc := types.NewBlockchain()
	defer bc.Db.Close()

	cli := CLI{bc}
	cli.Run()
}
