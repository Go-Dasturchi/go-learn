# Get Equal Substrings Within Budget

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: s = "abcd", t = "bcdf", maxCost = 3
Output: 3
Tushuntirish: "abc" -> "bcd" uchun xarajat: |a-b| + |b-c| + |c-d| = 1+1+1 = 3. Maksimal uzunlik 3.
```

## TASK

Sizga ikkita kichik ingliz harflaridan tuzilgan `s` va `t` satrlari, hamda butun son `maxCost` berilgan.

Siz `s` ni `t` ga aylantirmoqchisiz. `i`-chi belgini o'zgartirish xarajati `|s[i] - t[i]|` ga teng (belgilarning ASCII kodlari farqi).

Ko'pi bilan `maxCost` sarflagan holda `s` dan `t` ga o'zgartirish mumkin bo'lgan uzluksiz qism-satrning maksimal uzunligini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, t string, maxCost int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
