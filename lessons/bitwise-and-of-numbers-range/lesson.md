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

1. Barcha sonlarni birma-bir AND qilish shart emas — left va right ning ikkilik (binary) yozuvidagi umumiy old qismini (common prefix) topish kifoya, chunki oraliqda kamida bitta son bu prefiksdan keyingi bitlarni "buzadi".
2. left va right ni bir xil bo'lguncha bir vaqtda o'ngga siljiting (>>), har safar siljishlar sonini sanang; teng bo'lgach, natijani xuddi shuncha marta chapga siljitib (<<) qaytaring.
