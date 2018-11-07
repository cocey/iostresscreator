import psutil
import subprocess
import datetime




#define test scenarios
tests = {
    "default":{
        "command":['../iostress'],
        "description":"test with default values",
        "timeout": 5 #define timeout duration in seconds
    },
    "test1":{
        "command":['../iostress','-a','8','-b','4','-l','hard','-r'],
        "description":"test with basic config",
        "timeout":5,
    },
    "test2":{
        "command":['../iostress','-a','64','-b','8','-l','full','-r'],
        "description":"test with expensive options",
        "timeout":5,
    }
}

#results
results = ""
for test in tests:
    results += "---------------------\n"
    results += "starting test : "+test+"\n"
    results += "target timeout : "+str(tests[test]["timeout"])+" seconds\n"
    results += "starting time : "+str(datetime.datetime.now())+"\n"
    results += "desctiption : "+tests[test]["description"]+"\n"
    results += "command : "+" ".join(tests[test]["command"])+"\n"
    
    #start command
    subp = subprocess.Popen(tests[test]["command"], stdout=subprocess.PIPE)

    #get process id to terminate it
    p = psutil.Process(subp.pid)
    try:
        #wait untion end of timeout duration
        p.wait(timeout=tests[test]["timeout"])
    except psutil.TimeoutExpired:
        #killing process
        p.kill()
        pass
    #get command output and error signals
    out, err = subp.communicate()
    results += "command output : "+str(out)+"\n"
    results += "command errors : "+str(err)+"\n"
    results += "ending time : "+str(datetime.datetime.now())+"\n"
print(results)