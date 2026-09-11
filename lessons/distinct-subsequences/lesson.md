# Distinct Subsequences

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "rabbbit", t = "rabbit"
Output: 3
Tushuntirish:
Quyidagi tarzda s satridan 3 xil usulda "rabbit" so'zini hosil qilish mumkin:
rabbbit (uchinchi b ni o'chirish)
rabbbit (ikkinchi b ni o'chirish)
rabbbit (birinchi b ni o'chirish)
```

## TASK

Sizga `s` va `t` ikkita satr berilgan. `s` satrining qismlari (subsequence) ichida `t` satriga teng bo'lganlarining umumiy sonini qaytaring.

Satrning qismi (subsequence) bu asl satrdagi ba'zi belgilarni (yoki hech birini) o'chirish orqali hosil bo'ladigan yangi satrdir, qolgan belgilarni o'zaro joylashuv tartibini o'zgartirmagan holatda. Masalan, `"ACE"` bu `"ABCDE"` ning qismi hisoblanadi, ammo `"AEC"` emas.

Test keyslar javob 32-bitli butun songa (Int32) sig'ishiga moslab tuzilgan.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, t string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
