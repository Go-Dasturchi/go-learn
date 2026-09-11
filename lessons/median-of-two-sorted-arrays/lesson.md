# Median of two sorted arrays

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Tushuntirish: Birlashgan massiv = [1,2,3] va uning mediani 2.
```

## TASK

O'lchamlari mos ravishda `m` va `n` bo'lgan ikkita saralangan `nums1` va `nums2` massivlari berilgan. Ularning umumiy medianini toping.

Umumiy vaqt murakkabligi `O(log(m+n))` bo'lishi talab qilinadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums1 []int, nums2 []int) float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
