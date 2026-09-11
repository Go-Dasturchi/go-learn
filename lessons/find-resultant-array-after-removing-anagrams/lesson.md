# Find Resultant Array After Removing Anagrams

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: words = ["abba","baba","bbaa","cd","cd"]
Output: ["abba","cd"]
Tushuntirish:
- "baba" va "abba" anagramma. "baba" o'chiriladi. ["abba","bbaa","cd","cd"]
- "bbaa" va "abba" anagramma. "bbaa" o'chiriladi. ["abba","cd","cd"]
- "cd" va "cd" anagramma. Ikkinchi "cd" o'chiriladi. ["abba","cd"]
```

## TASK

Sizga satrlar massivi `words` berilgan, har bir element kichik ingliz harflaridan iborat.

Quyidagi amalni bajarish mumkin bo'lgunicha takrorlang:

1. Agar `0 < i < words.length` bo'lib, `words[i-1]` va `words[i]` o'zaro **anagramma** bo'lsa, `words[i]` ni o'chirib tashlang.
2. So'ng massivni qayta ko'rib chiqing.

Barcha operatsiyalar tugagandan so'ng qolgan massivni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(words []string) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
