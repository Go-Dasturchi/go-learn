# Remove element

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [3,2,2,3], val = 3
Output: 2
```

## TASK

Sizga `nums` massivi va `val` soni berilgan. Massivdagi barcha `val` ga teng bo'lgan elementlarni in-place o'chiring va massivda qolgan (val ga teng bo'lmagan) elementlar sonini `k` ni qaytaring.

Massiv elementlarining joylashuv tartibini o'zgartirishingiz mumkin. Swift dasturlash tilida odatda siz modifikatsiya o'rniga faqat shu qoladigan qiymatlar sonini hisoblab qaytarishingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, val int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
