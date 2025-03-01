package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// プレイヤーの設定
const (
	playerWidth    = 80                                  // プレイヤーの幅
	playerHeight   = 10                                  // プレイヤーの高さ
	initialPlayerX = (gameScreenWidth - playerWidth) / 2 // プレイヤーの初期位置
	initialPlayerY = gameScreenHeight - playerHeight     // プレイヤーの初期位置
	playerSpeed    = 6                                   // プレイヤーのスピード

)

type Player struct {
	x, y   float64
	width  float64
	height float64
	img    *ebiten.Image
	speed  float64
}

func NewPlayer() *Player {
	img := ebiten.NewImage(int(playerWidth), int(playerHeight))
	img.Fill(color.White)

	return &Player{
		x:      initialPlayerX,
		y:      initialPlayerY,
		width:  playerWidth,
		height: playerHeight,
		img:    img,
		speed:  playerSpeed,
	}
}
