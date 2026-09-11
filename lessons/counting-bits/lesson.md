# Counting Bits

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 2
Output: [0,1,1]
Tushuntirish:
0 -> 0b0 -> 0 ta '1'
1 -> 0b1 -> 1 ta '1'
2 -> 0b10 -> 1 ta '1'
```

## TASK

Sizga butun son `n` berilgan. Uzunligi `n + 1` bo'lgan `ans` massivini qaytaring, bunda `ans[i]` qiymati `i` sonining ikkilik (binary) ko'rinishidagi `1` larning sonidir.

Yechimingiz `O(n)` vaqt murakkabligida va `O(n)` dan kam qo'shimcha xotirada ishlashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
