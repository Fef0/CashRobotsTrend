# CashRobotsTrend
A Go script which calculates your earnings on Telegram bot "Cash Robots"

## Dependencies
- glot (go library for plotting)
    - gnu plot (glot dependency)
        - build gnu plot from [source](https://sourceforge.net/projects/gnuplot/files/gnuplot/)
        - linux users
           -  ```sudo apt-get update```
           -  ```sudo apt-get install gnuplot-x11``` 
        - mac users
           -  install homebrew
           -  ```brew cask install xquartz``` (for x-11)
           -  ```brew install gnuplot --with-x11```
    - Install glot itself: ```go get github.com/Arafatk/glot``` 
## Usage
This script is meant to be used on shell, it accepts the following arguments (all optionals):
- **-b** (int)
    -	Number of starting balance (default 300)
- **-d** (int)
    - Number of days for calculation (default 30)
- **-ub** (takes just the argument)
    -	Use bonus or not (default false)
- **-p** (string)
    -	Plot the trend graph (takes a string as the file name (**_png file only_**))
- **-v0** (int)
    -	Number of v0.01 robots (default 0)
- **-v1** (int)
    -	Number of v1.00 robots (default 0)
- -**v2** (int)
    -	Number of v2.00 robots (default 0)
- **-v3** (int)
    -	Number of v3.00 robots (default 0)
- **-v4** (int)
    -	Number of v4.00 robots (default 0)
- **-v5** (int)
    -	Number of v5.00 robots (default 0)
- **-v6** (int)
    -	Number of v6.00 robots (default 0)
- **-v7** (int)
    -	Number of v7.00 robots (default 0)
## Example
##### Command:
```go run CashRobotsTrend.go -d 7 -b 0 -ub -v0 100 -v1 10 -p Graph.png```
##### Output:
**Graph**

![](https://raw.githubusercontent.com/Fef0/CashRobotsTrend/master/Graph.png)
```
Welcome to RobotCashTrend!
This is your situation right now:
v0.01: 100
v1.00: 10
v2.00: 0
v3.00: 0
v4.00: 0
v5.00: 0
v6.00: 0
v7.00: 0

Day 1
*************************
Energy leftover: 0
Total energy produced today: 34080
Energy available: 34080
Total diamonds bought: 338
Diamonds division: 237 + 101
Bonus: 72
----------------------------------
     Available balance: 309
     Payout balance: 101
----------------------------------
Bought 3 robots v0.01
*************************


Day 2
*************************
Energy leftover: 280
Total energy produced today: 34584
Energy available: 34864
Total diamonds bought: 346
Diamonds division: 243 + 103
Bonus: 3
----------------------------------
     Available balance: 255
     Payout balance: 204
----------------------------------
Bought 2 robots v0.01
*************************


Day 3
*************************
Energy leftover: 264
Total energy produced today: 34920
Energy available: 35184
Total diamonds bought: 349
Diamonds division: 245 + 104
Bonus: 17
----------------------------------
     Available balance: 317
     Payout balance: 308
----------------------------------
Bought 3 robots v0.01
*************************


Day 4
*************************
Energy leftover: 284
Total energy produced today: 35424
Energy available: 35708
Total diamonds bought: 355
Diamonds division: 249 + 106
Bonus: 81
----------------------------------
     Available balance: 347
     Payout balance: 414
----------------------------------
Bought 3 robots v0.01
*************************


Day 5
*************************
Energy leftover: 208
Total energy produced today: 35928
Energy available: 36136
Total diamonds bought: 359
Diamonds division: 252 + 107
Bonus: 31
----------------------------------
     Available balance: 330
     Payout balance: 521
----------------------------------
Bought 3 robots v0.01
*************************


Day 6
*************************
Energy leftover: 236
Total energy produced today: 36432
Energy available: 36668
Total diamonds bought: 364
Diamonds division: 255 + 109
Bonus: 7
----------------------------------
     Available balance: 292
     Payout balance: 630
----------------------------------
Bought 2 robots v0.01
*************************


Day 7
*************************
Energy leftover: 268
Total energy produced today: 36768
Energy available: 37036
Total diamonds bought: 368
Diamonds division: 258 + 110
Bonus: 21
----------------------------------
     Available balance: 371
     Payout balance: 740
----------------------------------
Bought 3 robots v0.01
*************************

After 7 days this will be your situation:
v0.01: 119
v1.00: 10
v2.00: 0
v3.00: 0
v4.00: 0
v5.00: 0
v6.00: 0
v7.00: 0
Graph saved as Graph.png
```
