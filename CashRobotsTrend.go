package main

import (
  "fmt"
  "time"
  "math/rand"
  "github.com/Arafatk/glot"
  "flag"
)


type Robot struct{
  name string
  quantity int
  price int
  max_value int
}

func buy(balance *int, robot string, n_bots *int, min int, max int) {
  //Takes balance, robot name, number of robots,price and max_value
  if *balance >= min && *balance < max {
    old := *n_bots
    for *balance >= min{
      *balance -= min
      *n_bots += 1
    }
    fmt.Println("Bought",*n_bots-old,"robots",robot)
  }
}


func main(){

  //Set base values of robots
  robots := []Robot{
    {name:"v0.01", quantity:0, price:100, max_value:1000},
    {name:"v1.00", quantity:0, price:1000, max_value:5000},
    {name:"v2.00", quantity:0, price:5000, max_value:25000},
    {name:"v3.00", quantity:0, price:25000, max_value:100000},
    {name:"v4.00", quantity:0, price:100000, max_value:250000},
    {name:"v5.00", quantity:0, price:250000, max_value:500000},
    {name:"v6.00", quantity:0, price:500000, max_value:900000},
    {name:"v7.00", quantity:0, price:900000, max_value:9223372036854775807},
  }

  toPlot := [][]int{{},{}}  //2D slice of all the points to be plotted

  var days int = 0
  var balance int = 0
  var energy_produced int = 0
  var total_energy int = 0
  var payout int = 0
  var graph_file string
  var useBonus bool
  var bonus int = 0

  //Set flag
  flag.IntVar(&days,"d",30,"Number of days for calculation")
  flag.IntVar(&balance,"b",300,"Number of starting balance")
  flag.BoolVar(&useBonus,"ub",false,"Use bonus or not")
  flag.IntVar(&robots[0].quantity,"v0",0,"Number of v0.01 robots")
  flag.IntVar(&robots[1].quantity,"v1",0,"Number of v1.00 robots")
  flag.IntVar(&robots[2].quantity,"v2",0,"Number of v2.00 robots")
  flag.IntVar(&robots[3].quantity,"v3",0,"Number of v3.00 robots")
  flag.IntVar(&robots[4].quantity,"v4",0,"Number of v4.00 robots")
  flag.IntVar(&robots[5].quantity,"v5",0,"Number of v5.00 robots")
  flag.IntVar(&robots[6].quantity,"v6",0,"Number of v6.00 robots")
  flag.IntVar(&robots[7].quantity,"v7",0,"Number of v7.00 robots")
  flag.StringVar(&graph_file,"p","","Plot graphs (png file only)")
  flag.Parse()

  fmt.Println("\nWelcome to RobotCashTrend!\nThis is your situation right now:")
  for rb := range robots{
    fmt.Println(robots[rb].name+":",robots[rb].quantity)
  }


  //Set the random seed
  rand.Seed(time.Now().UTC().UnixNano())

  //Calculate the production for each day
  for day := 1; day <= days; day++ {
    fmt.Println()
    fmt.Println("Day",day)
    fmt.Println("*************************")
    fmt.Println("Energy leftover:",total_energy)

    //Calculate the energy production for each day
    energy_produced = ((robots[0].quantity*7) + (robots[1].quantity*72) +
                      (robots[2].quantity*375) +(robots[3].quantity*1920) +
                      (robots[4].quantity*9600) + (robots[5].quantity*25000) +
                      (robots[6].quantity*60000) + (robots[7].quantity*120000))*24

    //Update the total energy with the daily production
    total_energy += energy_produced

    fmt.Println("Total energy produced today:",energy_produced)
    fmt.Println("Energy available:",total_energy)

    //Turns energy in diamonds
    diamonds_produced := 0
    for total_energy >= 300{
      total_energy -= 100
      diamonds_produced += 1
    }

    //Calculate the daily bonus randomly (1,100 range)
    if useBonus {
      bonus = 1 + rand.Intn(100-1)
    }

    //Calculate the percentages of the payout and balance and update them
    perc := (diamonds_produced * 30) / 100
    payout += perc
    balance += (diamonds_produced - perc) + bonus

    fmt.Println("Total diamonds bought:",diamonds_produced)
    fmt.Println("Diamonds division:",diamonds_produced-perc,"+",perc)
    fmt.Println("Bonus:",bonus)
    fmt.Println("----------------------------------")
    fmt.Println("     Available balance:",balance)
    fmt.Println("     Payout balance:",payout)
    fmt.Println("----------------------------------")

    //Add points to be plotted (x=day, y=diamonds)
    toPlot[0] = append(toPlot[0],day)
    toPlot[1] = append(toPlot[1],diamonds_produced)

    //Buy operations
    for i := range robots{
      buy(&balance, robots[i].name, &robots[i].quantity, robots[i].price, robots[i].max_value)
    }

    fmt.Println("*************************\n")
  }

  //Print final result
  fmt.Println("After",days,"days this will be your situation:")
  for rb := range robots{
    fmt.Println(robots[rb].name+":",robots[rb].quantity)
  }

  if graph_file != "" {
    //Plot a graph of diamonds production trend
    plot, _ := glot.NewPlot(2,false,false)  //NewPlot(dimensions, persist, debug)
  	plot.AddPointGroup("Diamonds earning growing", "lines", toPlot)
  	plot.SetTitle("Cash Robots trend")
  	plot.SetXLabel("Days")
  	plot.SetYLabel("Diamonds")
  	plot.SavePlot(graph_file)
    fmt.Println("Graph saved as "+graph_file)
  }

}
