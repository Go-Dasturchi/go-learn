# 11 — Kafka

## THEORY

"Worker Pools" va "Producer/Consumer" darslarida, bitta dastur ichida, goroutine'lar orasida kanal orqali xabar almashishni ko'rgan edik. Lekin real tizimlarda, ko'pincha xabarni yuboruvchi (masalan, "buyurtma qabul qilish" xizmati) va uni qabul qiluvchi (masalan, "email yuborish" xizmati) — **butunlay boshqa-boshqa dasturlar**, hatto boshqa serverlarda ishlaydi. Ular orasida oddiy Go kanali ishlamaydi — ularga **tarmoq orqali ishlaydigan xabar navbati** kerak.

**Apache Kafka** — aynan shu vazifani bajaradigan, katta hajmdagi xabarlarni ishonchli yetkazib beruvchi tizim. U uchta asosiy tushunchaga tayanadi:

- **Producer (ishlab chiqaruvchi)** — xabar yuboruvchi dastur.
- **Topic (mavzu)** — xabarlar yig'iladigan "kanal nomi" (masalan, `"buyurtmalar"`).
- **Consumer (iste'molchi)** — xabarlarni o'qib, ularga ishlov beruvchi dastur.

### 1-qadam: Kafka'ni ishga tushirish (Docker orqali — eng oson yo'l)

Kafka'ni to'g'ridan-to'g'ri o'rnatish murakkab (u Zookeeper yoki KRaft kabi qo'shimcha komponentlarga muhtoj), shuning uchun "Docker Compose" darsida ko'rgan usul bilan ishga tushirish qulay:

```yaml
services:
  kafka:
    image: apache/kafka:3.7.0
    ports:
      - "9092:9092"
    environment:
      KAFKA_NODE_ID: 1
      KAFKA_LISTENERS: PLAINTEXT://:9092,CONTROLLER://:9093
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
```

```bash
docker compose up -d
```

### 2-qadam: Topic yaratish va xabar yuborish/o'qish (konteyner ichidan)

```bash
docker exec -it <konteyner_nomi> /opt/kafka/bin/kafka-topics.sh \
  --create --topic buyurtmalar --bootstrap-server localhost:9092

docker exec -it <konteyner_nomi> /opt/kafka/bin/kafka-console-producer.sh \
  --topic buyurtmalar --bootstrap-server localhost:9092
# > interaktiv rejimda xabar yozib, Enter bosing

docker exec -it <konteyner_nomi> /opt/kafka/bin/kafka-console-consumer.sh \
  --topic buyurtmalar --from-beginning --bootstrap-server localhost:9092
# yuborilgan xabarlarni shu yerda ko'rasiz
```

### 3-qadam: Go'dan Kafka bilan ishlash

Standart kutubxonada Kafka drayveri yo'q — mashhur kutubxona `github.com/segmentio/kafka-go`:

```go
import kafka "github.com/segmentio/kafka-go"

writer := &kafka.Writer{
	Addr:  kafka.TCP("localhost:9092"),
	Topic: "buyurtmalar",
}
writer.WriteMessages(context.Background(), kafka.Message{Value: []byte("yangi buyurtma: #42")})

reader := kafka.NewReader(kafka.ReaderConfig{
	Brokers: []string{"localhost:9092"},
	Topic:   "buyurtmalar",
})
msg, _ := reader.ReadMessage(context.Background())
fmt.Println(string(msg.Value))
```

**Nega Kafka, oddiy xabar navbatidan (masalan, RabbitMQ) farq qiladi.** Kafka, xabarlarni o'qilgandan keyin ham **darhol o'chirmaydi** — ular, belgilangan muddat (masalan, 7 kun) davomida saqlanadi, va bir nechta **turli** consumer, bir xil xabarlarni **mustaqil ravishda** (turli tezlikda) o'qishi mumkin. Bu, Kafka'ni "xabar jurnali" (log) sifatida ishlashga o'xshatadi.

**Nega bu darsda haqiqiy Kafka ishlatilmaydi.** Bu, checker ishlaydigan tarmoqsiz, izolyatsiyalangan muhitda mavjud emas. Shu sababli, producer/consumer va navbat g'oyasini, "Queue" darsida qurgan tuzilmaga o'xshab, oddiy Go slice orqali mashq qilamiz:

```go
type XabarNavbati struct {
	xabarlar []string
}

func (n *XabarNavbati) Yubor(xabar string) {
	n.xabarlar = append(n.xabarlar, xabar) // producer — navbat oxiriga qo'shadi
}

func (n *XabarNavbati) Ol() (string, bool) {
	if len(n.xabarlar) == 0 {
		return "", false // consumer uchun xabar yo'q
	}
	xabar := n.xabarlar[0]
	n.xabarlar = n.xabarlar[1:] // FIFO — birinchi kirgan, birinchi chiqadi
	return xabar, true
}
```

Bu — Kafka'dagi "topic"ning, bitta jarayon ichidagi, soddalashtirilgan ko'rinishi: `Yubor` — producer, `Ol` — consumer, va ikkalasi orasidagi FIFO navbat — Kafka topic'ining asosiy g'oyasi.

## EXAMPLE

```go
package main

import "fmt"

type XabarNavbati struct {
	xabarlar []string
}

func (n *XabarNavbati) Yubor(xabar string) {
	n.xabarlar = append(n.xabarlar, xabar)
}

func (n *XabarNavbati) Ol() (string, bool) {
	if len(n.xabarlar) == 0 {
		return "", false
	}
	xabar := n.xabarlar[0]
	n.xabarlar = n.xabarlar[1:]
	return xabar, true
}

func main() {
	navbat := &XabarNavbati{}

	navbat.Yubor("buyurtma #1")
	navbat.Yubor("buyurtma #2")

	xabar, bor := navbat.Ol()
	fmt.Println(xabar, bor)

	xabar, bor = navbat.Ol()
	fmt.Println(xabar, bor)

	xabar, bor = navbat.Ol()
	fmt.Println(xabar, bor)
}
```

Natija:

```
buyurtma #1 true
buyurtma #2 true
 false
```

## TASK

`XabarNavbati.Ol() (string, bool)` methodini to'ldiring:

1. Agar navbat bo'sh bo'lsa (`len(n.xabarlar) == 0`), `"", false` qaytaring.
2. Aks holda, **birinchi** xabarni (`n.xabarlar[0]`) oling, uni navbatdan olib tashlang (`n.xabarlar = n.xabarlar[1:]`), va `xabar, true` qaytaring.

`Yubor` methodi va `main()` funksiyasini o'zgartirish shart emas.

**Terminalda sinash uchun** (checker tekshirmaydi): Docker Compose orqali Kafka'ni ishga tushirib, `kafka-console-producer.sh`/`kafka-console-consumer.sh` bilan xabar yuborib-qabul qilib ko'ring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if len(n.xabarlar) == 0 { return "", false }`.
2. `xabar := n.xabarlar[0]; n.xabarlar = n.xabarlar[1:]; return xabar, true`.
