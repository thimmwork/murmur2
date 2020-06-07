# murmur2
Command line tool to calculate murmur2 hash (as used by Kafka to determine partition for a key)

Go programming language implementation of MurmurHash2, based on the works by [Austin Appleby](https://code.google.com/p/smhasher/) and [David Irvine](https://github.com/aviddiviner/go-murmur)

## Usage
Uses 60 buckets by default. To use customize the number of buckets, use ```-n <number>```

    > murmur2 hash-me
    46
    > murmur2 -n 30 hash-me
    16

## License
[MIT](LICENSE)
