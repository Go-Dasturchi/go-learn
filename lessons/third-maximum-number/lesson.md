# Third Maximum Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [3,2,1]
Output: 1
Tushuntirish: Uchinchi eng katta son 1.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Unda uchinchi eng katta **o'zaro farqli** (distinct) sonni qaytaring. Agar bunday son mavjud bo'lmasa, eng katta sonni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massivni saralash shart emas — faqat eng katta uchta noyob qiymatni (birinchi, ikkinchi, uchinchi) bitta o'tishda kuzatib borish mumkin.
2. Uchta o'zgaruvchi (birinchi, ikkinchi, uchinchi eng katta) saqlang; har bir sonni avval takrorlanmasligi uchun tekshiring, so'ng u qaysi o'ringa mos kelishiga qarab o'zgaruvchilarni siljiting; agar uchinchisi topilmagan bo'lsa, javob sifatida eng kattasini qaytaring.
