# Roman to Integer

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "III"
Output: 3
Tushuntirish: III = 3.
```

## TASK

Rim raqamlari yettita turli belgi orqali ifodalanadi: `I`, `V`, `X`, `L`, `C`, `D` va `M`.

Sizga Rim raqamlaridan tashkil topgan satr `s` berilgan. Uning butun son (integer) ekvivalentini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir belgining qiymatini map orqali saqlab, keyingi belgi bilan solishtirishni o'ylang — kichik qiymat kattadan oldin kelsa (masalan IV), bu ayirishni bildiradi.
2. Satrni chapdan o'ngga aylanib, har bir belgi qiymatini joriy belgidan keyingi belgi qiymati bilan solishtiring: agar joriysi kichik bo'lsa uni yig'indidan ayiring, aks holda qo'shing.
