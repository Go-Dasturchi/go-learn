# Longest common prefix

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

## TASK

Satrlar massivi `strs` berilgan. Bu satrlarning barchasida qatnashgan eng uzun umumiy prefiksni (boshlang'ich satrni) toping.

Agar umumiy prefiks mavjud bo'lmasa, bo'sh satr `""` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(strs []string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
