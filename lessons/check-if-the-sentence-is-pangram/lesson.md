# Check if the Sentence Is Pangram

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Tushuntirish: sentence ingliz alifbosidagi barcha 26 ta harfni o'z ichiga oladi.
```

## TASK

Pangram — bu ingliz alifbosidagi har bir harf kamida bir marta ishtirok etgan gap.
Sizga faqat kichik ingliz harflaridan iborat `sentence` satri berilgan.

Agar `sentence` pangram bo'lsa `true` qaytaring, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(sentence string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
