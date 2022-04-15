## Tumaty
This is a little try to create a code in go to start the Pomodoro focus technic.


### Use 
```
  -focusTime string
    	The focus time, where you can start a music and focus on your activity [syntax 0h0m] (default "0h50m")
  -intervalTime string
    	The interval time, where you can relax before starting the focus time again [syntax 0h0m] (default "0h10m")
  -repetition int
    	The amount of a repetition (default 1)

```

### Stdout
```
Focus:       [ 02h01m ] 
Interval:    [ 0h01m ] 
Total focus: [ 10h:5m ]
Round:       [  1/5  ]

[Focus] TIMER -- 2h:0m:56s --
```
At the end of each iteration, a bell will ring and a message will be appearing.