# Find All Numbers Disappeared in an Array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Tushuntirish: 1 dan 8 gacha bo'lgan sonlar ichida 5 va 6 yo'q.
```

## TASK

Sizga `n` o'lchamli `nums` massivi berilgan, bunda `nums[i]` qiymati `[1, n]` oralig'ida. Ba'zi elementlar ikki marta, ba'zilari esa bir marta uchraydi.

`[1, n]` oralig'idagi massivda mavjud bo'lmagan barcha sonlarni massiv ko'rinishida qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. [1, n] oralig'idagi har bir son massivda uchraganini belgilab boradigan boolean massiv haqida o'ylang.
2. n+1 uzunlikdagi boolean massiv yaratib, nums'dagi har bir qiymatni indeks sifatida ishlatib true deb belgilang, so'ngra 1 dan n gacha false qolgan indekslarni natijaga qo'shing.
