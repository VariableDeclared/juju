// Copyright 2021 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package queue

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/juju/clock/testclock"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
)

type BlockingOpQueueSuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(&BlockingOpQueueSuite{})

func (s *BlockingOpQueueSuite) TestEnqueue(c *gc.C) {
	now := time.Now()
	queue := NewBlockingOpQueue(testclock.NewClock(now))

	results := consumeN(c, queue, 1)

	err := queue.Enqueue(Operation{
		Command:  opName(0),
		Deadline: now.Add(time.Second),
	})
	c.Assert(err, jc.ErrorIsNil)

	var count int
	for result := range results {
		c.Assert(result, gc.DeepEquals, opName(count))
		count++
	}
	c.Assert(count, gc.Equals, 1)
}

func (s *BlockingOpQueueSuite) TestEnqueueWithError(c *gc.C) {
	now := time.Now()
	queue := NewBlockingOpQueue(testclock.NewClock(now))

	results := consumeNUntilErr(c, queue, 1, errors.New("boom"))

	err := queue.Enqueue(Operation{
		Command:  opName(0),
		Deadline: now.Add(time.Second),
	})
	c.Assert(err, gc.ErrorMatches, `boom`)

	var count int
	for result := range results {
		c.Assert(result, gc.DeepEquals, opName(count))
		count++
	}
	c.Assert(count, gc.Equals, 1)
}

func (s *BlockingOpQueueSuite) TestEnqueueTimesout(c *gc.C) {
	now := time.Now()
	clock := testclock.NewClock(now)
	queue := NewBlockingOpQueue(clock)

	go func() {
		c.Assert(clock.WaitAdvance(time.Millisecond, testing.ShortWait, 1), jc.ErrorIsNil)
	}()

	err := queue.Enqueue(Operation{
		Command:  []byte("abc-1"),
		Deadline: now.Add(time.Nanosecond),
	})
	c.Assert(err, gc.ErrorMatches, `deadline exceeded`)
}

func (s *BlockingOpQueueSuite) TestMultipleEnqueue(c *gc.C) {
	now := time.Now()
	queue := NewBlockingOpQueue(testclock.NewClock(now))

	results := consumeN(c, queue, 2)

	for i := 0; i < 2; i++ {
		err := queue.Enqueue(Operation{
			Command:  opName(i),
			Deadline: now.Add(time.Second),
		})
		c.Assert(err, jc.ErrorIsNil)
	}

	var count int
	for result := range results {
		c.Assert(result, gc.DeepEquals, opName(count))
		count++
	}
	c.Assert(count, gc.Equals, 2)
}

func (s *BlockingOpQueueSuite) TestMultipleEnqueueWithErrors(c *gc.C) {
	now := time.Now()
	clock := testclock.NewClock(now)
	queue := NewBlockingOpQueue(clock)

	results := make(chan []byte, 3)
	go func() {
		defer close(results)

		var count int
		for op := range queue.Queue() {
			results <- op.Command
			queue.Error() <- nil

			count++
			switch count {
			case 1:
				time.Sleep(time.Millisecond * 500)
				count++
			case 3:
				return
			}
		}
	}()

	err := queue.Enqueue(Operation{
		Command:  opName(0),
		Deadline: now.Add(time.Nanosecond),
	})
	c.Assert(err, jc.ErrorIsNil)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		c.Assert(clock.WaitAdvance(time.Millisecond, testing.ShortWait, 2), jc.ErrorIsNil)
	}()

	// Fail this one
	err = queue.Enqueue(Operation{
		Command:  opName(1),
		Deadline: now.Add(time.Nanosecond),
	})
	c.Assert(err, gc.ErrorMatches, `deadline exceeded`)

	err = queue.Enqueue(Operation{
		Command:  opName(2),
		Deadline: now.Add(time.Millisecond * 100),
	})
	c.Assert(err, jc.ErrorIsNil)

	var received []string
	for result := range results {
		received = append(received, string(result))
	}
	c.Assert(len(received), gc.Equals, 2)
	c.Assert(received, gc.DeepEquals, []string{
		"abc-0", "abc-2",
	})

	// Ensure that we actually did advance correctly.
	wg.Wait()
}

func (s *BlockingOpQueueSuite) TestMultipleEnqueues(c *gc.C) {
	now := time.Now()
	queue := NewBlockingOpQueue(testclock.NewClock(now))

	results := consumeN(c, queue, 10)

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			err := queue.Enqueue(Operation{
				Command:  opName(i),
				Deadline: now.Add(time.Second),
			})
			c.Assert(err, jc.ErrorIsNil)
		}(i)
	}
	wg.Wait()

	var received []string
	for result := range results {
		received = append(received, string(result))
	}
	c.Assert(len(received), gc.Equals, 10)
	c.Assert(received, jc.SameContents, []string{
		"abc-0", "abc-1", "abc-2", "abc-3", "abc-4",
		"abc-5", "abc-6", "abc-7", "abc-8", "abc-9",
	})
}

func opName(i int) []byte {
	return []byte(fmt.Sprintf("abc-%d", i))
}

func consumeN(c *gc.C, queue *BlockingOpQueue, n int) <-chan []byte {
	return consumeNUntilErr(c, queue, n, nil)
}

func consumeNUntilErr(c *gc.C, queue *BlockingOpQueue, n int, err error) <-chan []byte {
	results := make(chan []byte, n)

	go func() {
		defer close(results)

		var count int
		for op := range queue.Queue() {
			select {
			case results <- op.Command:
			case <-time.After(testing.LongWait):
				c.Fatal("timed out setting results")
			}

			count++
			var e error
			if count == n {
				e = err
			}
			select {
			case queue.Error() <- e:
			case <-time.After(testing.LongWait):
				c.Fatal("timed out setting error")
			}

			if count == n {
				break
			}
		}
	}()

	return results
}
