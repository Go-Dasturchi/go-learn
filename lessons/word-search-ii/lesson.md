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

1. Har bir so'zni jadvalda alohida-alohida qidirish (har biri uchun to'liq DFS) samarasiz bo'ladi — barcha so'zlarni bitta Trie (prefiks daraxti) ga yig'ib, jadval bo'ylab bitta umumiy DFS bilan bir yo'la qidiring.
2. `words` dan Trie quring, har bir tugunda "shu yergacha to'liq so'z tugaydimi" belgisini saqlang. Jadvaldagi har bir katakdan DFS boshlang: joriy harf Trie da mos yo'l bo'lsagina davom eting (aks holda darhol qaytish — bu keraksiz qidiruvlarni katta miqdorda kesib tashlaydi), agar joriy Trie tugunida to'liq so'z tugasa uni natijalar to'plamiga qo'shing. Qayta ishlatishdan saqlanish uchun tashrif buyurilgan katakni vaqtincha belgilab qo'ying va DFS dan qaytgach asl holatiga tiklang.
