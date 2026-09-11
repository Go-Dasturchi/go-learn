# The Skyline Problem

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **50 daqiqa**

## EXAMPLE

```
Input: buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
```

## TASK

Sizga shahar binolarining joylashuvi haqida ma'lumot berilgan. Barcha binolar tekis yuzada (x-o'qi) joylashgan va to'rtburchak shaklida.
Sizga `buildings` massivi berilgan bo'lib, har bir `buildings[i] = [lefti, righti, heighti]` ni bildiradi:

- `lefti` — binoning x-o'qidagi chap tomoni.
- `righti` — binoning x-o'qidagi o'ng tomoni.
- `heighti` — binoning balandligi.

Ushbu binolar hosil qilgan shahar siluetining konturini (skyline) toping.
Skyline quyidagicha qaytariladi: asosiy nuqtalar (key points) ketma-ketligi, bunda har bir nuqta `[x, y]` ko'rinishida bo'lib u gorizontal chiziqning boshlanishini ifodalaydi. Eng oxirgi nuqta har doim zaminga qaytganini bildirish uchun `[x, 0]` shaklida bo'ladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(buildings [][]int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
