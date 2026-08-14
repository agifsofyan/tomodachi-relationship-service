## Ambil service
kubectl get svc -n tomodachi-app

NAME                   TYPE
auth-service           ClusterIP
relationship-service   ClusterIP

```bash
kubectl port-forward svc/auth-service \
  -n tomodachi-app \
  8000:8000
```

atau

```bash
kubectl exec -it deployment/auth-service \
-n tomodachi-app -- sh
```

atau dari kong