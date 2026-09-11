# Count Asterisks

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: s = "l|*e*et|c**o|*de|"
Output: 2
Tushuntirish: |...| orasidagi qismlarni hisobga olmasak: "l", "c**o" qoladi. Ularning ichida 2 ta yulduzcha bor.
```

## TASK

Sizga ingliz harflari, bo'sh joylar, yulduzchalar `*` va vertikal chiziqlar `|` dan iborat `s` satri berilgan.

`s` da vertikal chiziqlar soni har doim juft bo'lishi kafolatlanadi.
Har bir ketma-ket ikki vertikal chiziqlar `|` orasidagi belgilarni bir guruh (pair) deb hisoblasak, siz **ushbu juftliklar orasiga kirmaydigan** yulduzchalar `*` sonini topishingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. '|' belgisiga duch kelganingizda holatni (juftlik ichidami yoki yo'qmi) almashtirib borishni o'ylang (boolean flag).
2. Satrni chapdan o'ngga aylanib, '|' ga duch kelganda flag'ni teskarisiga o'zgartiring, '*' ga duch kelganda flag juftlik tashqarisini bildirsa hisoblagichni oshiring.
