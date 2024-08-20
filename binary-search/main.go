package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var font rl.Font
var boxSize float32 = 100
var thick = boxSize / 10
var pad = -thick

type data struct {
	arr  []int
	low  int
	mid  int
	high int
}

func binarySearch(arr []int, target int, chOut chan data) int {
	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		chOut <- data{
			arr:  arr,
			low:  low,
			mid:  mid,
			high: high,
		}
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func vizWindow(low, high int, color rl.Color) {
	rec := rl.Rectangle{
		X:      (boxSize + pad) * float32(low),
		Y:      0,
		Width:  (boxSize+pad)*float32((high-low)+1) - pad,
		Height: boxSize,
	}
	rl.DrawRectangleLinesEx(rec, thick, color)
}

func vizVar(val int, pos float32) {
	rec := rl.Rectangle{
		X:      (boxSize + pad) * pos,
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
	rl.InitWindow(1600, 900, "Raylib")
	defer rl.CloseWindow()
	font = rl.LoadFontEx("./iosevka.ttf", 56, nil, 0)

	rl.SetTargetFPS(60)

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	target := 0
	tempData := data{
		arr:  arr,
		low:  -1,
		mid:  -1,
		high: -1,
	}
	chOut := make(chan data)
	timestart := rl.GetTime()
	go binarySearch(arr, target, chOut)
	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyN) || rl.GetTime()-timestart > 2 {
			fmt.Println("pressed next")
			timestart = rl.GetTime()
			tempData = <-chOut
			fmt.Println(tempData)
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0x181818ff))
		vizArray(tempData.arr)
		vizWindow(tempData.low, tempData.high, rl.Red)
		vizWindow(tempData.mid, tempData.mid, rl.Blue)
		rl.EndDrawing()
	}
}
