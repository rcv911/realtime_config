# realtime_config

### Описание работы:

1. **Инициализация**: Конфигурация загружается из `etcd` при запуске приложения через метод `NewRealTimeConfig`. Параметры можно также подгрузить из YAML файла для начальной загрузки.
2. **Наблюдение за изменениями**: С помощью `Watch` метод следит за изменениями в `etcd`. Как только конфигурация обновляется, вызывается функция `updateConfig`, которая безопасно обновляет значения.
3. **Получение конфигурации**: Для получения текущей конфигурации используется метод `GetConfig`, который предоставляет потокобезопасный доступ к текущей версии конфигурации.
4. **Завершение**: Для корректного завершения работы вызывается метод `Close`, который завершает `watch` и закрывает соединение с `etcd`.

### TODO:

- Обновлять точечно по полям, не весб конфиг
- Попробовать атомики для полей
- Интерфейс для etcd
- Нормально DI настроить

### Cmd

```shell
docker exec -it etcd1 etcdctl --endpoints=http://etcd1:2379 endpoint health
docker exec -it etcd2 etcdctl --endpoints=http://etcd2:2379 endpoint health
docker exec -it etcd3 etcdctl --endpoints=http://etcd3:2379 endpoint health
```

```shell
docker exec -it etcd1 etcdctl --endpoints=http://etcd1:2379 put foo bar
docker exec -it etcd1 etcdctl --endpoints=http://etcd1:2379 get foo
```

```shell
docker network inspect etcd-net
```

```shell
docker network connect etcd-net etcd1
docker network connect etcd-net etcd2
docker network connect etcd-net etcd3
```

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// Настройки подключения
	endpoints := []string{"localhost:2379"}
	dialTimeout := 5 * time.Second

	// Подключаемся к etcd
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer client.Close()

	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Запись данных в etcd
	key := "foo"
	value := "bar"
	_, err = client.Put(ctx, key, value)
	if err != nil {
		log.Fatalf("Failed to put key: %v", err)
	}
	fmt.Printf("✅ Key written: %s -> %s\n", key, value)

	// Чтение данных из etcd
	resp, err := client.Get(ctx, key)
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}

	// Выводим значение
	if len(resp.Kvs) > 0 {
		fmt.Printf("🔍 Key found: %s -> %s\n", resp.Kvs[0].Key, resp.Kvs[0].Value)
	} else {
		fmt.Println("⚠️ Key not found")
	}
}

```

### Обновить конфиг полностью

- из папки `config`

```shell
docker exec -it etcd-single etcdctl --endpoints=http://localhost:2379 put /app/config "$(cat config_template_updated.yaml)"
```
