# 12 — Password Hashing

## THEORY

Hech qanday jiddiy dastur foydalanuvchi parolini **ochiq matn** holida saqlamasligi kerak — agar ma'lumotlar bazasi kimningdir qo'liga tushib qolsa, barcha parollar darhol oshkor bo'lib qoladi. Buning o'rniga, parolning **xeshi (hash)** saqlanadi — bu, parolning "bir tomonlama" o'zgartirilgan ko'rinishi: xeshdan parolning o'ziga qaytib bo'lmaydi, lekin bir xil parol har doim bir xil xeshni beradi.

**Xesh hisoblash — `crypto/sha256`:**

```go
import (
	"crypto/sha256"
	"encoding/hex"
)

func parolniHash(parol string) string {
	summa := sha256.Sum256([]byte(parol))
	return hex.EncodeToString(summa[:])
}
```

`sha256.Sum256(...)` — berilgan baytlarning SHA-256 xeshini hisoblaydi, va **qat'iy o'lchamli** (32 bayt) massiv (`[32]byte`) qaytaradi — "Arrays" darsida ko'rgan qat'iy o'lchamli tur. `hex.EncodeToString(summa[:])` — "Encoding/Decoding" darsida ko'rgan `encoding/hex` paketi orqali, bu baytlarni odam o'qiy oladigan matn shakliga o'tkazadi (`summa[:]` — array'ni slice'ga aylantirish, chunki `hex.EncodeToString` slice kutadi).

**Tekshirish — parolni qayta hesh qilib, solishtirish:**

```go
func parolniTekshir(parol, saqlanganHash string) bool {
	return parolniHash(parol) == saqlanganHash
}
```

Ro'yxatdan o'tishda parol xeshlanib saqlanadi; kirishda, foydalanuvchi kiritgan parol **qayta xeshlanib**, saqlangan xesh bilan solishtiriladi — asl parolning o'zi hech qachon, hech qayerda saqlanmaydi.

**MUHIM XAVFSIZLIK OGOHLANTIRISHI: `sha256`ning o'zi haqiqiy parollar uchun YETARLI EMAS.** Bu dars, xeshlashning **mexanizmini** (bir tomonlama funksiya, qayta hisoblab solishtirish) o'rgatish uchun, ataylab eng oddiy vositadan (`sha256`, standart kutubxonada bor) foydalanadi. Lekin haqiqiy tizimlarda parollar uchun **maxsus, ataylab sekin** qilib yaratilgan algoritmlar (masalan, `bcrypt`, `scrypt`, `argon2`) ishlatilishi **shart**: `sha256` juda **tez** hisoblanadi, va zamonaviy kompyuterlar sekundiga milliardlab variantni "urinib ko'rish" (brute-force) orqali oddiy parollarni topib olishi mumkin. `bcrypt` kabi algoritmlar esa, ataylab sekin ishlaydi (va "tuz" — salt — qo'shadi), bu esa bunday hujumlarni amaliy jihatdan foydasiz qiladi. Bu — kursning "faqat standart kutubxona" tamoyiliga rioya qilish uchun ataylab soddalashtirilgan misol, real loyihada esa albatta maxsus parol xeshlash kutubxonasi ishlatilishi kerak.

## EXAMPLE

```go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func parolniHash(parol string) string {
	summa := sha256.Sum256([]byte(parol))
	return hex.EncodeToString(summa[:])
}

func parolniTekshir(parol, saqlanganHash string) bool {
	return parolniHash(parol) == saqlanganHash
}

func main() {
	hash := parolniHash("MaxfiyParol123")
	fmt.Println(hash)
	fmt.Println(parolniTekshir("MaxfiyParol123", hash))
	fmt.Println(parolniTekshir("notogriParol", hash))
}
```

Natija:

```
9822357d23a2be58c0f646b2be746a27185e5cef2552324b79ce25ca445c1fc7
true
false
```

## TASK

`parolniHash(parol string) string` funksiyasi berilgan. Uni `crypto/sha256` va `encoding/hex` yordamida shunday to'ldiringki, u `parol`ning SHA-256 xeshini, o'n oltilik (hex) matn ko'rinishida qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `summa := sha256.Sum256([]byte(parol))` — bu `[32]byte` massiv qaytaradi.
2. `return hex.EncodeToString(summa[:])` — massivni slice'ga aylantirib (`summa[:]`), hex matnga kodlang.
