package main

import (
	"os/exec"
	simModxml "GoSimulation/modxml"
	simLog "GoSimulation/log"
	"flag"
)


func main() {
	seedType := flag.String("seedType", "indoor", "input seed type: indoor or 19sector")
	flag.Parse()

	ueNums := []int{30}
	for _, ueNum := range ueNums {
		simModxml.ModChardware(ueNum, *seedType)
		callSimulation()
	}
}

func callSimulation() error{
	logger := simLog.NewLog()
	cmd := exec.Command("Simulation", "-t 10")
	err := cmd.Run()
	if err != nil {
		logger.Fatal("Simulation exec failure!")
		return err
	}
	return nil
}
