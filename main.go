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
	player *Player
}

func NewGame() *Game {
	g := &Game{}
	g.blocks = generateInitialBlocks()
	g.player = NewPlayer()

	return g
}

func (g *Game) Update() error {
	// プレイヤーの移動
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.x -= g.player.speed
		if g.player.x < 0 {
			g.player.x = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.x += g.player.speed
		if g.player.x > gameScreenWidth-playerWidth {
			g.player.x = gameScreenWidth - playerWidth
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// プレイヤーの描画
	var playerOpts ebiten.DrawImageOptions
	playerOpts.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.img, &playerOpts)

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
