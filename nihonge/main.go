package main

import (
	"BBS04/bbs"
	"fmt"

	"github.com/Nik-U/pbc"
)

func GenerateGroup() *bbs.PrivateKey {
	params := pbc.GenerateA(160, 512)
	// serialization params
	str := params.String()
	fmt.Println("params is : ", str)
	newParams, _ := pbc.NewPairingFromString(str)
	// get g1 and g2
	pairing := newParams
	g1 := pairing.NewG1().Rand()
	g2 := pairing.NewG2().Rand()
	// generate Group
	return bbs.GenerateGroup(g1, g2, pairing)
}
func GenerateMember(managerkey *bbs.PrivateKey) *bbs.Cert {
	return managerkey.Cert()
}
func GenerateSig(memberkey *bbs.Cert, message string) *bbs.Sig {
	return bbs.Sign(memberkey, message)
}

func main() {
	priv1 := GenerateGroup()
	priv2 := GenerateGroup()
	m11 := GenerateMember(priv1)
	m21 := GenerateMember(priv2)
	bbs.Verify_cert(m11)
	bbs.Verify_cert(m21)
}
