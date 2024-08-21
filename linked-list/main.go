package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var font rl.Font
var boxSize float32 = 100
var thick = boxSize / 10
var pad = -thick
var ArrowLen = 100

const fontSize = 58

type linkedlist struct {
	val  int
	next *linkedlist
}

func vizLinkedList(head *linkedlist) {
	pad = 0
	i := 0
	for head != nil {
		rec := rl.Rectangle{
			X:      (boxSize + pad + float32(ArrowLen)) * float32(i),
			Y:      0,
			Width:  boxSize,
			Height: boxSize,
		}
		rl.DrawRectangleLinesEx(rec, thick, rl.RayWhite)
		str := fmt.Sprintf("%d", head.val)
		fontMeasure := rl.MeasureTextEx(font, str, fontSize, 0)
		rl.DrawTextPro(font,
			str,
			rl.Vector2{X: rec.X + (boxSize / 2), Y: rec.Y + (boxSize / 2)},
			rl.Vector2{X: fontMeasure.X / 2, Y: fontMeasure.Y / 2},
			0,
			fontSize,
			0,
			rl.Green)
		lineStart := rl.Vector2{X: rec.X + rec.Width,
			Y: rec.Y + (rec.Height / 2)}
		lineEnd := rl.Vector2{X: lineStart.X + float32(ArrowLen-10),
			Y: lineStart.Y,
		}
		rl.DrawLineEx(lineStart, lineEnd, thick, rl.RayWhite)
		if head.next == nil {
			rl.DrawLineEx(rl.Vector2{X: lineEnd.X, Y: lineEnd.Y - 20},
				rl.Vector2{X: lineEnd.X, Y: lineEnd.Y + 20},
				thick,
				rl.RayWhite)
			rl.DrawLineEx(rl.Vector2{X: lineEnd.X + 20, Y: lineEnd.Y - 10},
				rl.Vector2{X: lineEnd.X + 20, Y: lineEnd.Y + 10},
				thick,
				rl.RayWhite)
		} else {
			rl.DrawTriangle(
				rl.Vector2{X: lineEnd.X + 10, Y: lineEnd.Y},
				rl.Vector2{X: lineEnd.X - 20, Y: lineEnd.Y - 20},
				rl.Vector2{X: lineEnd.X - 20, Y: lineEnd.Y + 20},
				rl.RayWhite)
		}
		head = head.next
		i += 1
	}
}

func main() {
	head := linkedlist{
		val:  1,
		next: nil,
	}
	node1 := linkedlist{
		val:  2,
		next: nil,
	}
	head.next = &node1
	node2 := linkedlist{
		val:  3,
		next: nil,
	}
	node1.next = &node2

	rl.InitWindow(1600, 900, "Raylib")
	defer rl.CloseWindow()
	font = rl.LoadFontEx("./iosevka.ttf", fontSize, nil, 0)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0x181818ff))
		vizLinkedList(&head)
		rl.EndDrawing()
	}
}
