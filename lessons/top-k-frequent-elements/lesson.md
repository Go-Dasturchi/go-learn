# Top K Frequent Elements

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

## TASK

Sizga butun sonlardan iborat `nums` massivi va `k` soni berilgan. Ushbu massivda eng ko'p takrorlanadigan `k` ta elementni qaytaring. Natijani istalgan tartibda qaytarish mumkin.

Yechimingiz `O(n log n)` dan tezroq ishlashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, k int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
