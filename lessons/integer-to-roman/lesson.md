# Integer to Roman

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: num = 3
Output: "III"
```

## TASK

Rim raqamlari yettita turli belgi orqali ifodalanadi: I, V, X, L, C, D, M. Butun son `num` berilgan. Uni Rim raqamlariga aylantiruvchi algoritm yozing.

Quyidagi funksiyani to'ldiring:

```go
func solve(num int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Rim raqamlarini alohida-alohida chiqarish o'rniga, qiymatlarni eng kattasidan boshlab greedy (ochko'z) tarzda ayirib borishni o'ylab ko'ring — bunda 4, 9, 40, 90, 400, 900 kabi maxsus kombinatsiyalarni ham alohida "qiymat" sifatida hisoblash qulay bo'ladi.
2. Qiymatlar va ularga mos belgilarni kamayish tartibida ikkita mos ro'yxatda saqlang (masalan 1000→"M", 900→"CM", ..., 1→"I"); ro'yxat bo'ylab yurib, num shu qiymatdan katta yoki teng bo'lguncha uni ayirib mos belgini natijaga qo'shing, so'ng keyingi qiymatga o'ting.
