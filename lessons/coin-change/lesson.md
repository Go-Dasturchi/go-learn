# Coin Change

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: coins = [1,5,11], amount = 11
Output: 1
Tushuntirish: 11 tanga bilan 11 ni yasash mumkin, bu eng kamosi.
```

## TASK

Sizga turli nominallardan iborat `coins` massivi va umumiy `amount` miqdori berilgan.

Shu miqdorni yasash uchun zarur bo'lgan eng kam tangalar sonini qaytaring. Agar hech qanday kombinatsiya bilan o'sha miqdorni yig'ish imkoni bo'lmasa, `-1` qaytaring.

Har bir nominal cheksiz miqdorda mavjud deb hisoblashingiz mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(coins []int, amount int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
