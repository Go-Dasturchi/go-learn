# Shortest Palindrome

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: s = "aacecaaa"
Output: "aaacecaaa"
Tushuntirish: O'zi "aacecaa" qismi palindrom, shuning uchun boshiga bitta "a" qo'shish kifoya.
```

## TASK

Sizga `s` satri berilgan. Satrning faqatgina boshiga istalgancha belgi qo'shish orqali uni palindromga (boshidan ham oxiridan ham bir xil o'qiladigan satr) aylantirishingiz mumkin.

Ushbu amallarni bajarish orqali hosil bo'lishi mumkin bo'lgan **eng qisqa** palindromni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Asosiy savol shu: "s ning boshidan boshlanadigan ENG UZUN palindrom prefiksi qancha?" — shuni tez topish uchun KMP algoritmining "failure function" (prefiks funksiyasi) g'oyasidan foydalanish mumkin.
2. `s` ni teskarisi bilan maxsus ajratuvchi belgi orqali birlashtiring: combined = s + "#" + reverse(s). Shu birlashtirilgan satr uchun KMP prefiks funksiyasini hisoblang; uning oxirgi qiymati — s ning boshidan boshlanuvchi eng uzun palindrom prefiksi uzunligini beradi. Qolgan (palindrom bo'lmagan) qismning teskarisini s ning boshiga qo'shib javobni hosil qiling.
