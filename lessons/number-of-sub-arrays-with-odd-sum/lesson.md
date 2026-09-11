# Number of Sub-arrays With Odd Sum

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: arr = [1,3,5]
Output: 4
Tushuntirish: [1], [3], [5] va [1,3,5] — barchasi toq yig'indili qism-massivlar.
```

## TASK

Sizga butun sonlardan iborat `arr` massivi berilgan. Yig'indisi **toq** (odd) bo'lgan barcha qism-massivlarning sonini toping.

Natijani `10^9 + 7` ga bo'lingandagi qoldig'ini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(arr []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
