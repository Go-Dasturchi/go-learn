# Summary Ranges

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Tushuntirish: Oraliqlar quyidagicha tuziladi:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
```

## TASK

Sizga tartiblangan va o'zaro takrorlanmaydigan butun sonlardan iborat `nums` massivi berilgan.

Massivdagi barcha raqamlarni to'liq o'z ichiga oladigan, eng kichik (bir-biri bilan tutashgan) oraliqlarni ifodalovchi ro'yxatni qaytaring. Har bir son aniq bir oraliqqa tegishli bo'lishi kerak.

Har bir oraliq `[a, b]` quyidagi ko'rinishdagi satr sifatida chiqarilishi lozim:

- `"a->b"`, agar `a != b` bo'lsa
- `"a"`, agar `a == b` bo'lsa

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
