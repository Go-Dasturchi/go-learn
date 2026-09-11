# Bitwise AND of Numbers Range

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: left = 5, right = 7
Output: 4
Tushuntirish: 5 AND 6 AND 7 = 4 (0101 & 0110 & 0111 = 0100)
```

## TASK

Sizga `left` dan `right` gacha bo'lgan oraliqtagi barcha sonlar berilgan (jumladan oraliq chegaralari ham kiritilgan). Ular ustida bitli VA (bitwise AND) amalini bajaring va natijani qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(left int, right int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
