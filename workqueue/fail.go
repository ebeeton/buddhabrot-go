// Package workqueue allows callers to enqueue plot requests.
package workqueue

import "log"

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
