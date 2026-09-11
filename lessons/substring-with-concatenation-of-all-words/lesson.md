# Substring with Concatenation

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0, 9]
```

## TASK

Satr `s` va bir xil uzunlikdagi so'zlardan iborat `words` massivi berilgan. `words` massividagi barcha so'zlarning har qanday tartibdagi kombinatsiyasidan iborat bo'lgan qism-satrlarning boshlang'ich indekslarini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, words []string) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
