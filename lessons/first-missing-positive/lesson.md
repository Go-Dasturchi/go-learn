# First missing positive

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,0]
Output: 3
Tushuntirish: 1 va 2 bor, demak eng kichik tushib qolgan son 3.
```

## TASK

Saralanmagan butun sonlar massivi `nums` berilgan. Unda ishtirok etmagan eng kichik musbat (1, 2, ...) butun sonni toping.

Siz yozgan algoritm `O(n)` vaqtda va iloji boricha doimiy `O(1)` xotira hajmi bilan ishlashi zarur.

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
