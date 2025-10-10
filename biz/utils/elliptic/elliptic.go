package elliptic

import "math/big"

type EllipticCurve struct {
	name string
	a    big.Int
	b    big.Int
	p    big.Int
	gx   big.Int
	gy   big.Int
}

func (ec *EllipticCurve) Name() string {
	return ec.name
}
func (ec *EllipticCurve) GetA() *big.Int {
	return &ec.a
}
func (ec *EllipticCurve) GetB() *big.Int {
	return &ec.b
}
func (ec *EllipticCurve) GetP() *big.Int {
	return &ec.p
}
func (ec *EllipticCurve) GetGx() *big.Int {
	return &ec.gx
}
func (ec *EllipticCurve) GetGy() *big.Int {
	return &ec.gy
}
