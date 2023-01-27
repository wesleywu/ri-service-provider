docker run -p 8888:8888 \
    -v $(pwd)/conf.yaml:/etc/pixiu/conf.yaml \
    -v $(pwd)/log.yml:/etc/pixiu/log.yml \
    thehackercat/dubbo-go-pixiu-gateway:dubbo-go-pixiu-gateway-0.5.1-rc
