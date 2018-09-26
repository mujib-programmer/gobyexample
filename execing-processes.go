// In the previous example we looked at spawning external processes. We do this when we need an external process accessible to a running Go process. Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one. To do this we’ll use Go’s implementation of the classic exec function.
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// For our example we’ll exec ls. Go requires an absolute path to the binary we want to execute, so we’ll use exec.LookPath to find it (probably /bin/ls).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form (as apposed to one big string). We’ll give ls a few common arguments. Note that the first argument should be the program name.
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use. Here we just provide our current environment.
	env := os.Environ()

	// Here’s the actual syscall.Exec call. If this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process. If there is an error we’ll get a return value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}

// When we run our program it is replaced by ls.
/*
$ go run execing-processes.go
total 296K
drwxr-xr-x 3 mujibur mujibur 4,0K Sep 26 13:13 .
drwxr-xr-x 6 mujibur mujibur 4,0K Sep 18 15:12 ..
-rw-r--r-- 1 mujibur mujibur 1013 Agu  6 13:49 arrays.go
-rw-r--r-- 1 mujibur mujibur 1,5K Agu  6 13:49 atomic-counters.go
-rw-r--r-- 1 mujibur mujibur 1,1K Sep 21 07:59 base64-encoding.go
-rw-r--r-- 1 mujibur mujibur  182 Agu  6 13:49 channel-buffering.go
-rw-r--r-- 1 mujibur mujibur  335 Agu  6 13:49 channel-directions.go
...
*/
