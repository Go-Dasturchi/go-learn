# Arithmetic Slices

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3,4]
Output: 3
Tushuntirish: [1,2,3], [2,3,4] va [1,2,3,4] uchta arifmetik ketma-ketlik.
```

## TASK

Arifmetik ketma-ketlik (arithmetic sequence) — bu uchta yoki undan ortiq elementdan iborat bo'lib, qo'shni elementlar orasidagi farq doim bir xil bo'ladigan ketma-ketlik.

Sizga `nums` butun sonlar massivi berilgan. Undagi arifmetik ketma-ketlik bo'ladigan qism-massivlar (subarrays) sonini toping.

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
