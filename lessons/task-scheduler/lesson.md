# Task Scheduler

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Tushuntirish: A -> B -> idle -> A -> B -> idle -> A -> B
```

## TASK

Sizga CPU vazifalari ro'yxati `tasks` va `n` soni berilgan. `n` — bu bir xil turdagi ikki vazifa orasida bo'lishi zarur minimal pauza (idle cycle) miqdori.

Barcha vazifalarni bajarishga kerak bo'lgan minimal vaqt birliklarini toping. Agar CPU biror vaqt birligida bo'sh tursa, bu ham 1 birlik hisob qilinadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(tasks []string, n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
