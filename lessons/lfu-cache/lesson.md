# LFU Cache

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: capacity = 2, queries = [["put", 1, 1], ["put", 2, 2], ["get", 1], ["put", 3, 3], ["get", 2], ["get", 3]]
Output: [1, -1, 3]
Tushuntirish:
put(1,1), put(2,2).
get(1) -> 1 qaytaradi (1 ning freq=2).
put(3,3) -> freq si eng kam bo'lgan 2 ni o'chiradi.
get(2) -> -1 qaytaradi (o'chirilgan).
get(3) -> 3 qaytaradi.
```

## TASK

Tayyor LFU (Least Frequently Used) cache xotira kutubxonalarisiz ushbu kesh tizimini yarating.

Tizim quyidagi operatsiyalarni O(1) o'rtacha vaqt murakkabligida qo'llab-quvvatlashi kerak:

- `get(key)`: keshda kalit mavjud bo'lsa uning qiymatini qaytaradi.
- `put(key, value)`: agar kalit mavjud bo'lsa uni yangilaydi. Aks holda, kalit-qiymat juftini qo'shadi. Agar kesh sig'imi to'lib qolsa, eng kam ishlatilgan (eng kichik chastotaga ega) elementni o'chirib yuborishi kerak.

Agar chastotalar bir xil bo'lsa, eng eski ishlatilgan (LRU) element o'chiriladi.

_Eslatma: Bu yerda funksiya test sifatida ketma-ket kesh operatsiyalari massivida bajariladi. `solve` funksiyasiga sig'im va test so'rovlari beriladi, u barcha `get` so'rovlarining natijalarini qaytaradi._

Quyidagi funksiyani to'ldiring:

```go
func solve(capacity int, ops []string, args [][]int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. O(1) vaqt uchun ikkita xesh-jadval kerak bo'ladi: biri kalitdan tugunga, ikkinchisi esa har bir chastota (frequency) qiymatidan o'sha chastotaga ega elementlarning ikki tomonlama bog'langan ro'yxatiga (doubly linked list) — shu bilan bir chastota ichida eng eski/eng yangi elementni ham kuzatib borish mumkin bo'ladi.
2. Har bir elementni get yoki put qilganda uning chastotasini bittaga oshirib, eski chastota ro'yxatidan yangi chastota ro'yxatining boshiga o'tkazing. Doim minFreq (eng kichik chastota) qiymatini kuzatib boring — agar element ko'chirilgandan keyin eski chastota ro'yxati bo'shab qolsa va u minFreq ga teng bo'lsa, minFreq ni bittaga oshiring; sig'im to'lganda esa aynan minFreq ro'yxatining oxiridagi (eng uzoq ishlatilmagan) elementni o'chiring.
