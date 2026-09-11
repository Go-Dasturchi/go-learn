# Generate parentheses

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

## TASK

Sizga `n` soni berilgan, u ochiq-yopiq qavslar juftligining sonini bildiradi. Hammasi to'g'ri yopilgan bo'lishi sharti bilan yasalishi mumkin bo'lgan barcha kombinatsiyalarni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
