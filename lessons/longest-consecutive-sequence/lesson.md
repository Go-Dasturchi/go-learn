# Longest Consecutive Sequence

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **35 daqiqa**

## EXAMPLE

```
Input: nums = [100,4,200,1,3,2]
Output: 4
Tushuntirish: Eng uzun ketma-ketlik bu [1, 2, 3, 4] bo'lib, uning uzunligi 4 ga teng.
```

## TASK

Sizga tartiblanmagan butun sonlardan iborat `nums` massivi berilgan. Uning ichidan eng uzun ketma-ketlik hosil qiladigan elementlar ro'yxatining uzunligini toping.

Sizning yechimingiz `O(n)` vaqt murakkabligida ishlashi shart.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massivni saralashsiz O(n) da yechish uchun barcha qiymatlarni hash-to'plamga (set) joylang va faqat "ketma-ketlikning boshlanish nuqtasi" bo'lgan sonlardan qidiruvni boshlang.
2. Bir son n uchun, agar to'plamda n-1 mavjud bo'lsa, bu n biror ketma-ketlikning o'rtasida joylashgan degani — uni o'tkazib yuboring. Aks holda (n-1 yo'q bo'lsa), n dan boshlab n+1, n+2, ... to'plamda mavjud bo'lguncha uzunlikni sanang; shu tarzda har bir son umumiy hisobda faqat bir marta ko'rib chiqiladi va yechim O(n) bo'ladi.
