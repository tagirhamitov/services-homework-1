# Архитектура
Код написан на языке Go версии 1.20

Есть 7 серверов для каждого из форматов:
- Нативный формат языка Go (GOB)
- XML
- JSON
- Google Protocol Buffers
- Apache Avro
- YAML
- MessagePack

У них один общий Dockerfile (server.dockerfile), в который передается аргумент SERIALIZATION_FORMAT. Он определяет, какой формат будет тестироваться каждым сервером.

Логика тестирования лежит в директории src/testing, для каждого из формата отдельный файл test_***.go (например, test_json.go).

Тестируемая структура описана в src/types/struct.go, а форматы в src/types/formats.go. Т.к. пакет xml в языке Go не поддерживает сериализацию словарей, я определил StringMap в файле src/types/string_map.go с собственной логикой сериализации.

Для Protocol Buffers proto-файл лежит в директории src/proto. Код генерируется в докейрфайле и размещается в директории src/pb.

Сервер слушает на порту 2000 и отвечает на запросы /get_result. Логика обработки http находится в src/app/server/main.go

Также есть прокси-сервер (src/app/proxy/main.go), который слушает на порту 2000 и отвечает на запросы /get_result?format=..., где в качестве формата нужно передать один из параметров:
- native
- xml
- json
- protobuf
- avro
- yaml
- message_pack

Прокси-сервер отправляет запрос в соответствующий контейнер и возвращает ответ.

# Запуск:
```bash
docker-compose up --build
```

После этого к прокси-серверу можно отправлять запросы вида:

http://localhost:2000/get_result?format=...

В качестве формата нужно передать один из параметров:
- native
- xml
- json
- protobuf
- avro
- yaml
- message_pack
