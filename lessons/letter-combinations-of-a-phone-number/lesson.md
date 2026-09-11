# Letter Combinations of a Phone Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

## TASK

Sizga 2 dan 9 gacha bo'lgan raqamlardan iborat `digits` satri beriladi. Raqamlarga mos (eski telefonlardagi kabi) harflar orqali yasalishi mumkin bo'lgan barcha so'zlarni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(digits string) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir raqamga mos harflardan birini tanlab, keyingi raqamga o'tadigan rekursiv qidiruv (backtracking) haqida o'ylab ko'ring — bu xuddi daraxt shaklidagi barcha yo'llarni yig'ib chiqishga o'xshaydi.
2. Raqam-harflar moslamasini (masalan '2' -> "abc") tayyorlab, joriy indeks va joriy yig'ilgan satrni oluvchi rekursiv funksiya yozing: indeks digits uzunligiga yetganda joriy satrni natijaga qo'shing, aks holda joriy raqamga mos har bir harf uchun uni satrga qo'shib keyingi indeksga rekursiya qiling.
