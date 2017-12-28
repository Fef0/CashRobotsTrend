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
This script is meant to be used on shell, it accepts the following arguments:
- **-b** (int)
    -	Number of starting balance (default 300)
- **-d** (int)
    - Number of days for calculation (default 30)
- **-ub** (takes just the argument)
    -	Use bonus or not (default false)
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
```go run CashRobotsTrend.go -d 7 -b 0 -ub -v0 100 -v1 10```
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
Bonus: 94
----------------------------------
     Available balance: 331
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
Bonus: 48
----------------------------------
     Available balance: 322
     Payout balance: 204
----------------------------------
Bought 3 robots v0.01
*************************


Day 3
*************************
Energy leftover: 264
Total energy produced today: 35088
Energy available: 35352
Total diamonds bought: 351
Diamonds division: 246 + 105
Bonus: 67
----------------------------------
     Available balance: 335
     Payout balance: 309
----------------------------------
Bought 3 robots v0.01
*************************


Day 4
*************************
Energy leftover: 252
Total energy produced today: 35592
Energy available: 35844
Total diamonds bought: 356
Diamonds division: 250 + 106
Bonus: 39
----------------------------------
     Available balance: 324
     Payout balance: 415
----------------------------------
Bought 3 robots v0.01
*************************


Day 5
*************************
Energy leftover: 244
Total energy produced today: 36096
Energy available: 36340
Total diamonds bought: 361
Diamonds division: 253 + 108
Bonus: 83
----------------------------------
     Available balance: 360
     Payout balance: 523
----------------------------------
Bought 3 robots v0.01
*************************


Day 6
*************************
Energy leftover: 240
Total energy produced today: 36600
Energy available: 36840
Total diamonds bought: 366
Diamonds division: 257 + 109
Bonus: 48
----------------------------------
     Available balance: 365
     Payout balance: 632
----------------------------------
Bought 3 robots v0.01
*************************


Day 7
*************************
Energy leftover: 240
Total energy produced today: 37104
Energy available: 37344
Total diamonds bought: 371
Diamonds division: 260 + 111
Bonus: 87
----------------------------------
     Available balance: 412
     Payout balance: 743
----------------------------------
Bought 4 robots v0.01
*************************

After 7 days this will be your situation:
v0.01: 122
v1.00: 10
v2.00: 0
v3.00: 0
v4.00: 0
v5.00: 0
v6.00: 0
v7.00: 0
```
