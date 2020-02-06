# dlog
'dlog' provides debug output that makes it easy to see where and what is happening.
This package outputs the name of the function that requested the output 
and the number of lines of source code, along with a log message.

This log mechanism converts any type of value into a character string 
if it can be referenced from the outside, and outputs it.

This makes it easier to output 'struct', 'map', 'array' and 'slice'.

# Installation
`go get -u github.com/WhiteRaven777/dlog`
