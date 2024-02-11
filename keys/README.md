# Keys

Place your key files inside this directory (`private.pem` and `public.pub`)

## Directory structure

```text
keys
├── private.pem
└── public.pub
```

## How to create a pair of key

```bash
  openssl genrsa -out ./keys/private.pem 2048 # Create Private Key
  openssl rsa -in ./keys/private.pem -pubout > ./keys/public.pub # Create Public Key
```
