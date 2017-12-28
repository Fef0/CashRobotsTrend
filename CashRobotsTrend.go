package main

import "fmt"
import "github.com/Arafatk/glot"

func main(){

  var days int = 60
  var v0 int = 208
  var v1 int = 15
  var v2 int = 0
  var v3 int = 0
  var v4 int = 0
  var v5 int = 0
  var v6 int = 0
  var v7 int = 0
  var left int = 0
  var diamonds int = 300
  var payout int= 0
  var temp int = 0
  toPlot := [][]int{{},{}}
  fmt.Println()

  for i := 1; i<=days;i++ {
    fmt.Println("Day",i)
    left += ((v0*7) + (v1*72) + (v2*375) + (v3*1920) + (v4*9600) + (v5*25000) + (v6*60000) + (v7*120000))*24

    fmt.Println("Total energy produced:",left)
    temp = 0
    for left >= 300{
      left -= 100
      temp += 1
    }

    perc := (temp * 30) / 100
    payout += perc

    diamonds += (temp - perc)
    fmt.Println("Total diamonds bought:",temp)
    fmt.Println("Payout:",payout)
    fmt.Println("Diamonds available:",diamonds)
    toPlot[0] = append(toPlot[0],i)
    toPlot[1] = append(toPlot[1],temp)
    if diamonds >= 100 && diamonds < 1000 {
      old := v0
      for diamonds >= 100{
        diamonds -= 100
        v0 += 1
      }
      fmt.Println("I've got",v0-old,"v0.1 robots")
    }

    if diamonds >= 1000 && diamonds < 5000 {
      old := v1
      for diamonds >= 1000{
        diamonds -= 1000
        v1 += 1
      }
      fmt.Println("I've got",v1-old,"v1 robots")
    }

    if diamonds >= 5000 && diamonds < 25000 {
      old := v2
      for diamonds >= 5000{
        diamonds -= 5000
        v2 += 1
      }
      fmt.Println("I've got",v2-old,"v2 robots")
    }

    if diamonds >= 25000 && diamonds < 100000 {
      old := v3
      for diamonds >= 25000{
        diamonds -= 25000
        v3 += 1
      }
      fmt.Println("I've got",v3-old,"v3 robots")
    }

    if diamonds >= 100000 && diamonds < 250000 {
      old := v4
      for diamonds >= 100000{
        diamonds -= 100000
        v4 += 1
      }
      fmt.Println("I've got",v4-old,"v4 robots")
    }

    if diamonds >= 250000 && diamonds < 500000 {
      old := v5
      for diamonds >= 250000{
        diamonds -= 250000
        v5 += 1
      }
      fmt.Println("I've got",v5-old,"v5 robots")
    }

    if diamonds >= 500000 && diamonds < 900000 {
      old := v6
      for diamonds >= 500000{
        diamonds -= 500000
        v6 += 1
      }
      fmt.Println("I've got",v6-old,"v6 robots")
    }

    if diamonds >= 900000 {
      old := v7
      for diamonds >= 900000{
        diamonds -= 900000
        v7 += 1
      }
      fmt.Println("I've got",v7-old,"v7 robots")
    }

    fmt.Println()
  }
  fmt.Println("Recap:")
  fmt.Printf("v0.01: %v\nv1.00: %v\nv2.00: %v\nv3.00: %v\nv4.00: %v\nv5.00: %v\nv6.00: %v\nv7.00: %v\n",v0,v1,v2,v3,v4,v5,v6,v7)

  plot, _ := glot.NewPlot(2,true,false)
        // Adding a point group
	plot.AddPointGroup("Diamonds earning growing", "lines", toPlot)
  plot.AddPointGroup("Diamonds earning growing", "circles", toPlot)
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle("Cash Robots trend")
	// Optional: Setting the title of the plot
	plot.SetXLabel("Days")
	plot.SetYLabel("Diamonds")
	plot.SavePlot("cash.png")

  fmt.Println("-------------------------------\n")
}
