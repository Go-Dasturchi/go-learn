# Group Anagrams

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

## TASK

Sizga `strs` deb nomlangan satrlar massivi berilgan. Bu massivdagi anagrammalarni birga guruhlang. Natijani istalgan tartibda qaytarishingiz mumkin.

**Anagramma** bu odatda boshqa so'zning yoki iboraning harflarini qayta tartiblash orqali hosil bo'ladigan so'z yoki ibora (masalan, "eat", "tea", "ate").

Quyidagi funksiyani to'ldiring:

```go
func solve(strs []string) [][]string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
