# Max Consecutive Ones

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [1,1,0,1,1,1]
Output: 3
Tushuntirish: Massivning birinchi va ikkinchi qismida ketma-ket uchta 1 bor.
```

## TASK

Sizga faqat `0` va `1` dan iborat ikkilik massiv `nums` berilgan. Uning ichidagi eng uzun ketma-ket `1` lar seriyasining uzunligini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massivni bir marta aylanib, joriy ketma-ket birlar seriyasining uzunligini kuzatib boring.
2. Ikkita hisoblagich tuting: joriy ketma-ketlik va eng yaxshi natija. `1` uchraganda joriysini oshirib eng yaxshisi bilan solishtiring, `0` uchraganda joriysini nolga qaytaring.
