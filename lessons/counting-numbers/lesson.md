# Counting Numbers

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: low = 3, high = 7
Output: 12
Tushuntirish: 3 dan 7 gacha juft sonlar: 4 va 6. 4 + 6 = 10 emas, balki 4, 6 = 10. Lekin to'g'ri javob 12.
```

_To'g'ri misol:_

```
Input: low = 1, high = 10
Output: 30
Tushuntirish: 2 + 4 + 6 + 8 + 10 = 30
```

## TASK

Sizga ikkita butun son `low` va `high` berilgan. `[low, high]` oralig'idagi barcha juft sonlar (even numbers) yig'indisini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(low int, high int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
