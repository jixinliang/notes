# 1. 配置源

``` bash
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF


dnf clean packages
```

# 2. 安装软件

```bash
dnf install kubelet kubeadm kubectl -y

# 自动补全
dnf install bash-completion -y
echo "source <(kubectl completion bash)" >> ~/.bashrc
```

```bash
kubeadm config images list
kubectl version --client
```



