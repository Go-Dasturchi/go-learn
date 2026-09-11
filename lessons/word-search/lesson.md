# Word Search

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```

## TASK

Sizga `m x n` o'lchamli belgilar (characters) dan iborat `board` to'ri va bitta `word` so'zi berilgan. Agar o'sha so'z to'rda mavjud bo'lsa `true` qaytaring, aks holda `false`.

So'z harflari gorizontal yoki vertikal bo'ylab bir-biriga qo'shni bo'lgan kataklardagi harflardan tuzilishi mumkin. Bitta katakdagi harfdan bitta so'z ichida faqat bir marta foydalanish mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(board [][]string, word string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
