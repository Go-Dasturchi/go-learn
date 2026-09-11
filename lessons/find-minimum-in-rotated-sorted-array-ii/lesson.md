# Find Minimum in Rotated Sorted Array II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **35 daqiqa**

## EXAMPLE

```
Input: nums = [2,2,2,0,1]
Output: 0
```

## TASK

Faraz qiling, dastlab o'sish tartibida joylashgan `nums` massivi (takrorlanuvchi elementlari bo'lishi mumkin) oldindan ma'lum bo'lmagan biror indeks atrofida aylantirilgan (rotated). Masalan, `[0,1,2,4,4,4,5,6,6,7]` massivi aylantirilib `[4,5,6,6,7,0,1,2,4,4]` ko'rinishiga kelishi mumkin.

Sizga shunday aylantirilgan `nums` massivi berilgan. Undagi eng kichik elementni qaytaring.

Yechimingiz iloji boricha o'rtacha holatda `O(log n)` vaqt murakkabligida bo'lishiga harakat qiling, ammo eng yomon holatda `O(n)` bo'lishi mumkinligini inobatga oling.

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
