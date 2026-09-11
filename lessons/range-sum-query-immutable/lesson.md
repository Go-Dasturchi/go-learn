# Range Sum Query - Immutable

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [-2,0,3,-5,2,-1], left = 0, right = 2
Output: 1
Tushuntirish: sumRange(0, 2) = -2 + 0 + 3 = 1
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. `left` va `right` indekslari orasidagi (ikki tomonlama ham kiritilgan) elementlarning yig'indisini hisoblaydigan funksiya yarating.

Masalan, `sumRange(nums, left, right)` funksiyasi `nums[left] + ... + nums[right]` qiymatini qaytarishi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, left int, right int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu masalada faqat bitta so'rov bo'lgani uchun murakkab prefiks-yig'indi tuzilmasi shart emas — oddiy tsikl yetarli.
2. `left` dan `right` gacha (ikkalasi ham kiritilgan holda) tsikl yurgizib, oraliqdagi barcha elementlarni bitta yig'indiga qo'shib chiqing.
