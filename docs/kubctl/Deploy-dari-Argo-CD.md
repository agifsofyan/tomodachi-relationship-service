# Tahap deploy service + Kong lewat ArgoCD.

## Pastikan placeholder sudah diganti dengan benar
Sebelum apply, cek dulu isi manifest kamu — terutama dua hal yang sempat kita bahas:
```bash
grep -r "<ORG>" argocd/
grep -r "<REGISTRY>" kubernetes/
```

Kalau masih muncul hasil, berarti masih ada placeholder yang belum diganti — perbaiki dulu:

<ORG> → URL repo Git kamu yang sebenarnya (yang sudah di-push)
<REGISTRY> → k3d-registry.localhost:5000 (port 5000, bukan 6000 — karena ini dibaca oleh node k3d di dalam cluster, bukan dari host kamu)

Kalau sudah diganti, commit & push ulang ke repo:
```bash
git add .
git commit -m "update registry and repo placeholders"
git push
```

## Apply namespace dasar
```bash
kubectl apply -f kubernetes/namespace.yaml
kubectl apply -f kubernetes/kong/namespace.yaml
```

## Apply root Application (app-of-apps)
```bash
kubectl apply -f argocd/root-app.yaml
```

## Pantau sinkronisasi
```bash
kubectl get applications -n argocd
```
Tunggu sampai kolom SYNC STATUS = Synced dan HEALTH STATUS = Healthy untuk keempat app (fastapi-service, golang-service, kong-controller, kong-routes). Bisa juga dipantau langsung dari ArgoCD UI (https://localhost:8443) — akan muncul 5 card: root app + 4 child app.

Kalau ada yang stuck di OutOfSync/Missing/Progressing lama, cek detail errornya:
```bash
kubectl describe application <nama-app> -n argocd
```

## Setelah semua Synced, verifikasi resource benar-benar hidup
```bash
kubectl -n kong get pods
kubectl -n microservice-app get pods,svc,ingress
```

## Test akses lewat Kong
```bash
curl -i http://localhost:8000/api/v1/auth/health
curl -i http://localhost:8000/api/v1/relationships/friends/request
```

## Hapus semua aplikasi
```bash
kubectl delete applications --all -n argocd
```

hapus namespace aplikasi
```bash
kubectl delete namespace tomodachi-app
kubectl delete namespace kong
kubectl delete namespace argocd
```

hapus cluster k3d
```bash
k3d cluster delete tomodachi
```

optional hapus registry local
```bash
docker stop k3d-registry.localhost
docker rm k3d-registry.localhost
```
atau
```bash
k3d registry delete registry.localhost
```

optional hapus volume ls
```bash
docker volume ls
```

hapus jika masih ada
```bash
docker volume rm k3d-tomodachi-images
```

bersihkan docker image
```bash
docker image rm \
k3d-registry.localhost:5000/auth-service \
k3d-registry.localhost:5000/relationship-service
```

atau hapus semua (riskan)
```bash
docker image prune -a
```

pastikan sudah bersih
```bash
k3d cluster list

kubectl config get-contexts

docker ps

docker volume ls

docker images
```

lalu urutan build ulang:
1. Buat registry local
   ```bash
   k3d registry create k3d-registry.localhost \
  --port 6000
  ```
2. k3d cluster create
   ```bash
   k3d cluster create tomodachi \
  --registry-use k3d-k3d-registry.localhost:6000 \
  -p "8181:80@loadbalancer" \
  -p "8443:443@loadbalancer" \
  -p "8001:8001@loadbalancer"
  ```
3.  Install ArgoCD
4.  Login ArgoCD
5.  Deploy Root App
6.  Deploy Kong
7.  Deploy Namespace
8.  Deploy Auth Service
9.  Deploy Relationship Service
10. Test Kong
11. Test API

Generate sealed secret
```bash
  kubectl create secret generic auth-service-secret \
  --namespace=tomodachi-app \
  --from-literal=DATABASE_URL='postgresql+psycopg://<USER>:<PASSWORD>@<HOST>:<PORT>/<DB_NAME>' \
  --from-literal=SECRET_KEY='<JWT_SECRET_KEY>' \
  --dry-run=client -o yaml | kubeseal --format=yaml \
  > kubernetes/application/auth-service/sealed-secret.yaml

  kubectl create secret generic relationship-service-secret \
  --namespace=tomodachi-app \
  --from-literal=DATABASE_HOST='<DB_HOST>' \
  --from-literal=DATABASE_PORT='<DB_PORT' \
  --from-literal=DATABASE_NAME='<DB_NAME' \
  --from-literal=DATABASE_USER='DB_USER' \
  --from-literal=DATABASE_PASSWORD='DB_PASS' \
  --from-literal=SECRET_KEY='<JWT_SECRET_KEY>' \
  --dry-run=client -o yaml | kubeseal --format=yaml \
  > kubernetes/application/relationship-service/sealed-secret.yaml
```

```bash
  kubectl create secret generic relationship-service-secret \
  --namespace=tomodachi-app \
  --from-literal=DATABASE_HOST='aws-0-ap-southeast-1.pooler.supabase.com' \
  --from-literal=DATABASE_PORT='5432' \
  --from-literal=DATABASE_NAME='tomodachi-relationship' \
  --from-literal=DATABASE_USER='postgres.vkagugtiiigehrvvmlpo' \
  --from-literal=DATABASE_PASSWORD='L11Vjip7SvznIT36BfZQibYkiXe9fKD' \
  --from-literal=SECRET_KEY='c2f68238b99ae64b45b5e7203345296c689a4ef3f1a709406f1db19adf9587a0' \
  --dry-run=client -o yaml | kubeseal --format=yaml \
  > kubernetes/application/relationship-service/sealed-secret.yaml

  ❯ kubectl get secret auth-service-secret -n tomodachi-app -o jsonpath='{.data.DATABASE_URL}' | base64 -d
postgresql+psycopg://<user>:<password>@<host>:5432/tomodachi%
```

## Check log Service
```bash
kubectl logs -f deployment/auth-service -n tomodachi-app
```