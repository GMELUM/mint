```markdown
# TON Mint API

Этот проект предоставляет API для управления транзакциями с кошельками TON.

## Требования

- [Go 1.23.2](https://golang.org/dl/)
- Настроенные переменные окружения (все обязательны):

  - `PORT`: Номер порта, на котором будет слушать сервер.
  - `HOST`: Имя хоста или IP-адрес, к которому будет привязан сервер.
  - `WALLET_WORDS`: Слова для восстановления или генерации кошелька, разделенные пробелами.
  - `WALLET_DESTINATION`: Адрес кошелька для обработки транзакций.
  - `WALLET_JETTON`: Идентификатор токена или джеттона, используемого в транзакциях.
  - `SECRET`: Секретный ключ, необходимый для всех API-запросов, передаваемый через заголовок авторизации или как параметр запроса.
  - `MYSQL_HOST`: Хост MySQL базы данных.
  - `MYSQL_USERNAME`: Имя пользователя для MySQL.
  - `MYSQL_PASSWORD`: Пароль для MySQL.
  - `MYSQL_DATABASE`: Имя базы данных MySQL.
  - `MYSQL_PORT`: Порт MySQL сервера.
  - `MYSQL_MAX_CONNECTIONS`: Максимальное количество подключений к MySQL базе данных.
  - `MYSQL_CACHE_ENABLED`: Включение кэширования запросов для MySQL.
  - `MYSQL_MUTEX_ENABLED`: Включение Redis-базированного mutex для кэшированных данных MySQL.
  - `MYSQL_QUERY_DURATION`: Продолжительность запроса MySQL.
  - `CALLBACK_URL`: URL для обратных вызовов.

### Пример файла `.env`

Создайте файл `.env` в корневом каталоге:

```env
PORT=18300
HOST=0.0.0.0
WALLET_WORDS="your seed words here"
WALLET_DESTINATION=your_wallet_destination_address
WALLET_JETTON=your_jetton_identifier
SECRET=your_secret_key
MYSQL_HOST=localhost
MYSQL_USERNAME=root
MYSQL_PASSWORD=your_db_password
MYSQL_DATABASE=your_db_name
MYSQL_PORT=3306
MYSQL_MAX_CONNECTIONS=10
MYSQL_CACHE_ENABLED=false
MYSQL_MUTEX_ENABLED=false
MYSQL_QUERY_DURATION=1s
CALLBACK_URL=http://your_callback_url.com
```

## Установка и запуск

1. Клонируйте репозиторий и перейдите в каталог проекта:

   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   ```

2. Установите зависимости:

   ```bash
   go mod download
   ```

3. Запустите приложение:

   ```bash
   go run main.go
   ```

Сервер будет доступен по адресу `http://<HOST>:<PORT>`.

## Использование API

### Аутентификация

Все API-запросы должны включать `SECRET` для проверки подлинности. Это можно сделать следующими способами:
- **Заголовок авторизации:**

  Включите `SECRET` в заголовок авторизации HTTP:

  ```http
  Authorization: your_secret_key
  ```

- **Параметр запроса:**

  Включите `SECRET` как параметр запроса в URL:

  ```http
  http://<HOST>:<PORT>/withdraw?secret=your_secret_key
  ```

### Запрос на вывод средств

- **Маршрут:** `POST /withdraw`
- **Тело запроса:**

  В теле запроса должны быть указаны следующие поля JSON:

  ```json
  {
    "transaction": "transaction_detail",
    "wallet": "recipient_wallet_address",
    "amount": 1000,
    "message": "Transaction message" // необязательное сообщение для транзакции
  }
  ```

### Успешный ответ

- **Формат:**

  Изменен формат успешного ответа от сервера на запрос вывода средств. Теперь ответ будет выглядеть так:

  ```json
  {
    "response": {
      "result": true
    }
  }
  ```

### Обработка обратных вызовов

На указанный `CALLBACK_URL` отправляется объект следующего формата при успешной транзакции:

```json
{
  "id": 123,  // уникальный идентификатор
  "transaction": "transaction_detail",  // подробная информация о транзакции
  "hash": "LdSOGgjcvBuAPmCIEsL8Z48H8LvEiXXRFMxaeYSJeF4=",  // хеш транзакции
  "created_at": "2023-10-10T10:00:00Z",  // время создания объекта
  "updated_at": "2023-10-10T10:00:00Z"   // время последнего обновления объекта
}
```

### Пример обработки обратных вызовов

Функция обработки обратных вызовов может быть реализована следующим образом:

```go
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var success Success
		if err := json.NewDecoder(r.Body).Decode(&success); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		// Implement your logic here using the `success` object

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
```

### Ответ с ошибкой

Когда возникает ошибка, ответ будет содержать объект `error` в следующем формате:

- **Формат ошибки:**

  ```json
  {
    "error": {
      "code": 123,
      "message": "Error message",
      "critical": true
    }
  }
  ```
```
```