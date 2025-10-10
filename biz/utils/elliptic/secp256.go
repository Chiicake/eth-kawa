package elliptic

import (
	"math/big"
)

// SecP256 表示椭圆曲线secp256的参数
type SecP256 struct {
	A, B, P *big.Int // 曲线参数
	Gx, Gy  *big.Int // 生成点坐标
}

// K1 是secp256k1曲线的实例
var K1 = &SecP256{
	A: big.NewInt(0), // 参数a
	B: big.NewInt(7), // 参数b
}

func init() {
	// 初始化参数p: FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F
	p, ok := new(big.Int).SetString(
		"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F",
		16,
	)
	if ok {
		K1.P = p
	}

	// 初始化生成点G的横坐标
	gx, ok := new(big.Int).SetString(
		"79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798",
		16,
	)
	if ok {
		K1.Gx = gx
	}

	// 初始化生成点G的纵坐标
	gy, ok := new(big.Int).SetString(
		"483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8",
		16,
	)
	if ok {
		K1.Gy = gy
	}
}

// GetA 获取参数a
func (s *SecP256) GetA() *big.Int {
	return s.A
}

// GetB 获取参数b
func (s *SecP256) GetB() *big.Int {
	return s.B
}

// GetP 获取参数p
func (s *SecP256) GetP() *big.Int {
	return s.P
}

// GetGx 获取生成点G的横坐标
func (s *SecP256) GetGx() *big.Int {
	return s.Gx
}

// GetGy 获取生成点G的纵坐标
func (s *SecP256) GetGy() *big.Int {
	return s.Gy
}
