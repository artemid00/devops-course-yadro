# Vault + k8s via External Secret Operator (ESO)

Данная инструция описывает подчлючение Vault к k8s, как централизованное хранилище секретов.

В данном примере описывается синхронизацию кредов для dockerhub из Vault в Kubernetes ImagePullSecret

---

### Установка

1. Необходимо добавить репозиторий ESO через helm:

```sh
help repo add external-secrets https://charts.external-secrets.io
helm repo update
```

2. Установка в отдельный namespace:

```sh
helm install external-secrets external-secrets/external-secrets -n external-secrets --create-namespace
```

3. Проверка

```sh
kubectl get ns | grep -i external-secrets
```

---

### Настройка Vault

0. Создать kv секрет
1. Включить метод аунтентификации `kubernetes`
2. Настроить конфиг `auth/kubernetes/config`, в котором указать ca.crt кластера и issuer - имя кластера, например,  https://kubernetes.default.svc.cluster.local
3. Создать policy 
4. Создать role

---

### Настройка ESO

В данной директории приведены 3 файла, которые описывают настройку ImagePullSecret как External Secret из Vautl kv storage

1. `cluster_secret_store.yaml`

В данном файле создается ресурс ClusterSecretStore, в которо описывается подключение к Vault kv secret, с методом аунтентификации `auth/kubernetes/dockerhub-role` от имени ServiceAccount external-secret

2. `cluster_role_tokenreview.yaml`

В данном файле для сервис аккаунта external-secrets дается право выполнять token-review в kubernetes API. Это необходимо для авторизации в Vault с помощью JWT токена

3. `external_secret.yaml`

В данном файле описывается External Secret:
- Имя секрета и неймспейс
- Из какого SecretStore брать секрет
- В каком формате и какие поля сохранять

ImagePullSecret описывается в файле .dockerconfigjson

---

### Архитектура работы
1. Pod external-secrets получает JWT своего ServiceAccount
2. Отправляет его в Vault
3. Vault вызывает Kubernetes TokenReview API
4. Kubernetes подтверждает валидность токена
5. Vault дает доступ к секрету
6. External Secrets создаёт Kubernetes Secret

---

### Проверка работы

```sh
kubectl describe clustersecretstore vault-store
```

Должно быть 

```sh
Ready

True
```

