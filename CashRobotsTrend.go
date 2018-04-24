package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/Arafatk/glot"
)

// Robot represents a robot instance
type Robot struct {
	name     string
	quantity int
	price    int
	maxValue int
}

func buy(balance *int, robot string, nBots *int, min int, max int) {
	//Takes balance, robot name, number of robots,price and maxValue
	if *balance >= min && *balance < max {
		old := *nBots
		for *balance >= min {
			*balance -= min
			*nBots++
		}
		fmt.Println("Bought", *nBots-old, "robots", robot)
	}
}

func main() {

	//Set base values of robots
	robots := []Robot{
		{name: "v0.01", quantity: 0, price: 100, maxValue: 1000},
		{name: "v1.00", quantity: 0, price: 1000, maxValue: 5000},
		{name: "v2.00", quantity: 0, price: 5000, maxValue: 25000},
		{name: "v3.00", quantity: 0, price: 25000, maxValue: 100000},
		{name: "v4.00", quantity: 0, price: 100000, maxValue: 250000},
		{name: "v5.00", quantity: 0, price: 250000, maxValue: 500000},
		{name: "v6.00", quantity: 0, price: 500000, maxValue: 900000},
		{name: "v7.00", quantity: 0, price: 900000, maxValue: 9223372036854775807},
	}

	toPlot := [][]int{{}, {}} //2D slice of all the points to be plotted

	var days int
	var balance int
	var energyProduced int
	var totalDiamonds int
	var totalEnergy int
	var payout int
	var graphFile string
	var useBonus bool
	var bonus int

	//Set flag
	flag.IntVar(&days, "d", 30, "Number of days for calculation")
	flag.IntVar(&balance, "b", 300, "Number of starting balance")
	flag.BoolVar(&useBonus, "ub", false, "Use bonus or not")
	flag.IntVar(&robots[0].quantity, "v0", 0, "Number of v0.01 robots")
	flag.IntVar(&robots[1].quantity, "v1", 0, "Number of v1.00 robots")
	flag.IntVar(&robots[2].quantity, "v2", 0, "Number of v2.00 robots")
	flag.IntVar(&robots[3].quantity, "v3", 0, "Number of v3.00 robots")
	flag.IntVar(&robots[4].quantity, "v4", 0, "Number of v4.00 robots")
	flag.IntVar(&robots[5].quantity, "v5", 0, "Number of v5.00 robots")
	flag.IntVar(&robots[6].quantity, "v6", 0, "Number of v6.00 robots")
	flag.IntVar(&robots[7].quantity, "v7", 0, "Number of v7.00 robots")
	flag.StringVar(&graphFile, "p", "", "Plot graphs (png file only)")
	flag.Parse()

	fmt.Println("\nWelcome to RobotCashTrend!\nThis is your situation right now:")
	for rb := range robots {
		fmt.Println(robots[rb].name+":", robots[rb].quantity)
	}

	//Set the random seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Calculate the production for each day
	for day := 1; day <= days; day++ {
		fmt.Println()
		fmt.Println("Day", day)
		fmt.Println("*************************")
		fmt.Println("Energy leftover:", totalEnergy)

		//Calculate the energy production for each day
		energyProduced = ((robots[0].quantity * 7) + (robots[1].quantity * 72) +
			(robots[2].quantity * 375) + (robots[3].quantity * 1920) +
			(robots[4].quantity * 9600) + (robots[5].quantity * 25000) +
			(robots[6].quantity * 60000) + (robots[7].quantity * 120000)) * 24

		//Update the total energy with the daily production
		totalEnergy += energyProduced

		fmt.Println("Total energy produced today:", energyProduced)
		fmt.Println("Energy available:", totalEnergy)

		//Turns energy in diamonds
		diamondsProduced := 0
		for totalEnergy >= 300 {
			totalEnergy -= 100
			diamondsProduced++
		}

		//Calculate the daily bonus randomly (1,100 range)
		if useBonus {
			bonus = 1 + rand.Intn(100-1)
		}

		//Calculate the percentages of the payout and balance and update them
		perc := (diamondsProduced * 30) / 100
		payout += perc
		balance += (diamondsProduced - perc) + bonus
		totalDiamonds += balance

		fmt.Println("Total diamonds produced so far:", totalDiamonds)
		fmt.Println("Total diamonds bought:", diamondsProduced)
		fmt.Println("Diamonds division:", diamondsProduced-perc, "+", perc)
		fmt.Println("Bonus:", bonus)
		fmt.Println("----------------------------------")
		fmt.Println("     Available balance:", balance)
		fmt.Println("     Payout balance:", payout)
		fmt.Println("----------------------------------")

		//Add points to be plotted (x=day, y=diamonds)
		toPlot[0] = append(toPlot[0], day)
		toPlot[1] = append(toPlot[1], diamondsProduced)

		//Buy operations
		for i := range robots {
			buy(&balance, robots[i].name, &robots[i].quantity, robots[i].price, robots[i].maxValue)
		}

		fmt.Println("*************************")
		fmt.Println()
	}

	//Print final result
	fmt.Println("After", days, "days this will be your situation:")
	for rb := range robots {
		fmt.Println(robots[rb].name+":", robots[rb].quantity)
	}

	if graphFile != "" {
		//Plot a graph of diamonds production trend
		plot, _ := glot.NewPlot(2, false, false) //NewPlot(dimensions, persist, debug)
		plot.AddPointGroup("Diamonds earning growing", "lines", toPlot)
		plot.SetTitle("Cash Robots trend")
		plot.SetXLabel("Days")
		plot.SetYLabel("Diamonds")
		plot.SavePlot(graphFile)
		fmt.Println("Graph saved as " + graphFile)
	}

}
