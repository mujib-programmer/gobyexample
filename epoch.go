// A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanoseconds, respectively.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Note that there is no UnixMillis, so to get the milliseconds since epoch you’ll need to manually divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

// My output:
/*
2018-09-14 10:30:10.520372035 +0700 WIB m=+0.000340658
1536895810
1536895810520
1536895810520372035
2018-09-14 10:30:10 +0700 WIB
2018-09-14 10:30:10.520372035 +0700 WIB
*/
