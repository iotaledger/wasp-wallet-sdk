This is a PoC for now and untested.

Memory leaks will most likely build up as no controlled freeing of memory is in place.
Wrapper is very low level for now and does not offer fancy easy to use classes.

The native dll function `internal_listen_wallet` is not ported correctly yet.
It requires a handler callback and an array as parameters which are not yet adjusted.

It will most likely crash the application once called.

# Testing

As this is a wrapper for a native library, tests don't run out of the box. 

They require the compiled native lib installed on the machine, and it's not shipped in this repo.



