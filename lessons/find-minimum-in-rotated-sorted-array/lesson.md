# Find Minimum in Rotated Sorted Array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums = [3,4,5,1,2]
Output: 1
Tushuntirish: Dastlabki massiv [1,2,3,4,5] bo'lib u 3 marta aylantirilgan.
```

## TASK

Faraz qiling, dastlab o'sish tartibida joylashgan `nums` massivi oldindan ma'lum bo'lmagan biror indeks atrofida aylantirilgan (rotated). Masalan, `[0,1,2,4,5,6,7]` massivi aylantirilib `[4,5,6,7,0,1,2]` ko'rinishiga kelishi mumkin.

Sizga shunday aylantirilgan `nums` massivi berilgan (takrorlanuvchi elementlarsiz). Undagi eng kichik elementni `O(log n)` vaqt murakkabligida toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
