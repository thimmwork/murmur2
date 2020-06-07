// Based on the MurmurHash2.cpp source from SMHasher & MurmurHash,
// https://code.google.com/p/smhasher/

// Copyright (c) 2015 David Irvine

package main

// Mixing constants; generated offline.
const (
	M     = 0x5bd1e995
	BIG_M = 0xc6a4a7935bd1e995
	R     = 24
	BIG_R = 47
)

// 32-bit mixing function.
func mmix(h uint32, k uint32) (uint32, uint32) {
	k *= M
	k ^= k >> R
	k *= M
	h *= M
	h ^= k
	return h, k
}

// The original MurmurHash2 32-bit algorithm by Austin Appleby.
func MurmurHash2(data []byte) (h uint32) {
	var k uint32

	// Initialize the hash to a 'random' value
	h = 0x9747b28c ^ uint32(len(data))

	// Mix 4 bytes at a time into the hash
	for l := len(data); l >= 4; l -= 4 {
		k = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
		h, k = mmix(h, k)
		data = data[4:]
	}

	// Handle the last few bytes of the input array
	switch len(data) {
	case 3:
		h ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(data[0])
		h *= M
	}

	// Do a few final mixes of the hash to ensure the last few bytes are well incorporated
	h ^= h >> 13
	h *= M
	h ^= h >> 15

	return
}
