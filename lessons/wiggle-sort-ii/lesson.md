# Wiggle Sort II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: nums = [1,5,1,1,6,4]
Output: [1,6,1,5,1,4]
Tushuntirish: [1,4,1,5,1,6] ham to'g'ri javob hisoblanadi.
```

## TASK

Sizga o'sish tartibida tartiblanmagan `nums` butun sonlar massivi berilgan. Uni `nums[0] < nums[1] > nums[2] < nums[3]...` shart bajarilishi uchun in-place usulda qayta joylashtiring.

_Eslatma: Bu yerda tizim testining maqsadida, siz o'zgartirilgan massivni qaytaring (in-place talab qilinmaydi)._

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massivni saralab, uni kichik va katta yarimlarga bo'lib, keyin ularni juft va toq pozitsiyalarga navbat bilan taqsimlashni o'ylab ko'ring — shunda katta sonlar kichiklarning orasiga tushib, "to'lqin" shakli hosil bo'ladi.
2. Massivni saralang, uni ikki qismga bo'ling (kichikroq yarim va kattaroq yarim); kichikroq yarimning elementlarini teskari tartibda juft indekslarga (0,2,4,...), kattaroq yarimning elementlarini ham teskari tartibda toq indekslarga (1,3,5,...) joylashtiring — teskari tartib teng qiymatlarning yonma-yon tushib qolishining oldini oladi.
