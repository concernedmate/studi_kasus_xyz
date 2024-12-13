echo copying to bin..
docker cp app_xyz:usr/app ./bin

echo creating docker images
docker save studi_kasus_xyz-app_xyz -o ./docker-image/image-app.tar
echo done
