# Guess Number Higher or Lower

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: n = 10, pick = 6
Output: 6
```

## TASK

Biz u yoki bu sonni taxmin qilish o'yinini o'ynamoqdamiz. Men 1 dan `n` gacha bo'lgan oraliqdan bitta son tanladim.

Sizga bir `guess(num)` API mavjud bo'lib u quyidagicha javob qaytaradi:

- `-1`: Siz taxmin qilgan son `pick` dan katta.
- `1`: Siz taxmin qilgan son `pick` dan kichik.
- `0`: Siz taxmin qilgan son `pick` ga teng. Siz topdingiz!

_Test tizimida biz `pick` qiymatini sizga parameter sifatida uzatamiz. Siz `num == pick` bo'lgan num ni topishingiz kerak._

Quyidagi funksiyani to'ldiring:

```go
func solve(n int, pick int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
