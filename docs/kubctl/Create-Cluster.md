# Recreate Cluster

## 1.  Tomodachi CLuster
```bash
k3d cluster delete tomodachi
```
lalu
```bash
k3d cluster create tomodachi \
  --registry-use k3d-registry.localhost:5000 \
  -p "8000:80@loadbalancer" \
  -p "8443:443@loadbalancer"
```

lalu
```bash
docker ps
```

akan menjadi kira-kira seperti ini
```bash
k3d-tomodachi-serverlb

PORTS

0.0.0.0:8000->80/tcp
0.0.0.0:8443->443/tcp
0.0.0.0:51142->6443/tcp
```

Hasil akhir. Alur seperti ini:
```bash
Browser
      │
http://localhost:8000
      │
      ▼
k3d LoadBalancer
      │
      ▼
Kong Proxy
      │
      ├────────► auth-service
      │
      └────────► relationship-service
```
Ini adalah pola yang umum dipakai untuk development dengan k3d.

## 2. Install Argo CD
```bash
kubectl create namespace argocd

```bash
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```
atau

```bash
kubectl apply --server-side -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

jika error
```bash
kubectl delete namespace argocd
```

Check
```bash
kubectl get pods -n argocd
```
Tunggu sampai semuanya Running

## 3. Deploy Kong
```bash
kubectl get pods -n kong
kubectl get svc -n kong
```

### prepare kong jika belum ada
#### Apply namespace dasar
```bash
kubectl apply -f kubernetes/namespace.yaml
kubectl apply -f kubernetes/kong/namespace.yaml
```

#### Apply root Application (app-of-apps)
```bash
kubectl apply -f argocd/root-app.yaml
```

#### Pantau sinkronisasi
```bash
kubectl get applications -n argocd
```
Tunggu sampai kolom SYNC STATUS = Synced dan HEALTH STATUS = Healthy untuk keempat app (fastapi-service, golang-service, kong-controller, kong-routes). Bisa juga dipantau langsung dari ArgoCD UI (https://localhost:8443) — akan muncul 5 card: root app + 4 child app.

Kalau ada yang stuck di OutOfSync/Missing/Progressing lama, cek detail errornya:
```bash
kubectl describe application <nama-app> -n argocd
```

#### Setelah semua Synced, verifikasi resource benar-benar hidup
```bash
kubectl -n kong get pods
kubectl -n microservice-app get pods,svc,ingress
```

## 4. Deploy Auth Service
Pastikan muncul
```bash
kubectl get all -n tomodachi-app
```

## 5. Deploy Relationship Service

## 6. Verifikasi Ingress
```bash
kubectl get ingress -n tomodachi-app
```

harus muncul kira-kira:
```bash
NAME                           CLASS   HOSTS   ADDRESS
auth-service-ingress           kong    *
relationship-service-ingress   kong    *
```