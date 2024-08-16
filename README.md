# Базовая настройка

## Запуск minikube

[Инструкция по установке](https://minikube.sigs.k8s.io/docs/start/)

```bash
minikube start
```


## Добавление токена авторизации GitHub

[Получение токена](https://github.com/settings/tokens/new)

```bash
kubectl create secret docker-registry ghcr --docker-server=https://ghcr.io --docker-username=<github_username> --docker-password=<github_token> -n default
```


## Установка API GW kusk

[Install Kusk CLI](https://docs.kusk.io/getting-started/install-kusk-cli)

```bash
kusk cluster install
```


## Настройка terraform

[Установите Terraform](https://yandex.cloud/ru/docs/tutorials/infrastructure-management/terraform-quickstart#install-terraform)


Создайте файл ~/.terraformrc

```hcl
provider_installation {
  network_mirror {
    url = "https://terraform-mirror.yandexcloud.net/"
    include = ["registry.terraform.io/*/*"]
  }
  direct {
    exclude = ["registry.terraform.io/*/*"]
  }
}
```

## Применяем terraform конфигурацию 

```bash
cd terraform
terraform apply
```

## Настройка API GW

```bash
kusk deploy -i api.yaml
```

## Проверяем работоспособность

```bash
kubectl port-forward svc/kusk-gateway-envoy-fleet -n kusk-system 8080:80
curl localhost:8080/hello
```


## Delete minikube

```bash
minikube delete
```

----

# Задание 1.1

## Текущий функционал

- Получение нагревательной системы по ID
- Обновление нагревательной системы по ID
- Включение нагревательной системы по ID
- Выключение нагревательной системы по ID
- Установка температуры для нагревательной системы по ID (взаимодействие с датчиком температуры)
- Получение текущей температуры системы по ее ID

## Архитектура

Язык: Java
СУБД: Postgresql

Система монолитная, с разделением на контроллер, репозиторий и сервис. Все компоненты системы находятся в одном сервисе.
Запросы обрабатываются синхронно.
Масштабировать можно только всю систему целиком.

## Домены

### Дома (системы)

В данном домене есть следующие сущности:
- Система
- Пользователь

Можно выделить следующие функции:
- CRULD систем
- Управление пользователями
- Связь пользователя с системой
- Список подключенных к системе устройств (ID)

### Датчики

В данном домене есть следующие сущности:
- Устройство
- Температура

Можно выделить сдедующие функции:
- CRUDL устройств
- Прием данных с устройства
- Отображение данных выбранного устройства