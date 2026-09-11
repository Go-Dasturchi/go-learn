# Remove Duplicate Letters

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "bcabc"
Output: "abc"
```

## TASK

Sizga kichik ingliz harflaridan tuzilgan `s` satri berilgan. Ushbu satrdan takrorlanuvchi harflarni shu tartibda olingki, hosil bo'lgan satrda har bir harf aynan bir marta qatnashsin va u barcha mumkin bo'lgan natijalar ichida leksikografik jihatdan eng kichigi bo'lsin.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Monotonik stek (stack) g'oyasidan foydalaning: natijaviy satrni harf-harf quring va har bir harfning satrda oxirgi marta qayerda uchrashini oldindan bilib oling.
2. Har bir belgi uchun: agar u allaqachon stekda bo'lsa o'tkazib yuboring. Aks holda, stekning tepasidagi harf joriy harfdan katta bo'lsa VA o'sha tepadagi harf satrning qolgan qismida yana uchraydigan bo'lsa (ya'ni uning lastIndex joriy pozitsiyadan katta bo'lsa), uni stekdan chiqarib tashlang — shu shartlarga mos kelmay qolguncha bu jarayonni takrorlang, so'ng joriy harfni stekka qo'shing.
