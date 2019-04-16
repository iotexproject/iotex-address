# iotex-address

This is the golang library of the address used in IoTeX blockchain and relatant products and services.

## Address Generation Algorithm

A human readable address looks like `io1nyjs526mnqcsx4twa7nptkg08eclsw5c2dywp4`. It takes the following steps to be constructed:

1. Generating public key (`pk`) by using secp256k1's elliptic curve.
2. Apply keccak256 hash function public key bytes, exluding the first byte (`hash := keccak256(pk[1:]`).
3. Taking the late 20 bytes as the payload (`payload := hash[12:]`), which is the byte representation of the address.
4. Appying bech32 encoding on the payload and adding `io` prefix.
