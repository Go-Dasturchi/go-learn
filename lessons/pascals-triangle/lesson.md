# Pascal's Triangle

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

## TASK

Sizga butun son ko'rinishidagi `numRows` berilgan. Paskal uchburchagining (Pascal's triangle) dastlabki `numRows` qatorini qaytaring.

Paskal uchburchagida har bir raqam uning to'g'ridan-to'g'ri yuqorisidagi ikkita raqam yig'indisiga teng.

Quyidagi funksiyani to'ldiring:

```go
func solve(numRows int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir qatorni oldingi qatordan hosil qilish mumkin — chetlardagi elementlar doim 1, ichkilari esa yuqoridagi ikkita qo'shni elementning yig'indisi.
2. Qatorlarni birma-bir quring: har qator uchun boshi va oxirini 1 qilib belgilang, o'rtadagi har bir indeks uchun oldingi qatordagi `j-1` va `j` indeksdagi qiymatlarni qo'shib joylashtiring.
