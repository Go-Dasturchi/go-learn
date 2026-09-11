# Decode the Message

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: key = "the quick brown fox jumps over the lazy dog", message = "vkbs bs t suepuv"
Output: "this is a secret"
Tushuntirish: 't' -> 'a', 'h' -> 'b', 'e' -> 'c' va hokazo (birinchi marta paydo bo'lishiga qarab).
Shundan foydalanib xabar shifrdan yechiladi.
```

## TASK

Sizga ikkita satr `key` va `message` berilgan.
`key` bu o'ziga xos shifrlash jadvalini yaratadi. Undagi harflar tartibi alifboning 26 ta harfiga mos keladi (birinchi kelgan harf `a`, ikkinchisi `b` va hokazo, takrorlangan harflar o'tkazib yuboriladi, bo'sh joy esa jadvalga ta'sir qilmaydi).
Bo'sh joylar ` ` kodi o'zgarishsiz qoladi.

Siz `message` ni shu `key` asosida topilgan jadval orqali dekodlab qaytarishingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(key string, message string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
