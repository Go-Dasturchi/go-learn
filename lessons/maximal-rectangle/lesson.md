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

1. Bu masalani "Largest Rectangle in Histogram" masalasiga qaytarib yechish mumkin — har bir qatorni gistogrammaning asosi deb, o'sha qatorgacha har bir ustunda tepasidan qancha ketma-ket `1` borligini balandlik sifatida hisoblang.
2. `heights` massivini qator bo'ylab yurib yangilab boring: agar matrix[qator][ustun]=="1" bo'lsa heights[ustun]++ qiling, aks holda heights[ustun]=0 qiling. Har bir qatordan keyin shu heights massiviga nisbatan "eng katta to'rtburchak gistogrammada" algoritmini (monotonik stek bilan) qo'llab, natijalarning eng kattasini saqlang.
