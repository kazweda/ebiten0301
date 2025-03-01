package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	blockRowNums    = 5  // ブロックを何行ならべるか
	blockCloumnNums = 10 // ブロックを何列ならべるか
	blockPadding    = 10 // ブロックとブロックの間の間隔
	blockTopOffset  = 50 // 画面の上端とブロックの間の間隔
	blockWidth      = 60 // ブロックの幅
	blockHeight     = 20 // ブロックの高さ
)

type Block struct {
	x, y      float64       // ブロックの描画位置 (x, y)。左上が (0, 0) になる
	width     float64       // ブロックの幅
	height    float64       // ブロックの高さ
	isVisible bool          // ブロックが見えるかどうか。ボールが衝突したら false になる
	img       *ebiten.Image // ブロックの画像
}

// ゲーム開始した時のブロックを生成する
func generateInitialBlocks() []*Block {
	var blocks []*Block
	// 指定した行、列の分だけ Block 構造体を作成してスライスに追加する
	for row := 0; row < blockRowNums; row++ {
		for col := 0; col < blockCloumnNums; col++ {
			color := color.RGBA{
				R: uint8(200 - row*30), // 何行目かで色を少し変える
				G: uint8(200 - row*30), // 何行目かで色を少し変える
				B: 255,
				A: 255,
			}

			// ブロックの画像を生成
			img := ebiten.NewImage(blockWidth, blockHeight)
			img.Fill(color)

			block := &Block{
				x:         float64(col)*(blockWidth+blockPadding) + blockPadding,    // 左端からの距離
				y:         float64(row)*(blockHeight+blockPadding) + blockTopOffset, // 上端からの距離
				width:     blockWidth,
				height:    blockHeight,
				isVisible: true, // 初期化時は true にしてみえるようにする
				img:       img,
			}
			blocks = append(blocks, block) // スライスに追加
		}
	}

	return blocks
}
