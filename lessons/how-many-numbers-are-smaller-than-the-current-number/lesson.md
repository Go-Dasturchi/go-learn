# How Many Numbers Are Smaller Than the Current Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [8,1,2,2,3]
Output: [4,0,1,1,3]
Tushuntirish:
- 8 uchun 4 ta undan kichik son bor: 1, 2, 2, 3.
- 1 uchun undan kichik son yo'q.
- 2 uchun bitta kichik son bor: 1.
- 3 uchun 3 ta kichik son bor: 1, 2, 2.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan.

Massivdagi har bir `nums[i]` elementi uchun, massivda undan qat'iy kichik bo'lgan nechta raqam borligini hisoblang. Ya'ni, barcha `j != i` indekslar uchun `nums[j] < nums[i]` shartini qanoatlantiradigan `j` lar sonini hisoblashingiz kerak.

Natijani massiv ko'rinishida qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
