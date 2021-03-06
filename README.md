# Укорачиватель ссылок

Сервис, который предоставляет API по созданию сокращённых ссылок следующего формата:
- Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться
только одна сокращенная ссылка.
- Ссылка должна быть длинной 10 символов
- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем
регистре, цифр и символа _ (подчеркивание)

Сервис написан на Go и принимает следующие запросы по gRPC:
1. Метод Create, который сохраняет оригинальный URL в базе и возвращает сокращённый
2. Метод Get, который принимает сокращённый URL и возвращает оригинальный URL

Сервис распространен в виде Docker-образа. В качестве хранилища используется PostgreSQL
API описано в proto файле 

```
    ./api/proto/shorter.proto
```

### Настройка и запуск

(Использовать для первоначальной настройки проекта)

```bash
    make app-setup-and-up
```

### Запуск

(Использовать для запуска клиента)

```bash
    make run
```
