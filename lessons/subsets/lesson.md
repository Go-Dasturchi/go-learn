# Subsets

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

## TASK

Sizga o'zaro takrorlanmaydigan (farqli) butun sonlardan iborat `nums` massivi berilgan. Uning barcha mumkin bo'lgan qism-to'plamlarini (subsets, ya'ni power set) qaytaring.

Natijadagi ro'yxatda takrorlanuvchi qism-to'plamlar bo'lishi mumkin emas. Natijani istalgan tartibda qaytarishingiz mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir yangi son qo'shilganda mavjud qism-to'plamlar sonini ikki barobarga oshirish mumkinligini o'ylab ko'ring — yoki bo'lmasa har bir sonni "olish/olmaslik" tanlovi bilan yuruvchi rekursiv qidiruvni ko'rib chiqing.
2. Natijani bo'sh to'plam `[[]]` bilan boshlang; har bir yangi n uchun natijadagi hozirgi barcha qism-to'plamlarning nusxasini olib, ularga n ni qo'shib, shu yangi qism-to'plamlarni ham natijaga qo'shib boring — shunda oxirida barcha 2^n ta kombinatsiya hosil bo'ladi.
