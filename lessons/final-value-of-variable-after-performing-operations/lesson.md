# Final Value of Variable After Performing Operations

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: operations = ["--X","X++","X++"]
Output: 1
Tushuntirish: X dastlab 0.
--X: X = -1
X++: X = 0
X++: X = 1
```

## TASK

Dastlab `X` nomli o'zgaruvchi bor va uning qiymati `0` ga teng. Sizga satrlar ro'yxati `operations` berilgan. Har bir operatsiya `++X`, `X++`, `--X` yoki `X--` lardan biri.

Barcha operatsiyalar bajarilgandan keyin `X` ning yakuniy qiymatini qaytaring.

- `++X` va `X++` qiymatni 1 ga oshiradi.
- `--X` va `X--` qiymatni 1 ga kamaytiradi.

Quyidagi funksiyani to'ldiring:

```go
func solve(operations []string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir operatsiya satrida faqat '+' yoki '-' belgisi bor-yo'qligini bitta belgidan tekshirish kifoya.
2. Har bir operatsiya satrining ikkinchi belgisini (op[1]) tekshiring — u '+' bo'lsa hisoblagichni oshiring, aks holda kamaytiring.
