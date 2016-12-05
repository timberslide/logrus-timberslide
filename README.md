# logrus-timberslide

logrus-timberslide is a Logrus hook for Timberslide!

Example usage:

```
package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/timberslide/logrus-timberslide"
)

func main() {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	hook, err := tslogrus.NewTimberslideHook(
		&tslogrus.Hook{
			Token: "TqPeDFG30Fz147NeCB59SUgT",
			Topic: "kris/logrustest",
		})
	if err != nil {
		log.Fatalln(err)
	}
	log.Hooks.Add(hook)

	for {
		log.WithField("Example", "Field").Infoln("Hello Timberslide!")
		time.Sleep(1 * time.Second)
	}
}
```
