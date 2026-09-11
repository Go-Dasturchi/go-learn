# Basic Calculator

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: s = "(1+(4+5+2)-3)+(6+8)"
Output: 23
Tushuntirish:
(1+11-3)+14 = 9+14 = 23
```

## TASK

Sizga ifoda yozilgan `s` satri berilgan. Ifoda musbat butun sonlar, `+`, `-`, `(`, `)` va bo'sh joylardan (` `) iborat.
Ushbu ifodani hisoblab, natijasini qaytaring.

Natija 32-bitli butun son chegarasida bo'lishi kafolatlanadi. Tayyor `eval()` kabi funksiyalardan foydalanmang.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ifodani chapdan o'ngga bitta o'tishda o'qib chiqing va joriy natija bilan joriy ishorani alohida o'zgaruvchilarda saqlang; qavslarni to'g'ri hisoblash uchun stek (stack) tuzilmasidan foydalanishni o'ylab ko'ring.
2. Har bir belgini o'qiganda: raqam bo'lsa uni yig'ing, `+`/`-` uchrasa joriy sonni ishora bilan natijaga qo'shib ishorani yangilang. `(` uchraganda joriy natija va ishorani stekka joylab ularni nolga qaytaring; `)` uchraganda ichki natijani hisoblab, stekdan chiqargan qiymatlar yordamida tashqi natijaga qo'shing.
