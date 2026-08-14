## Menggunakan k3d untuk local
```bash
kubectl config current-context
k3d cluster delete dev-cluster

k3d registry create registry.localhost --port 6000

k3d cluster create tomodachi \
  --registry-use k3d-registry.localhost:6000 \
  -p "8000:80@loadbalancer" \
  -p "8080:8081@loadbalancer"

kubectl config current-context
kubectl cluster-info
kubectl get nodes
```

## Build & push image ke registry lokal k3d
```bash
docker build -t k3d-registry.localhost:6000/auth-service:latest ./auth-service
docker push k3d-registry.localhost:6000/auth-service:latest

docker build -t k3d-registry.localhost:6000/relationship-service:latest ./relationship-service
docker push k3d-registry.localhost:6000/relationship-service:latest
```

## Buat CLuster
### CLuster Tomodachi
```bash
k3d cluster create tomodachi \
  --registry-use k3d-registry.localhost:6000 \
  -p "8000:80@loadbalancer" \
  -p "8080:8081@loadbalancer"
```

### Check namespace
```bash
kubectl get namespace
```

seharusnya namespace **tomodachi-app** sudah ada. tinggal namespace **argocd**

### Namespace ArgoCD
```bash
kubectl create namespace argocd
```

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

## Verifikasi cluster
```bash
k3d cluster list
kubectl config current-context
kubectl cluster-info
kubectl get nodes
```

## Apply manifest
```bash
kubectl apply -f kubernetes/namespace.yaml
```

## Post forward ArgoCD ke server
```bash
kubectl -n argocd port-forward svc/argocd-server 8443:443
```

## Ambil password admin
```bash
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d && echo
```
atau via CLI
```bash
argocd login localhost:8443 --username admin --insecure
```

## Akses Kong
```bash
curl -i http://localhost:8000/api/v1/auth/health
curl -i http://localhost:8000/api/v1/relationships/friends/request
```

## Check apakah service sudah tersambung kong
```bash
# Pastikan Kong controller sudah jalan
kubectl -n kong get pods

# Pastikan Ingress terbaca oleh Kong dan sudah dapat address
kubectl -n microservice-app get ingress

# Detail lebih lengkap
kubectl -n microservice-app describe ingress fastapi-ingress
kubectl -n microservice-app describe ingress golang-ingress
```
Kalau kolom ADDRESS pada get ingress kosong terus, biasanya artinya Kong controller belum sinkron — cek log:
```bash
kubectl -n kong logs -l app.kubernetes.io/name=kong -c ingress-controller --tail=50
```

## Pastikan pod service hidup
```bash
kubectl -n microservice-app get pods
kubectl -n microservice-app get svc
```

## Check apakah ArcoCD sudah tersambung ke manifest
```bash
kubectl get applications -n argocd
```

## Detail per App
```bash
argocd app get fastapi-service
argocd app get golang-service
argocd app get kong-controller
argocd app get kong-routes
```

## Recreate Cluster
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