1) Why don’t we see “goroutine hello”?
We don’t see “goroutine hello” because the main function finishes before the goroutine has time to run.

2) Why is the final counter value not 1000?
The final value is not 1000 because of a race condition. Multiple goroutines update the same variable at the same time without synchronization.