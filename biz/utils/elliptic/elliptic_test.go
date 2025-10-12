package elliptic

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math/big"
	"testing"
)

func TestEllipticCurveCreation(t *testing.T) {
	curve, err := NewEllipticCurve(K1.A, K1.B, K1.P)
	if err != nil {
		t.Errorf("NewEllipticCurve failed: %v", err)
	}
	// 检查曲线是否为奇异曲线
	if curve.IsSingular() {
		t.Errorf("NewEllipticCurve created a singular curve")
	}
	// 检查曲线参数是否正确
	if curve.A.Cmp(K1.A) != 0 {
		t.Errorf("NewEllipticCurve created with incorrect a parameter")
	}
	if curve.B.Cmp(K1.B) != 0 {
		t.Errorf("NewEllipticCurve created with incorrect b parameter")
	}
	if curve.P.Cmp(K1.P) != 0 {
		t.Errorf("NewEllipticCurve created with incorrect p parameter")
	}
}

// print a curve figure
func TestPrintCurveFigure(t *testing.T) {
	// 创建一个小素数域上的椭圆曲线用于可视化 (y^2 = x^3 + 0x + 7 mod 17)
	p := big.NewInt(17)
	a := big.NewInt(0)
	b := big.NewInt(7)

	curve, err := NewEllipticCurve(a, b, p)
	if err != nil {
		t.Fatalf("Failed to create elliptic curve: %v", err)
	}

	t.Log("椭圆曲线: y^2 = x^3 +", a, "x +", b, "(mod", p, ")")
	t.Log("正在生成曲线图形...")

	// 计算曲线上的所有点
	points := collectCurvePoints(curve)

	// 使用gonum/plot创建散点图
	err = plotCurve(curve, points, t)
	if err != nil {
		t.Fatalf("Failed to plot curve: %v", err)
	}

	t.Log("曲线图形已保存为 curve_plot.png")
}

// 新增函数：收集曲线上的所有点
func collectCurvePoints(curve *EllipticCurve) plotter.XYs {
	var points plotter.XYs
	p := curve.P
	maxX := new(big.Int).Sub(p, big.NewInt(1))

	for x := big.NewInt(-5); x.Cmp(maxX) <= 0; x.Add(x, big.NewInt(1)) {
		// 计算 y² = x³ + ax + b mod p
		ySquared := new(big.Int).Exp(x, big.NewInt(3), p) // x³ mod p
		ax := new(big.Int).Mul(curve.A, x)                // ax
		ax.Mod(ax, p)                                     // ax mod p
		ySquared.Add(ySquared, ax)                        // x³ + ax
		ySquared.Add(ySquared, curve.B)                   // x³ + ax + b
		ySquared.Mod(ySquared, p)                         // mod p

		// 检查是否为二次剩余并计算y值
		y := new(big.Int).ModSqrt(ySquared, p)
		if y != nil {
			// 添加(x, y)和(x, p-y)两个点
			points = append(points, plotter.XY{
				X: float64(x.Int64()),
				Y: float64(y.Int64()),
			})
			if y.Cmp(big.NewInt(0)) != 0 { // 避免重复添加原点
				pMinusY := new(big.Int).Sub(p, y)
				points = append(points, plotter.XY{
					X: float64(x.Int64()),
					Y: float64(pMinusY.Int64()),
				})
			}
		}
	}
	return points
}

// 新增函数：使用gonum/plot绘制曲线
func plotCurve(curve *EllipticCurve, points plotter.XYs, t *testing.T) error {
	// 创建新图表
	p := plot.New()

	p.Title.Text = fmt.Sprintf("椭圆曲线 y² = x³ + %dx + %d (mod %d)",
		curve.A.Int64(), curve.B.Int64(), curve.P.Int64())
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	// 创建散点图
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		return err
	}
	scatter.GlyphStyle.Radius = vg.Points(3) // 设置点大小

	p.Add(scatter)

	// 保存为PNG文件
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "curve_plot.png"); err != nil {
		return err
	}
	return nil
}
