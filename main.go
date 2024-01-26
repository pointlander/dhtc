// Copyright 2024 The DHTC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func main() {
	rng := rand.New(rand.NewSource(1))
	data := [sha256.Size + 4]byte{}
	for i := range data {
		data[i] = byte(rng.Uint32())
	}
	max := 0
	for n := 0; n < 32*1024*1024; n++ {
		data[0] = byte(n & 0xFF)
		data[1] = byte((n >> 8) & 0xFF)
		data[2] = byte((n >> 16) & 0xFF)
		data[3] = byte((n >> 24) & 0xFF)
		hash := sha256.Sum256(data[:])
		d := data[4:]
		same := 0
	outer:
		for i := range hash {
			x, y := hash[i], d[i]
			for j := 0; j < 8; j++ {
				if x&1 != y&1 {
					break outer
				}
				x >>= 1
				y >>= 1
				same++
			}
		}
		if same > max {
			max = same
		}
	}
	fmt.Println(max)
}
