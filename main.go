package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var font rl.Font

type data struct {
	arr []int
	m   int
}

func maximum(arr []int, chOut chan data) int {
	m := -1
	for _, n := range arr {
		if n > m {
			m = n
		}
		chOut <- data{
			arr: arr,
			m:   m,
		}
	}
	return m
}

func vizVar(val int) {
	var boxSize float32 = 100
	thick := boxSize / 10
	rec := rl.Rectangle{
		X:      0,
		Y:      200,
		Width:  boxSize,
		Height: boxSize,
	}
	rl.DrawRectangleLinesEx(rec, thick, rl.RayWhite)
	str := fmt.Sprintf("%d", val)
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
func vizArray(arr []int) {
	var boxSize float32 = 100
	thick := boxSize / 10
	pad := -thick
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
}

func main() {
	rl.InitWindow(800, 600, "Raylib")
	defer rl.CloseWindow()
	font = rl.LoadFontEx("./iosevka.ttf", 56, nil, 0)

	rl.SetTargetFPS(60)

	arr := []int{5, 2, 8, 7, 9}
	temparr := arr
	tempm := -1
	chOut := make(chan data)
	timestart := rl.GetTime()
	go maximum(arr, chOut)
	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyN) || rl.GetTime()-timestart > 2 {
			fmt.Println("pressed next")
			timestart = rl.GetTime()
			temp := <-chOut
			temparr = temp.arr
			tempm = temp.m
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0x181818ff))
		vizArray(temparr)
		vizVar(tempm)
		rl.EndDrawing()
	}
}
