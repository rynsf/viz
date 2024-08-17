package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var font rl.Font

func maximum(arr []int, chOut chan []int) int {
	m := 1
	for i := range arr {
		//if n > m {
		//	m = n
		//}
		chOut <- arr
		arr[i] = m
	}
	return m
}

func vizArray(arr []int) {
	var boxSize float32 = 100
	thick := boxSize / 10
	pad := -thick
	rl.BeginDrawing()
	for n := range arr {
		rec := rl.Rectangle{
			X:      (boxSize + pad) * float32(n),
			Y:      0,
			Width:  boxSize,
			Height: boxSize,
		}
		rl.DrawRectangleLinesEx(rec, thick, rl.RayWhite)
		str := fmt.Sprintf("%d", arr[n])
		fontMeasure := rl.MeasureTextEx(font, str, 56, 0)
		rl.DrawTextPro(font,
			str,
			rl.Vector2{X: rec.X + (boxSize / 2), Y: rec.Y + (boxSize / 2)},
			rl.Vector2{X: fontMeasure.X / 2, Y: fontMeasure.Y / 2},
			0,
			56,
			0,
			rl.Green)
	}
	rl.EndDrawing()
}

func main() {
	rl.InitWindow(800, 600, "Raylib")
	defer rl.CloseWindow()
	font = rl.LoadFontEx("./iosevka.ttf", 56, nil, 0)

	rl.SetTargetFPS(60)

	arr := []int{5, 2, 8, 7, 9}
	temparr := arr
	chOut := make(chan []int)
	timestart := rl.GetTime()
	go maximum(arr, chOut)
	for !rl.WindowShouldClose() {
		if rl.IsKeyReleased(rl.KeyN) || rl.GetTime()-timestart > 2 {
			fmt.Println("pressed next")
			timestart = rl.GetTime()
			temparr = <-chOut
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0x181818ff))
		vizArray(temparr)
		rl.EndDrawing()
	}
}
