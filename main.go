package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Hello World")
	rl.InitWindow(800, 600, "Raylib")
	defer rl.CloseWindow()
	//font := rl.LoadFont("./iosevka.ttf")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Red)
		rl.DrawText("Hello World", 0, 0, 32, rl.Black)
		rl.EndDrawing()
	}
}
