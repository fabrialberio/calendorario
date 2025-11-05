## Calendorario

### Installazione
1. [Installa Docker](https://docs.docker.com/engine/install/)
2. [scarica i file](https://github.com/fabrialberio/calendorario/archive/refs/heads/main.zip) oppure clona il repository con Git:
```
git clone https://github.com/fabrialberio/calendorario
```
3. crea una chiave privata:
```
openssl genrsa -out private_key.pem 4096
```
4. crea una chiave pubblica corrispondente:
```
openssl rsa -in private_key.pem -out public_key.pem -outform PEM -pubout
```
5. crea un file `.env` con le credenziali necessarie per l'avvio, ad esempio:
```env
POSTGRES_USER: calendorario
POSTGRES_PASSWORD: example
POSTGRES_ROOT_PASSWORD: example
POSTGRES_DB: calendorario
POSTGRES_CONTAINER_NAME: calendorario-postgres
ADMIN_PASSWORD: example
```
6. avvia il server di sviluppo:
```
docker compose -f docker-compose.dev.yml up
```
7. puoi accedere al server su `http://localhost:80`, con il nome utente `amdin` e la password inserita nel file `.env`.
