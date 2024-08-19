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

[C4 схема](./task_1.1.puml)

Язык: Java
СУБД: Postgresql

Система монолитная, с разделением на контроллер, репозиторий и сервис. Все компоненты системы находятся в одном сервисе.
Запросы обрабатываются синхронно.
Масштабировать можно только всю систему целиком.

Судя по требованиям к конечному продукту, команду разработки придется раширять. Предлогается часть команды оставить для поддержки legacy на время переезда на сервисы.
Для разработки сервисов придется нанимать людей, почему бы при этом не переехать на го? 😀

## Домены

### Smart Home

Контексты:
- **Heateheating control** - управляет температурой дома, отвечает на запросы пользователей
- **Temperature monitoring** - получает данные о температуре от датчиков

# Задание 1.2

Схемы:
- [Схема контейнеров](./task_1.2_containers.puml)
- [Схема компонент разбораданныех сенсора](./task_1.2_components.puml)
- [Последовательность обработки данных от датчиков](./task_1.2_seq_sensor_write.puml)

## Задание 1.3

- [Схема таблиц](./task_1.3_er.puml)

## Задание 1.4

- [swagger описание сервиса управления устройствами](./task_1.4_command_swagger.yaml)