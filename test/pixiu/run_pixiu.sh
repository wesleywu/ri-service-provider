docker run --name pixiug-gateway -p 8888:8888 \
    -v conf.yaml:/etc/pixiu/conf.yaml \
    -v log.yml:/etc/pixiu/log.yml \
    dubbogopixiu/dubbo-go-pixiu:latest