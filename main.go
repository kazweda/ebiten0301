package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	gameScreenWidth  = 640 // ゲームのウィンドウの横幅
	gameScreenHeight = 480 // ゲームのウィンドウの縦幅
)

type Game struct {
	blocks []*Block
	player *Player
	ball   *Ball
}

func NewGame() *Game {
	g := &Game{}
	g.initialize()

	return g
}

func (g *Game) initialize() {
	g.blocks = generateInitialBlocks()
	g.player = NewPlayer()
	g.ball = NewBall()
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

	// ボールの移動
	g.ball.x += g.ball.speedX
	g.ball.y += g.ball.speedY

	// ボールと壁との衝突判定
	// 左の壁：g.ball.x < g.ball.radius
	// 右の壁：g.ball.x > gameScreenWidth-g.ball.radius
	if g.ball.x < g.ball.radius || g.ball.x > gameScreenWidth-g.ball.radius {
		g.ball.speedX *= -1
	}
	// 上の壁：g.ball.y < g.ball.radius
	if g.ball.y < g.ball.radius {
		g.ball.speedY *= -1
	}
	// 下の壁はゲームオーバー
	if g.ball.y > gameScreenHeight-g.ball.radius {
		g.initialize() // ゲームオーバーになったら、ゲームを初期化する。
	}

	// プレイヤーとボールの衝突判定
	if (g.ball.y+g.ball.radius >= g.player.y) && (g.ball.y+g.ball.radius <= g.player.y+g.player.height) {
		if g.ball.x >= g.player.x && g.ball.x <= g.player.x+g.player.width {
			// ボールの速度を更新
			relativeIntersectX := (g.ball.x - (g.player.x + g.player.width/2)) / (g.player.width / 2)
			bounceAngle := relativeIntersectX * (math.Pi / 3) // 最大60度の角度で反射

			g.ball.speedX = ballSpeed * math.Sin(bounceAngle)
			g.ball.speedY = -ballSpeed * math.Cos(bounceAngle)

			// ボールをプレイヤーの上に位置させる
			g.ball.y = g.player.y - g.ball.radius - 1
		}
	}

	// ブロックとボールの衝突判定
	for _, block := range g.blocks {
		if block.isVisible {
			if g.ball.x >= block.x && g.ball.x <= block.x+block.width &&
				g.ball.y-g.ball.radius <= block.y+block.height && g.ball.y+g.ball.radius >= block.y {
				g.ball.speedY *= -1
				block.isVisible = false
				break
			}
		}
	}

	// // speedX と speedY をターミナルに出力
	// log.Printf("Ball speedX: %f, speedY: %f", g.ball.speedX, g.ball.speedY)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// プレイヤーの描画
	var playerOpts ebiten.DrawImageOptions
	playerOpts.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.img, &playerOpts)

	// ボールの描画
	DrawBall(screen, g.ball, color.White)

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

// 指定した座標にボールを描画する。円に近くなるように多角形を描いている。
func DrawBall(screen *ebiten.Image, ball *Ball, clr color.Color) {
	segments := 72
	for i := 0; i < segments; i++ {
		angle1 := float64(i) / float64(segments) * 2 * math.Pi
		angle2 := float64(i+1) / float64(segments) * 2 * math.Pi
		x1 := ball.x + ball.radius*math.Cos(angle1)
		y1 := ball.y + ball.radius*math.Sin(angle1)
		x2 := ball.x + ball.radius*math.Cos(angle2)
		y2 := ball.y + ball.radius*math.Sin(angle2)
		vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1, clr, false)
	}
}
