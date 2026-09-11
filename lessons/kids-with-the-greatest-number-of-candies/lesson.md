# Kids With the Greatest Number of Candies

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true]
Tushuntirish:
- 1-bola: 2 + 3 = 5 (>= 5, true)
- 2-bola: 3 + 3 = 6 (>= 5, true)
- 3-bola: 5 + 3 = 8 (>= 5, true)
- 4-bola: 1 + 3 = 4 (< 5, false)
- 5-bola: 3 + 3 = 6 (>= 5, true)
```

## TASK

Sizga `n` ta bolaning konfetlari sonini bildiruvchi `candies` massivi va `extraCandies` (qo'shimcha konfetlar) butun soni berilgan.

Har bir bola uchun, agar biz barcha `extraCandies` ni unga bersak, uning konfetlari soni barcha bolalar orasidagi hozirgi eng ko'p konfetlar sonidan katta yoki teng bo'lishini aniqlang.
Natijani har bir bola uchun mantiqiy (boolean) massiv ko'rinishida qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(candies []int, extraCandies int) []bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
