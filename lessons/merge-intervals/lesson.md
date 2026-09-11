# Merge Intervals

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Tushuntirish: [1,3] va [2,6] oraliqlar o'zaro kesishganligi sababli ular [1,6] ga birlashtirildi.
```

## TASK

Sizga `intervals` massivi berilgan bo'lib, uning har bir elementi `intervals[i] = [start_i, end_i]` ko'rinishida oraliqni bildiradi. Barcha o'zaro kesishadigan oraliqlarni birlashtiring va faqatgina kesishmaydigan oraliqlardan tashkil topgan massivni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(intervals [][]int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Kesishuvni tekshirish uchun oraliqlarni qandaydir tartibga solish foydali bo'ladi — boshlanish nuqtasi bo'yicha saralashni o'ylab ko'ring, shunda kesishuvchi oraliqlar bir-biriga yaqin joylashadi.
2. Oraliqlarni boshlanish qiymati bo'yicha o'sish tartibida saralang, so'ng natija ro'yxatini yuritib boring: agar joriy oraliqning boshlanishi oxirgi qo'shilgan oraliqning tugashidan katta bo'lmasa (kesishadi), ularning tugash qiymatlarining kattasini olib birlashtiring, aks holda joriy oraliqni yangi element sifatida ro'yxatga qo'shing.
