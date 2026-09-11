# Jewels and Stones

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: jewels = "aA", stones = "aAAbbbb"
Output: 3
Tushuntirish: jewels = 'a' va 'A'. stones = "aAAbbbb" ichida bitta 'a' va ikkita 'A' bor. Jami 3 ta.
```

## TASK

Sizga qimmatbaho toshlar turlarini ifodalovchi `jewels` satri va o'zingizda bor bo'lgan toshlarni ifodalovchi `stones` satri berilgan.
`stones` satridagi har bir belgi sizdagi bir dona toshni ifodalaydi. `jewels` satridagi belgilar esa qimmatbaho toshlarni. Katta va kichik harflar farq qiladi (masalan, `'a'` va `'A'` har xil toshlar hisoblanadi).

Sizdagi toshlar ichida nechtasi qimmatbaho tosh (jewel) ekanligini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(jewels string, stones string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. jewels satridagi belgilarni tez tekshirish uchun set (yoki boolean massiv) tuzib olish qulay.
2. Avval jewels harflarini set'ga joylang, so'ngra stones satridagi har bir belgi shu set'da bor-yo'qligini tekshirib, mos kelganlarni sanang.
