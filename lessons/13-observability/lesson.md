# 13 — Observability

## THEORY

Dasturingiz production'da ishlab turibdi. Bir kuni, foydalanuvchilar "sayt sekin ishlayapti" deb shikoyat qiladi. Savol: **qayerda** muammo? Baza sekinmi? Qaysi endpoint? Qaysi soatda boshlangan? **Observability (kuzatuvchanlik)** — aynan shu savollarga javob berish uchun, dastur ichidan **tashqariga signal chiqarish** amaliyoti.

Observability, an'anaviy ravishda **uchta ustunga** tayanadi:

### 1. Logging (loglar) — "Nima bo'ldi?"

"Logging" darsida ko'rganimizdek — har bir muhim hodisa (xatolik, so'rov, holat o'zgarishi) haqida yozib boriladigan, vaqt tamg'asi bilan belgilangan xabarlar.

### 2. Metrics (metrikalar) — "Necha marta? Qancha vaqt?"

Metrikalar — vaqt bo'yicha **yig'iladigan raqamlar**: so'rovlar soni, xatoliklar foizi, javob vaqti. Ular, alohida hodisalar emas, balki **umumiy tendensiyani** ko'rsatadi.

**Eng keng tarqalgan metrika turlari:**

- **Counter (hisoblagich)** — faqat **ko'payadigan** raqam (masalan, "jami qabul qilingan so'rovlar soni").
- **Gauge (o'lchagich)** — yuqoriga ham, pastga ham o'zgarishi mumkin bo'lgan raqam (masalan, "hozirgi faol ulanishlar soni").
- **Histogram** — qiymatlarning **taqsimotini** kuzatadi (masalan, "so'rovlarning necha foizi 100ms'dan tez javob berdi").

Go'da, mashhur `prometheus/client_golang` kutubxonasi yordamida:

```go
so'rovlarSoni := prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_sorovlar_jami",
	Help: "Qabul qilingan HTTP so'rovlar soni",
})
so'rovlarSoni.Inc() // har bir so'rovda +1
```

### 3. Tracing (tracing) — "Qayerda vaqt ketdi?"

Bitta so'rov, ko'pincha bir nechta xizmat orqali o'tadi ("Microservices" darsini eslang: Foydalanuvchi xizmati → Buyurtma xizmati → Bildirishnoma xizmati). **Distributed tracing** — bitta so'rovning, **barcha xizmatlar bo'ylab** qancha vaqt olganini, qaysi qismda "tiqilib qolganini" ko'rsatadi.

### Terminalda ko'rish — Prometheus + Grafana (Docker Compose orqali)

```yaml
services:
  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
```

```bash
docker compose up -d
```

`http://localhost:9090` — Prometheus konsoli (metrikalarni yig'ish), `http://localhost:3000` — Grafana (metrikalarni **grafik** ko'rinishida ko'rsatish).

**Nega bu darsda haqiqiy Prometheus ishlatilmaydi.** Bu, tashqi tizim va tarmoq ulanishini talab qiladi. Shu sababli, **Counter** va **so'rov vaqtini o'lchash** g'oyasini, oddiy Go struct orqali mashq qilamiz — "Redis" darsidagi kabi, vaqtni **tashqaridan in'ektsiya qilib** (deterministik test uchun):

```go
type Metrikalar struct {
	sorovlarSoni int
	jamiVaqtMs   int64
}

func (m *Metrikalar) SorovniQayd(davomiylikMs int64) {
	m.sorovlarSoni++
	m.jamiVaqtMs += davomiylikMs
}

func (m *Metrikalar) OrtachaVaqtMs() float64 {
	if m.sorovlarSoni == 0 {
		return 0
	}
	return float64(m.jamiVaqtMs) / float64(m.sorovlarSoni)
}
```

`SorovniQayd` — har bir so'rov tugagach chaqiriladi (Counter'ning `Inc()`iga o'xshab, lekin vaqtni ham qo'shadi), `OrtachaVaqtMs` esa — yig'ilgan ma'lumotdan **o'rtacha javob vaqtini** hisoblaydi, aynan Prometheus'dagi Histogram'ning soddalashtirilgan versiyasi.

## EXAMPLE

```go
package main

import "fmt"

type Metrikalar struct {
	sorovlarSoni int
	jamiVaqtMs   int64
}

func (m *Metrikalar) SorovniQayd(davomiylikMs int64) {
	m.sorovlarSoni++
	m.jamiVaqtMs += davomiylikMs
}

func (m *Metrikalar) OrtachaVaqtMs() float64 {
	if m.sorovlarSoni == 0 {
		return 0
	}
	return float64(m.jamiVaqtMs) / float64(m.sorovlarSoni)
}

func main() {
	m := &Metrikalar{}

	m.SorovniQayd(100)
	m.SorovniQayd(200)
	m.SorovniQayd(300)

	fmt.Println(m.sorovlarSoni, m.OrtachaVaqtMs())
}
```

Natija:

```
3 200
```

## TASK

`Metrikalar.OrtachaVaqtMs() float64` methodini to'ldiring:

1. Agar hali birorta ham so'rov qayd qilinmagan bo'lsa (`m.sorovlarSoni == 0`), `0` qaytaring (nolga bo'lishning oldini olish uchun).
2. Aks holda, `m.jamiVaqtMs`ni `m.sorovlarSoni`ga bo'lib (ikkalasini ham `float64`ga aylantirib), o'rtacha qiymatni qaytaring.

`SorovniQayd` methodi va `main()` funksiyasini o'zgartirish shart emas.

**Terminalda sinash uchun** (checker tekshirmaydi): yuqoridagi Docker Compose bilan Prometheus va Grafana'ni ishga tushirib, `http://localhost:9090` va `http://localhost:3000` manzillarini brauzerda ochib ko'ring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if m.sorovlarSoni == 0 { return 0 }`.
2. `return float64(m.jamiVaqtMs) / float64(m.sorovlarSoni)`.
