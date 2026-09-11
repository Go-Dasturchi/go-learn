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

1. O(n) vaqt va O(1) qo'shimcha xotira bilan yechish uchun berilgan massivning o'zini xesh-jadval o'rnida ishlatishni o'ylab ko'ring — 1 dan n gacha bo'lgan har bir musbat son x uchun uning "to'g'ri" joyi x-1 indeksi hisoblanadi.
2. Massiv bo'ylab yurib, agar nums[i] 1 dan n oralig'ida bo'lsa va hali o'z to'g'ri joyida (nums[nums[i]-1]) turmagan bo'lsa, nums[i] bilan nums[nums[i]-1] ni almashtiring va shu tekshiruvni shu indeksda takrorlayvering. Shundan so'ng yana bir marta yurib, nums[i] != i+1 bo'lgan birinchi indeksni toping — javob i+1; agar barchasi joyida bo'lsa javob n+1.
