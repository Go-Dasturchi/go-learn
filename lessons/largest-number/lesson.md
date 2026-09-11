# Largest Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [10,2]
Output: "210"
```

## TASK

Sizga manfiy bo'lmagan butun sonlardan iborat `nums` massivi berilgan. Ularni tartiblashtiring va eng katta raqamni hosil qiling, natijani satr sifatida qaytaring.

Natija juda katta bo'lgani sababli uni satr sifatida qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Sonlarni oddiy raqamli tartibda solishtirish yetarli emas — ularni satrga aylantirib, ikkita satrni qo'shib qo'yish natijasiga qarab maxsus solishtirish qoidasi kerakligini o'ylab ko'ring.
2. Har bir sonni satrga aylantiring va ikkita a, b satrni solishtirishda `a+b` bilan `b+a` ni taqqoslab, kattasi oldinda turadigan tartibda saralang; saralangandan so'ng ularni birlashtiring, faqat natija "0000" kabi bo'lib qolsa, alohida "0" qaytarishni unutmang.
