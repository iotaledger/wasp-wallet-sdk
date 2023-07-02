Internal wallet/secure storage library for Wasp, built on top of the IOTA SDK native lib.
For now, it's only purpose is to secure the seed in the wasp-cli, but could be extended if required.

Does not support `listen_wallet` as it's currently not required.

# Testing

As this is a wrapper for a native library, tests don't run out of the box.
They require the compiled native lib installed on the machine, and it's not shipped in this repo.

Instructions will follow once the native library is merged into iota-sdk.


