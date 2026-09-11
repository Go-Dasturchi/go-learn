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

1. Har bir elementda "shu yerda tugaydigan eng uzun o'suvchi ketma-ketlik uzunligi"ni saqlaydigan oddiy DP bilan boshlash mumkin, lekin tezroq yechim uchun "har bir uzunlikdagi o'suvchi ketma-ketlikning eng kichik oxirgi elementi"ni saqlab borishni o'ylab ko'ring.
2. `tails` deb nomlangan yordamchi massiv yuriting — tails[i] shu uzunlikdagi o'suvchi ketma-ketlikning eng kichik mumkin bo'lgan oxirgi qiymati; har bir yangi son uchun binary search orqali uni tails ichida qayerga qo'yish (yoki almashtirish) mumkinligini toping — agar songa joy bo'lmasa, tails oxiriga qo'shing; oxirida tails uzunligi javob bo'ladi.
