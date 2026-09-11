# Longest palindromic substring

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: s = "babad"
Output: "bab"
Tushuntirish: "aba" ham to'g'ri javob hisoblanadi.
```

## TASK

Sizga `s` satri berilgan. `s` ichidagi eng uzun palindrom qism-satrni (substring) toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir palindromning "markazi" bo'lishi mumkinligini o'ylab ko'ring — markazdan ikki tomonga kengayib borish (expand around center) orqali har bir mumkin bo'lgan palindromni tekshirish mumkin.
2. Har bir indeks i uchun ikki xil markazni tekshiring: toq uzunlikdagi palindromlar uchun (i,i) va juft uzunlikdagilar uchun (i,i+1); har ikkalasidan chap va o'ng tomonga bir xil belgi turgunicha kengaytirib boring va shu jarayonda topilgan eng uzun palindromning boshlanish/tugash indekslarini saqlab qo'ying.
