# Test Results
## Load Test
tested with custom python script
``` sh
Tool : Custom Python Script
Test file : docs/test.py
```

### test results
``` sh
---------------------
starting test : default
target timeout : 5 seconds
starting time : 2018-05-28 14:22:42.901833
desctiption : test with default values
command : ./iostress
command output : 
command errors : None
ending time : 2018-05-28 14:22:47.924481
---------------------
starting test : test1
target timeout : 5 seconds
starting time : 2018-05-28 14:22:47.924506
desctiption : test with basic config
command : ./iostress -a 8 -b 4 -l hard -r
command output : 
command errors : None
ending time : 2018-05-28 14:22:52.948512
---------------------
starting test : test2
target timeout : 5 seconds
starting time : 2018-05-28 14:22:52.948537
desctiption : test with expensive options
command : ./iostress -a 64 -b 8 -l full -r
command output : 
command errors : None
ending time : 2018-05-28 14:22:57.976648```
```


## Security Test
There is no security obligation for this application

