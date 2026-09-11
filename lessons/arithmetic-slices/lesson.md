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

1. Har bir indeksda "aynan shu elementda tugaydigan arifmetik ketma-ketliklar soni" degan holatni saqlab boruvchi dinamik dasturlash (DP) haqida o'ylab ko'ring.
2. i=2 dan boshlab yuring: agar nums[i]-nums[i-1] == nums[i-1]-nums[i-2] bo'lsa, joriy hisoblagichni 1 ga oshirib uni umumiy javobga qo'shing; shart bajarilmasa hisoblagichni nolga qaytaring.
