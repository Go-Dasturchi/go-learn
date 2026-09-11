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

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
