package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gameScreenWidth  = 640 // ゲームのウィンドウの横幅
	gameScreenHeight = 480 // ゲームのウィンドウの縦幅
)

type Game struct {
	blocks []*Block
}

func NewGame() *Game {
	g := &Game{}
	g.blocks = generateInitialBlocks()

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ブロックの描画
	for _, block := range g.blocks {
		// isVisible == false の Block（ボールが衝突した場合）は表示しない
		if block.isVisible {
			var opts ebiten.DrawImageOptions      // オプションの宣言
			opts.GeoM.Translate(block.x, block.y) // 描画位置を指定
			screen.DrawImage(block.img, &opts)    // 画像を指定したオプションで描画
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameScreenWidth, gameScreenHeight
}

func main() {
	g := NewGame() // ゲームの初期化
	ebiten.SetWindowSize(gameScreenWidth, gameScreenHeight)
	ebiten.SetWindowTitle("ブロック崩し")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
