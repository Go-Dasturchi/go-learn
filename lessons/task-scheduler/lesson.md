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

1. Vaqt jadvalini bevosita simulyatsiya qilish shart emas — eng ko'p takrorlanadigan vazifa qancha "bo'shliq" yaratishini matematik formula orqali hisoblashni o'ylab ko'ring.
2. Har bir vazifaning chastotasini sanang, eng katta chastotani (maxFreq) va shu chastotaga ega vazifalar sonini (maxCount) toping; formula bo'yicha minimal vaqt `(maxFreq-1)*(n+1) + maxCount` ga teng bo'ladi, ammo agar umumiy vazifalar soni bundan katta bo'lsa (bo'shliqqa ehtiyoj yo'q holat), javob shunchaki vazifalar sonining o'ziga teng bo'ladi.
