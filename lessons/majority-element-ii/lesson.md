# Majority Element II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [3,2,3]
Output: [3]
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Massivda `n/3` dan ko'p marta (strictly more than `⌊n/3⌋` times) takrorlanadigan barcha elementlarni toping.

Yechimingiz `O(1)` qo'shimcha xotira bilan ishlashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. n/3 dan ko'p marta uchraydigan elementlar ko'pi bilan ikkita bo'lishi mumkinligini payqang — bu Boyer-Moore Voting algoritmini ikkita nomzod bilan umumlashtirishga olib keladi.
2. Ikkita nomzod (cand1, cand2) va ularning hisoblagichlarini yuriting: joriy son mavjud nomzodlardan biriga teng bo'lsa hisoblagichini oshiring, bo'sh nomzod bo'lsa joriy sonni nomzod qiling, aks holda ikkala hisoblagichni ham kamaytiring; oxirida massivni qayta aylanib, ikkala nomzodning haqiqiy takrorlanish sonini hisoblab, n/3 dan ko'p bo'lganlarinigina natijaga qo'shing.
