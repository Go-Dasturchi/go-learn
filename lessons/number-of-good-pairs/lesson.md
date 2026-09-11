# Number of Good Pairs

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3,1,1,3]
Output: 4
Tushuntirish: Yaxshi juftliklar indekslari bo'yicha: (0,3), (0,4), (3,4), (2,5). Ularning soni 4 ta.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan.

Agar `nums[i] == nums[j]` va `i < j` bo'lsa, u holda `(i, j)` juftlikni **yaxshi juftlik (good pair)** deb ataymiz.

Massivdagi jami yaxshi juftliklar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir sonni ilgari necha marta ko'rganingizni bilsangiz, juftliklar sonini ikkinchi ichki tsiklsiz ham hisoblash mumkin — buning uchun map dan foydalaning.
2. Massivni bir marta aylanib, har bir son uchun uni map'ga qo'shishdan oldin shu songa mos hisoblagichning joriy qiymatini natijaga qo'shing (chunki u qadar necha marta uchragan bo'lsa, shuncha yangi juftlik hosil bo'ladi), keyin hisoblagichni oshiring.
