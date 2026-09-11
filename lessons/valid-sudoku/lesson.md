# Valid Sudoku

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
```

## TASK

Sizga `9 x 9` o'lchamdagi qisman to'ldirilgan Sudoku doskasi berilgan. Quyidagi qoidalarga ko'ra uning to'g'riligini (valid ekanligini) aniqlang:

1. Har bir qatorda `1-9` gacha raqamlar faqat bir marta ishtirok etishi kerak.
2. Har bir ustunda `1-9` gacha raqamlar faqat bir marta ishtirok etishi kerak.
3. Doskadagi har bir `3 x 3` maydonchada `1-9` gacha raqamlar faqat bir marta ishtirok etishi kerak.

Bo'sh kataklar `'.'` belgisi bilan ifodalanadi.

Izoh: Doska faqatgina hozirgi holatiga ko'ra to'g'ri (valid) bo'lishi yetarli, uning to'liq yechilishi shart emas.

Quyidagi funksiyani to'ldiring:

```go
func solve(board [][]string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir katak bir vaqtning o'zida uchta guruhga (o'z qatori, o'z ustuni, o'z 3x3 bloki) tegishli — shuning uchun har bir guruh uchun alohida "ko'rilgan raqamlar" to'plamini saqlashni o'ylab ko'ring.
2. Har bir qator, ustun va 3x3 blok uchun alohida set (masalan map yoki bool massiv) yuriting; doskani bir marta aylanib chiqing va har bir to'ldirilgan katak uchun, agar uning qiymati mos qator/ustun/blok setida allaqachon bo'lsa false qaytaring, aks holda uni shu uchala setga ham qo'shing (blok indeksini `(r/3)*3 + c/3` orqali hisoblang).
