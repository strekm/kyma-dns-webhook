# Kyma DNS Webhook

The Kyma DNS Webhook is a component that can be used with [Cert Manager](https://cert-manager.netlify.com/).

//TODO @piotrmsc describe repo, configuration, how to build & test
## Build DNS webhook

 `docker build -f Dockerfile.webhook . `
 
 ## Sample Issuer and Cert CR for cert-manager
 
 ```
apiVersion: certmanager.k8s.io/v1alpha1
kind: Issuer
metadata:
  name: letsencrypt-staging
spec:
  acme:
    email: admin@kyma-project.io
    server: https://acme-staging-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-staging-account-key
    dns01:
      providers:
        - name: kyma-dns
          webhook:
            groupName: acme.kyma-project.io
            solverName: kyma-dns
            

                           
```                

```
apiVersion: certmanager.k8s.io/v1alpha1
kind: Certificate
metadata:
  name: sel-letsencrypt-crt
  namespace: default
spec:
  secretName: example-com-tls
  commonName: example.com
  dnsNames:
  - example.com
  - www.example.com
  issuerRef:
    name: letsencrypt-staging
    kind: Issuer
  acme:
    config:
      - dns01:
          provider: kymsa-dns
        domains:
          - example.com
          - www.example.com
          
```

## Installation

Helm charts are available in `deploy ` directory. Execute following command to install webhook on a cluster:

``` helm install ./deploy/kyma-dns-webhook  --name kyma-dns-webhook --namespace istio-system  ```
