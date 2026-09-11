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

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
