# Count Odd Numbers in an Interval Range

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: low = 3, high = 7
Output: 3
Tushuntirish: Toq sonlar: 3, 5, 7.
```

## TASK

Sizga ikkita manfiy bo'lmagan butun `low` va `high` sonlari berilgan. `[low, high]` oralig'idagi toq sonlar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(low int, high int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Barcha sonlarni birma-bir sanab chiqish shart emas — oraliqdagi son va uning chekkalarining juft/toqligidan formula orqali topish mumkin.
2. (high-low)/2 orqali asosiy sonni toping, so'ngra agar low yoki high toq bo'lsa natijaga 1 qo'shing.
