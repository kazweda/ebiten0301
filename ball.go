package main

import (
	"math"
)

// ボールの設定
const (
	ballRadius   = 5                    // ボールの半径
	ballSpeed    = 5                    // ボールが移動するスピード
	ballInitialX = gameScreenWidth / 2  // ボールの初期位置
	ballInitialY = gameScreenHeight / 2 // ボールの初期位置
)

type Ball struct {
	x, y           float64
	radius         float64
	speedX, speedY float64
}

func NewBall() *Ball {
	return &Ball{
		x:      ballInitialX,
		y:      ballInitialY,
		radius: ballRadius,
		// 右下斜め45度に向かって、速さ ballSpeed で進む
		speedX: ballSpeed * math.Cos(math.Pi/4),
		speedY: ballSpeed * math.Sin(math.Pi/4),
	}
}
