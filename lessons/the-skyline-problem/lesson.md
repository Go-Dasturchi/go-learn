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

1. Har bir binoni ikkita "hodisa" (event) sifatida ifodalang — chap chetida binoning boshlanishini, o'ng chetida esa tugashini — va bu hodisalarni x koordinatasi bo'yicha sweep-line (chapdan o'ngga yurish) usulida qayta ishlang.
2. Hodisalarni x bo'yicha saralang (bir xil x da avval barcha "boshlanish"larni, keyin "tugash"larni qayta ishlang, boshlanishlar orasida balandroq bino avval kelsin). Joriy vaqtdagi barcha faol binolar balandligini maksimal-heap (yoki hisoblagichli ko'p to'plam) da saqlang: har bir hodisadan keyin heap tepasidagi (eng katta) balandlik oldingi natijaviy balandlikdan farq qilsa, [x, yangi balandlik] ni natijaga qo'shing.
