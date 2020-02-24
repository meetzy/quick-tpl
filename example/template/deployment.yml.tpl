apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: {{.k8s.name}}
  namespace: {{.k8s.namespace}}
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: {{.k8s.name}}
  template:
    metadata:
      labels:
        k8s-app: {{.k8s.name}}
    spec:
      containers:
      - image: {{.docker.name}}:{{.docker.version}}
        imagePullPolicy: Always
        name: {{.k8s.name}}
        ports:
        - containerPort: 80
          protocol: TCP