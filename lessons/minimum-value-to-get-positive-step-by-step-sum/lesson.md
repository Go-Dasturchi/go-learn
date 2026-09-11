# Minimum Value to Get Positive Step by Step Sum

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [-3,2,-3,4,2]
Output: 5
Tushuntirish:
startValue = 5 bilan:
5 + (-3) = 2 >= 1
2 + 2 = 4 >= 1
4 + (-3) = 1 >= 1
1 + 4 = 5 >= 1
5 + 2 = 7 >= 1
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. `startValue` dan boshlang va `nums` elementlarini chapdan o'ngga birinchisi bilan qo'shing.

Quyidagi shart bajarilishi uchun zarur bo'lgan minimal musbat `startValue` ni toping:

- Har bir qadamdagi oraliq yig'indi (step by step sum) kamida `1` ga teng bo'lsin.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Kerakli `startValue` qiymati oraliq yig'indining eng past (eng manfiy) nuqtasiga bog'liq ekanini o'ylang.
2. `nums` bo'ylab yurib joriy yig'indini hisoblang va u qancha pastga tushsa ham eng kichik qiymatni saqlab boring; javob — shu eng kichik qiymatni 1 dan ayirib, teskari ishora bilan olish, ya'ni `1 - min`.
