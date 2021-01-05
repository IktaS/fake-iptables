# Fake IPTables

Fake IPTables is a simple program to simulate the act of [masquerading](https://en.wikipedia.org/wiki/Network_address_translation)

## Server

Currently, the server only does 2 things, when it recieves a request, it will read the body, determine if the source address should be accepted or not, and then accordingly deliver response, in this case, unaccepted source address will be returing "Not Permitted" as the message response, and "Hello!" as the message from accepted source address

## Router

Router does 2 things, it will map any request to another address HWAddr to their respective original source address, and then they will replace that request source address with their own, and then proceeds to forward that request to their destination. And it will replace the destination of any request coming to the router with their actual address based on the hardware address, and then forward it to the actual address

## Client

Client makes 2 request of the same content every 4 second, but the first request they send to the router, and the second they send directly to server. With this we can see how router masquerade the "source address" so that the server can accept it.
