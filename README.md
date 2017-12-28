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
Bonus: 38
----------------------------------
     Available balance: 275
     Payout balance: 101
----------------------------------
Bought 2 robots v0.01
*************************


Day 2
*************************
Energy leftover: 280
Total energy produced today: 34416
Energy available: 34696
Total diamonds bought: 344
Diamonds division: 241 + 103
Bonus: 20
----------------------------------
     Available balance: 336
     Payout balance: 204
----------------------------------
Bought 3 robots v0.01
*************************


Day 3
*************************
Energy leftover: 296
Total energy produced today: 34920
Energy available: 35216
Total diamonds bought: 350
Diamonds division: 245 + 105
Bonus: 78
----------------------------------
     Available balance: 359
     Payout balance: 309
----------------------------------
Bought 3 robots v0.01
*************************


Day 4
*************************
Energy leftover: 216
Total energy produced today: 35424
Energy available: 35640
Total diamonds bought: 354
Diamonds division: 248 + 106
Bonus: 79
----------------------------------
     Available balance: 386
     Payout balance: 415
----------------------------------
Bought 3 robots v0.01
*************************


Day 5
*************************
Energy leftover: 240
Total energy produced today: 35928
Energy available: 36168
Total diamonds bought: 359
Diamonds division: 252 + 107
Bonus: 87
----------------------------------
     Available balance: 425
     Payout balance: 522
----------------------------------
Bought 4 robots v0.01
*************************


Day 6
*************************
Energy leftover: 268
Total energy produced today: 36600
Energy available: 36868
Total diamonds bought: 366
Diamonds division: 257 + 109
Bonus: 27
----------------------------------
     Available balance: 309
     Payout balance: 631
----------------------------------
Bought 3 robots v0.01
*************************


Day 7
*************************
Energy leftover: 268
Total energy produced today: 37104
Energy available: 37372
Total diamonds bought: 371
Diamonds division: 260 + 111
Bonus: 10
----------------------------------
     Available balance: 279
     Payout balance: 742
----------------------------------
Bought 2 robots v0.01
*************************

After 7 days this will be your situation:
v0.01: 120
v1.00: 10
v2.00: 0
v3.00: 0
v4.00: 0
v5.00: 0
v6.00: 0
v7.00: 0
Graph saved as Graph.png
```
