/*
Buffered Channels:
- Synchronous way of communication between the goroutines
- ie. when pinger(sender) sends the msg to channel,
    it will wait until the printer(receiver) is ready to receive the message.
- also called as `blocking`
*/

package channels
