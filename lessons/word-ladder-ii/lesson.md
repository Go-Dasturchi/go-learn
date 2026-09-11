# Word Ladder II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **50 daqiqa**

## EXAMPLE

```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
```

## TASK

Sizga boshlang'ich so'z `beginWord`, tugash so'zi `endWord`, hamda so'zlar ro'yxati `wordList` berilgan.
Boshlang'ich so'zdan tugash so'ziga o'tishning **barcha eng qisqa yo'llarini** toping.
O'tish shartlari:

1. Har bir qadamda faqatgina bitta harfni o'zgartirish mumkin.
2. O'zgartirilgan har bir so'z `wordList` da mavjud bo'lishi shart. `beginWord` ro'yxatda bo'lishi shart emas.

Agar bunday yo'l mavjud bo'lmasa, bo'sh ro'yxat qaytaring.

_Eslatma: Output massivining tartibi ahamiyatsiz (Testda bir xil natija ekanligiga e'tibor qilinadi, lekin satr shaklida qat'iy tartib kutilgani uchun javobingizni leksikografik tartiblab oling yoki shunchaki bir xil test javobi formatida chiqaring. Test faqat elementlarga qarab tekshiradi)._

Quyidagi funksiyani to'ldiring:

```go
func solve(beginWord string, endWord string, wordList []string) [][]string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Faqat "Word Ladder" dagi kabi BFS bilan uzunlikni topish yetarli emas — bu safar barcha eng qisqa yo'llarni saqlab qolish kerak, shuning uchun BFS ni "qatlam-qatlam" (level by level) bajarib, har bir so'zga olib boruvchi barcha yo'llarni to'plang.
2. Har bir qatlamda joriy so'zlarga qadar yig'ilgan barcha to'liq yo'llarni saqlang (xarita: so'z -> yo'llar ro'yxati). Bitta harfni almashtirib navbatdagi qatlamdagi so'zlarga o'ting, lekin FAQAT shu qatlamda topilgan so'zlarni keyingi qatlamga o'tkazing va ularni umumiy so'zlar to'plamidan o'chirib tashlang (aks holda ortga qaytish yoki cheksiz aylanish yuzaga kelishi mumkin); endWord birinchi marta uchragan qatlamda to'xtang — undan keyingi qatlamlarga hojat yo'q, chunki ular endi eng qisqa bo'lmaydi.
