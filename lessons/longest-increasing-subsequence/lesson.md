# Longest Increasing Subsequence

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Tushuntirish: Eng uzun o'suvchi qism-ketma-ketlik [2,3,7,101] bo'lib, uzunligi 4.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Uning eng uzun qat'iy o'suvchi qism-ketma-ketligining (strictly increasing subsequence) uzunligini toping.

_Qism-ketma-ketlik — bu massivning ba'zi elementlarini olib tashlash yoki olmaslik orqali hosil bo'lgan ketma-ketlik, bunda qolgan elementlarning tartiblanishi o'zgarmaydi._

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
