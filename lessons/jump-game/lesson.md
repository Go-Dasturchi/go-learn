# Jump Game

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums = [2,3,1,1,4]
Output: true
Tushuntirish: 0-indeksdan 1 qadam sakrab 1-indeksga o'tamiz, u yerdan esa 3 qadam sakrab oxirgi indeksga yetib boramiz.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Siz dastlab massivning birinchi indeksida joylashgansiz va har bir element sizning o'sha joydan maksimal qancha sakrashingiz mumkinligini bildiradi.

Agar oxirgi indeksga yetib bora olsangiz `true` qaytaring, aks holda `false`.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
