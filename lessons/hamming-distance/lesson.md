# Hamming Distance

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: x = 1, y = 4
Output: 2
Tushuntirish:
1  = 0 0 0 1
4  = 0 1 0 0
          ^  ^   - ikkita bit farq qiladi
```

## TASK

Ikkita butun son `x` va `y` o'rtasidagi **Hamming masofasi** ularning ikkilik (binary) ko'rinishlarida farq qiluvchi bitlar sonidir.

Sizga `x` va `y` berilgan, Hamming masofasini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(x int, y int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sonning bitlab farqini topish uchun bitwise XOR amalini eslang — XOR natijasida farqli bitlar 1 bo'ladi.
2. x va y'ni XOR qiling, so'ng natijadagi 1 bitlar sonini hisoblang (masalan math/bits paketidagi tayyor funksiya yordamida).
