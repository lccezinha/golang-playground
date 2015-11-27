// cria um channel Buffered

/*
This creates a buffered channel with a capacity of 1. 
Normally channels are synchronous; both sides of the channel will wait until the other side is ready. 
A buffered channel is asynchronous; 
sending or receiving a message will not wait unless the channel is already full.
*/

c := make(chan int, 1)