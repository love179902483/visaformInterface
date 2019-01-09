package main

import (
	"fmt"
	"getCPU"
	"log"

	"github.com/robfig/cron"
)

func main() {
	fmt.Println("CPU quantity", getCPU.GetCPUinfo())

	i := 0
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println(getCPU.GetCPUinfo())
		log.Println("cron ruing", i)
	})

	c.Start()
	select {}
}
