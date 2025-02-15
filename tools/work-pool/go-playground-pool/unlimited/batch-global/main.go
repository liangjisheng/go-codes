package main

import (
	"fmt"
	"time"

	"gopkg.in/go-playground/pool.v3"
)

var gpool = pool.New()

func main() {

	// OK so maybe you want a long running pool to maximize throughput
	// yet limit the # of workers eg. email provider may limit the # of
	// concurrent connection you can have so spin up a pool with the #
	// of workers being that limit and then can batch
	// (or send per unit if desired) then can maximize email sending throughput
	// without breaking your providers limits.

	batch := gpool.Batch()

	// for max speed Queue in another goroutine
	// but it is not required, just can't start reading results
	// until all items are Queued.

	go func() {
		for i := 0; i < 10; i++ {
			batch.Queue(sendEmail("email content"))
		}

		// DO NOT FORGET THIS OR GOROUTINES WILL DEADLOCK
		// if calling Cancel() it calles QueueComplete() internally
		batch.QueueComplete()
	}()

	for email := range batch.Results() {

		if err := email.Error(); err != nil {
			// handle error
			// maybe call batch.Cancel()
		}

		// use return value
		fmt.Println(email.Value().(bool))
	}
}

func sendEmail(email string) pool.WorkFunc {

	return func(wu pool.WorkUnit) (interface{}, error) {

		// simulate waiting for something, like TCP connection to be established
		// or connection from pool grabbed
		time.Sleep(time.Second * 1)

		if wu.IsCancelled() {
			// return values not used
			return nil, nil
		}

		// ready for processing...

		return true, nil // everything ok, send nil, error if not
	}
}