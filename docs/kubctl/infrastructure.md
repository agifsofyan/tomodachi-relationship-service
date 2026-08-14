# Arsitektur

A. Service:
  - Fastapi
    - port: 8000
    - host: localhost
    - prefix untuk:
      - api/v1/auth
      - api/v1/profiles
      - api/v1/interests
      - api/v1/addresses
      - api/v1/regions
      
  - Golang
    - port: 8080
    - host: localhost
    - prefix untuk:
      - api/v1/relationships/friends/request
      
B. Infrastucture:
  - Kubernetes
  - Kong
  - ArgoCD

## Architecture
```
                                 ┌────────────────────┐
                                 │      Client        │
                                 │ Postman / Browser  │
                                 └─────────┬──────────┘
                                           │
                                           │
                                 localhost:4001
                                           │
                                           ▼
                               ┌────────────────────┐
                               │        Kong        │
                               │--------------------│
                               │ Proxy :4001        │
                               │ Admin :4000        │
                               └─────────┬──────────┘
                                         │
                    ┌────────────────────┴─────────────────────┐
                    │                                          │
                    ▼                                          ▼
        ┌──────────────────────────┐              ┌──────────────────────────┐
        │      Auth Service        │              │ Relationship Service     │
        │--------------------------│              │--------------------------│
        │ FastAPI                  │              │ Go (Gin)                 │
        │ localhost:8000           │              │ localhost:8080           │
        │                          │              │                          │
        │ /api/v1/auth             │              │ /api/v1/relationships    │
        │ /api/v1/profiles         │              │ /friends/request         │
        │ /api/v1/interests        │              │                          │
        │ /api/v1/addresses        │              └──────────────────────────┘
        │ /api/v1/regions          │
        └──────────────────────────┘

──────────────────────────────────────────────────────────────────────────────

                        Kubernetes Cluster (k3d/k3s)

        Namespace : kong
            └── Kong Gateway

        Namespace : apps
            ├── auth-service
            └── relationship-service

        Namespace : argocd
            └── ArgoCD Server
```

## Request Flow
```
Client

POST /api/v1/auth/login
            │
            ▼
localhost:4001
            │
            ▼
Kong
            │
            ▼
auth-service:8000
```

## Kubernetes Namespace
```tree
apps
│
├── auth-service
└── relationship-service

kong
│
└── kong-gateway

argocd
│
└── argocd-server
```


## Infrastructure Repository
``` tree
micro-infra/
│
├── README.md
├── compose.yml
│
├── argocd/
│   ├── install.yml
│   ├── project.yml
│   └── applications/
│       ├── auth-service.yml
│       └── relationship-service.yml
│
├── kong/
│   ├── gateway.yml
│   ├── httproutes/
│   │   ├── auth.yml
│   │   └── relationship.yml
│   ├── plugins/
│   └── consumers/
│
├── namespaces/
│   ├── apps.yml
│   ├── kong.yml
│   └── argocd.yml
│
├── cluster/
│   ├── k3d.yml
│   └── bootstrap.sh
│
└── scripts/
    ├── start.sh
    ├── stop.sh
    ├── install-kong.sh
    └── install-argocd.sh
```

## Deployment Cycle
```
                GitHub

                   │

           micro-infra Repository

                   │

             ArgoCD Watch Repo

                   │

                   ▼

           Kubernetes Cluster

      ┌────────────┴────────────┐

      ▼                         ▼

    Kong                 Applications

                           │

          ┌────────────────┴───────────────┐

          ▼                                ▼

    auth-service              relationship-service
```

## Port yang Diekspos ke Host

| Komponen             |      Host Port | Keterangan              |
| -------------------- | -------------: | ----------------------- |
| Kong Proxy           |       **4001** | Entry point seluruh API |
| Kong Admin           |       **4000** | Administrasi Kong       |
| ArgoCD Server        |       **8081** | UI ArgoCD               |
| Auth Service         | Tidak diekspos | Diakses melalui Kong    |
| Relationship Service | Tidak diekspos | Diakses melalui Kong    |

## Roadmap Bertahap
1. Bootstrap cluster
   - cluster/k3d.yml
   - scripts/start.sh
   - scripts/stop.sh
2. Namespace
   - apps
   - kong
   - argocd
3. Kong
   - Install Kong
   - Expose Proxy:4001
   - Expose Admin:4000
4. ArgoCD
   - Install ArgoCD
   - Expose UI (misalnya localhost:8081)
   - Buat AppProject
5. Routing
   - auth-service
   - relationship-service
