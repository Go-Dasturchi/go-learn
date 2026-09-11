# Word Search II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
Tushuntirish: Jadvalda "oath" va "eat" so'zlarini gorizontal/vertikal yo'llar bilan topish mumkin. "pea" va "rain" ni esa to'liq topib bo'lmaydi.
```

## TASK

Sizga `m x n` o'lchamli belgilar jadvali (board) va satrlar ro'yxati (words) berilgan. Jadval ichidan `words` ro'yxatiga tegishli bo'lgan barcha so'zlarni toping.

Har bir so'z ketma-ket yopishgan harflardan yig'ilishi kerak. Yopishgan (qo'shni) harflar - bu gorizontal yoki vertikal joylashgan harflar. Bitta so'zni yig'ishda bir hil xujayrani (cell) qayta ishlatish mumkin emas.

Quyidagi funksiyani to'ldiring:

```go
func solve(board [][]string, words []string) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
