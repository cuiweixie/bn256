package bn256

import (
	"math/big"
)

func bigFromBase10(s string) *big.Int {
	n, _ := new(big.Int).SetString(s, 10)
	return n
}

// p is a prime over which we form a basic field: 36u⁴+36u³+24u³+6u+1.
var p = bigFromBase10("65000549695646603732796438742359905742825358107623003571877145026864184071783")

var p2 = [4]uint64{0x185cac6c5e089667, 0xee5b88d120b5b59e, 0xaa6fecb86184dc21, 0x8fb501e34aa387f9}
var np = [4]uint64{0x2387f9007f17daa9, 0x734b3343ab8513c8, 0x2524282f48054c12, 0x38997ae661c3ef3c}

// var r = &gfP{0xe7a35393a1f76999, 0x11a4772edf4a4a61, 0x559013479e7b23de, 0x704afe1cb55c7806}
var r2 = &gfP{0x9c21c3ff7e444f56, 0x409ed151b2efb0c2, 0xc6dc37b80fb1651, 0x7c36e0e62c2380b7}

// Order is the number of elements in both G₁ and G₂: 36u⁴+36u³+18u³+6u+1.
// order-1 = (2**5) * 3 * 5743 * 280941149 * 130979359433191 * 491513138693455212421542731357 * 6518589491078791937
var Order = bigFromBase10("65000549695646603732796438742359905742570406053903786389881062969044166799969")
