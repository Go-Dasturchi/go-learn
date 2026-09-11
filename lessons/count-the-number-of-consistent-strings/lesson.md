# Count the Number of Consistent Strings

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
Output: 2
Tushuntirish: "aaab" va "baa" satrlari faqat 'a' va 'b' harflaridan tashkil topgan.
```

## TASK

Sizga ruxsat etilgan harflardan iborat `allowed` satri va `words` satrlar massivi berilgan.

Agar `words` dagi biror satr faqatgina `allowed` satrida mavjud bo'lgan harflardan tuzilgan bo'lsa, uni **mos keluvchi (consistent)** deb ataymiz.

Mos keluvchi satrlar sonini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(allowed string, words []string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
