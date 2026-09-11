# Goal Parser Interpretation

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: command = "G()(al)"
Output: "Goal"
Tushuntirish: "G" -> "G", "()" -> "o", "(al)" -> "al". Ularni qo'shsak "Goal".
```

## TASK

Sizga Goal parsari uchun qoidalar berilgan:

- `"G"` satri `"G"` deb tarjima qilinadi.
- `"()"` satri `"o"` deb tarjima qilinadi.
- `"(al)"` satri `"al"` deb tarjima qilinadi.

Sizga `command` satri berilgan. Yuqoridagi qoidalar asosida uni o'zgartirib, hosil bo'lgan tarjimani qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(command string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
