# Generate parentheses

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

## TASK

Sizga `n` soni berilgan, u ochiq-yopiq qavslar juftligining sonini bildiradi. Hammasi to'g'ri yopilgan bo'lishi sharti bilan yasalishi mumkin bo'lgan barcha kombinatsiyalarni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir qadamda ochiq yoki yopiq qavs qo'shish mumkinmi degan tanlov bilan ishlaydigan rekursiv qidiruv (backtracking) haqida o'ylab ko'ring, faqat har doim to'g'ri (valid) natijaga olib boradigan tanlovlarni qiling.
2. Joriy satr, ishlatilgan ochiq qavslar soni va ishlatilgan yopiq qavslar sonini parametr sifatida uzatuvchi rekursiv funksiya yozing: ochiq qavslar soni n dan kichik bo'lsa "(" qo'shishingiz mumkin, yopiq qavslar soni ochiq qavslardan kam bo'lsa ")" qo'shishingiz mumkin; satr uzunligi 2n ga yetganda uni natijaga saqlang.
