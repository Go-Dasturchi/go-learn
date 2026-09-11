# Maximum Number of Vowels in a Substring of Given Length

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: s = "abciiidef", k = 3
Output: 3
Tushuntirish: "iii" qism-satri 3 ta unli harfdan iborat.
```

## TASK

Sizga kichik ingliz harflaridan iborat `s` satri va butun son `k` berilgan. `k` uzunlikdagi har qanday qism-satrida mavjud bo'lishi mumkin bo'lgan maksimal unli harflar sonini qaytaring.

Unli harflar: `'a'`, `'e'`, `'i'`, `'o'`, `'u'`.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, k int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
