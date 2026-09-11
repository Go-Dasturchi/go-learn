# Nim Game

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: n = 4
Output: false
Tushuntirish: Siz 1 ta, 2 ta yoki 3 ta tosh olsangiz ham, do'stingiz keyingi navbatda qolgan barcha toshlarni olib yutadi. Sizda yutish imkoni yo'q.
```

## TASK

Siz do'stingiz bilan Nim o'yinini o'ynamoqdasiz:

1. Stolda bir uyum toshlar (jami `n` ta) turibdi.
2. Siz va do'stingiz navbatma-navbat bittadan uchtagacha (1, 2 yoki 3 ta) toshni olishingiz mumkin.
3. O'yinni birinchi bo'lib siz boshlaysiz.
4. Eng oxirgi toshni olgan o'yinchi yutadi.

Agar ikkala o'yinchi ham eng optimal darajada o'ynasa, siz g'alaba qozona olishingiz mumkinligini toping. G'alaba qozonsangiz `true`, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
