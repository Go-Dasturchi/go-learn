# Maximal Rectangle

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 6
Tushuntirish: Eng katta to'rtburchak (2,2) dan (3,4) gacha bo'lgan qism bo'lib, yuzasi 6 ga teng.
```

## TASK

Sizga `0` va `1` lardan iborat bo'lgan `rows x cols` o'lchamli 2D binar `matrix` berilgan. Matritsa ichidagi faqat `1` lardan iborat bo'lgan eng katta to'rtburchakni toping va uning yuzini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(matrix [][]string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
