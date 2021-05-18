package mpt

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/joeqian10/neo3-gogogo/helper"
)

func TestVerifyProof(t *testing.T) {
	proofStr := "080c0000000102050008720003a69d0798767e0166c14d03e31fc81e147b91287560e7fa19d1561f29759b957703a9495885fc335dd5daa222e49ce0d542c8d6b1094ca51e30b5b2ea17d9b6facc0404040404040404040404040403251d0de36de3bfd935fee7c7c3bee40f68777b87e08f4b3c9eb6164fac558e3a04fd120100040403de0449d75800e662cc5c07b804e72cb2aa7e6bb3c3c8812fa178587b9586089b040404040403ebd5e01d9738d1158597b04e452deb599305fef2f92c1f8e1a112cfb7f94d8830403a065ed1a35318147644d985278da6c38a6c0cfaac1856b388e4f8cb492afd91103de0449d75800e662cc5c07b804e72cb2aa7e6bb3c3c8812fa178587b9586089b03652b8a83ead517596149de57733acfb988d709f46a78fd2e1d7cdca7798e2d16034ee881ac96d9a53b8fc20f18cf4fe7e2ee998276272d385d2ca3db57776022ce03aa2dc96edb7ef06ecf9674559046b8dbd03e1b3c7a1f3a62eaefd9ea13b64b9603d90b274d165198c2045263551c524a390b70ccc13ca527a076db7b5cf714de7e042a01070000000000000003791d0651089230b366bc3bb286aaf6008e7298803ebdaf6a92e92204879717ed92000403a7de2e8df729aa720e5bda2dd4903a092aeb6eb147300a31f7b54b016544e31d03d340ff18882c9b4f6b5dee49b6be575938019edbad4404633e8f27f65d41c2a803f1f3e3558f3a751b5bb2511493e19f382af6b80d38ae306881e3db4bb1d0e4a604037fdf37ec7a2b533746412242f763d665c20593baaabb2f6a528a51fa4d11172b04040404040404040404042401010003538699ec1fb90dfc037ced3617f4d0c8d941a72c5a24e868a5ad96d8d29eb6b352000403904fb8e816446e1b3d1501783006a2f01fb13166117497ed714cc74a385221ad03a2497231297a2a46abf007954393ae6b3e7d9c1f09c3f869aefcaeb8425eb74504040404040404040404040404042701040005000003497d33633b08061b6cc8882bc5a797cb048d1ff916fd984f4bd09671f64fb88dc902c700201ca7d4dcecd3705602d6e285deafa2c8240f2883045d237822ae33d50451f59f208d735f64159b5f9c7a5670dbab887dbb895c012f39003e5fbc55b6fbff9df44414cc9f88a9e96be8e91131b06ea674fbba51c7c99e0500000000000000144f5f702b3f459f222d371052940bb9ce2d86d2ed06756e6c6f636b4a14e14fdd69cf7bf6afb9265ac806e09fea438df7b81425820465d41a57dca24529e88387ac2d787227780f42400000000000000000000000000000000000000000000000000000000000"
	proofData := helper.HexToBytes(proofStr)

	root, _ := helper.UInt256FromString("0x721697bf93a8f96ed125ba481585b2f7e604e962262062df92ce7c7448101bf1")

	id, key, proofs, err := ResolveProof(proofData)
	assert.Nil(t, err)

	value, err := VerifyProof(root, id, key, proofs)
	assert.Nil(t, err)

	assert.Equal(t, "00201ca7d4dcecd3705602d6e285deafa2c8240f2883045d237822ae33d50451f59f208d735f64159b5f9c7a5670dbab887dbb895c012f39003e5fbc55b6fbff9df44414cc9f88a9e96be8e91131b06ea674fbba51c7c99e0500000000000000144f5f702b3f459f222d371052940bb9ce2d86d2ed06756e6c6f636b4a14e14fdd69cf7bf6afb9265ac806e09fea438df7b81425820465d41a57dca24529e88387ac2d787227780f42400000000000000000000000000000000000000000000000000000000000", helper.BytesToHex(value))
}
