echo "====== Starting to Download Fabric ========"
curl -sSL https://bit.ly/2ysbOFE -o bootstrap.sh
chmod 755 ./bootstrap.sh
bash ./bootstrap.sh -- 2.2.3 1.5.0