#include <raylib.h>
#include <math.h>
#include <stdio.h>

Font iosevka;

int maximum(int* arr, int n) {
    int m = -1;
    for (int i = 0; i < n; i++) {
        if (arr[i] > m) {
            m = arr[i];
        }
    }
    return m;
}

void vizArray(int* arr, int n) {
    int boxSize = 100;
    int thick = 10;
    int pad = -thick;
        for (int i = 0; i < n; i++) {
            Rectangle rec = {
                .x = (boxSize+pad)*i,
                .y = 0,
                .width = boxSize,
                .height = boxSize,
            };
            DrawRectangleLinesEx(rec, 10, GetColor(0xbbbbbbff));
            const char* str =  TextFormat("%d", arr[i]);
            Vector2 fontMeasure = MeasureTextEx(iosevka, str, 58, 0);
            DrawTextPro(iosevka, str, (Vector2){rec.x+(boxSize/2),rec.y+(boxSize/2)}, (Vector2){fontMeasure.x/2,fontMeasure.y/2}, 0, 58, 0, GREEN);
        }
}

int main() {
    int w = 800;
    int h = 600;
    int rw = 100;
    int rh = 100;
    float t = GetTime();
    int arr[] = {2, 5, 3, 2, 4};
    int n = 5;


    InitWindow(w, h, "viz");
    SetTargetFPS(60);

    iosevka = LoadFont("./iosevka.ttf");

    while (!WindowShouldClose()) {
        BeginDrawing();
            ClearBackground(GetColor(0x181818ff));
            //DrawRectangle((w-rw)*((sinf(t)+1.0f)*0.5f), 0, rw, rh, BLUE);
            vizArray(arr, n);
        EndDrawing();
        t = GetTime();
    }
    CloseWindow();
    return 0;
}
