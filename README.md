# iotex-address

This is the golang library of the address used in IoTeX blockchain and relatant products and services.

## Address Generation Algorithm

A human readable address looks like `io1nyjs526mnqcsx4twa7nptkg08eclsw5c2dywp4`. It takes the following steps to be constructed:

1. Generating a random private key and the corresponding public key using secp256k1's elliptic curve;
2. Apply keccak256 hash function to the public key, exluding the first byte (hash := keccak256(pk[1:]);
3. Take the late 20 bytes as the payload (payload := hash[12:]), which is the byte representation of the address;
4. Apply [bech32](https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki) encoding on the payload and adding io prefix.
