# Find first and last position

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3, 4]
```

## TASK

O'sish tartibida saralangan `nums` massivi berilgan. `target` sonining birinchi va oxirgi marta qatnashgan indekslarini massiv ko'rinishida `[first, last]` qaytaring. Topilmasa `[-1, -1]` qayting. Algoritm `O(log n)` murakkablikda ishlashi zarur.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, target int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massiv saralangan bo'lgani uchun oddiy binary search kifoya qilmaydi — target ning birinchi va oxirgi o'rnini topish uchun ikkita alohida binary search yurgizish kerakligini o'ylab ko'ring.
2. Bittasi target dan kichik bo'lgan eng chap chegarani (lower bound), ikkinchisi target dan katta bo'lgan eng chap chegarani (upper bound) topadigan ikkita binary search yozing; birinchisi natijaning boshlanish indeksini, ikkinchisidan 1 ayirilgani esa tugash indeksini beradi — agar boshlanish indeksidagi qiymat target ga teng bo'lmasa, `[-1,-1]` qaytaring.
