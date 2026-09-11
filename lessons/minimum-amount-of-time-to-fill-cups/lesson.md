# Minimum Amount of Time to Fill Cups

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: amount = [1,4,2]
Output: 4
Tushuntirish:
1-sekund: Iliq va issiq stakanga quyish. amount = [1,3,1]
2-sekund: Sovuq va iliq stakanga quyish. amount = [0,2,1]
3-sekund: Iliq va issiq stakanga quyish. amount = [0,1,0]
4-sekund: Iliq stakanga quyish. amount = [0,0,0]
Jami 4 sekund ketadi.
```

## TASK

Sizga sovuq, iliq va issiq suv quyish kerak bo'lgan stakanlar soni berilgan `amount` massivi uzunligi 3. Har sekundda siz:

- Yoki ikkita har xil turdagi stakanga suv quyishingiz mumkin.
- Yoki bitta stakanga suv quyishingiz mumkin.

Barcha stakanlarni to'ldirish uchun ketadigan minimal sekundlar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(amount []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
