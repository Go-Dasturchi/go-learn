# Move zeroes

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [0,1,0,3,12]
Output: [1, 3, 12, 0, 0]
```

## TASK

Butun sonlardan iborat `nums` massivi berilgan. Barcha nollarni (0) uning oxiriga o'tkazing. Boshqa nol bo'lmagan elementlarning o'zaro joylashuv tartibini saqlab qoling.

Izoh: Bu masalada massiv elementlarini qaytarishingiz talab etiladi (aslida in-place qilinadi). Natijani array ko'rinishida qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Nolmas elementlarni massiv boshiga "siljitish" pozitsiyasi (write-pointer) g'oyasini o'ylab ko'ring.
2. Bitta yozuv indeksi tuting; massivni chapdan o'ngga aylanib, nolga teng bo'lmagan har bir elementni shu indeksga qo'yib indeksni oshiring, keyin qolgan barcha o'rinlarni nol bilan to'ldiring.
