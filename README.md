# How to start docker container

1. docker compose up --build -d
2. docker exec -it db_xyz sh
3. login to mysql in container and create schema studi_kasus_xyz
4. docker exec -i db_xyz mysql -u root -proot studi_kasus_xyz < src/migrations/dump.sql
5. app is on port 8001

- Use build.sh to copy binaries from container and save docker image

# NOTES 
- Uncomment and change src/Dockerfile on line below to use non root user on container
- (change userid as needed)
-   RUN adduser -u 1000 -D app
-   RUN chown -R app:app /usr/app
-   USER app

- Generate RSA private key (e.g https://cryptotools.net/rsagen) and copy it to src/certs/private.key

- default private.key and .env is committed on purpose for ease of use