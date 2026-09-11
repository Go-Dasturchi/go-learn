# Best Time to Buy and Sell Stock

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: prices = [7,1,5,3,6,4]
Output: 5
Tushuntirish: 2-kuni sotib oling (narx = 1) va 5-kuni soting (narx = 6), foyda = 6 - 1 = 5.
E'tibor bering, 2-kuni sotib olib, 1-kuni sotish mumkin emas, chunki sotishdan oldin sotib olishingiz kerak.
```

## TASK

Sizga `prices` deb nomlangan massiv berilgan, unda `prices[i]` ma'lum bir aksiyaning `i`-kuni narxini bildiradi.

Siz bir kunni aksiya sotib olish uchun va kelajakdagi boshqa bir kunni aksiyani sotish uchun tanlab, maksimal foydani topishingiz kerak.

Agar hech qanday foyda topa olmasangiz, 0 ni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(prices []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
