1. Completador de texto en docker:

https://github.com/huggingface/text-generation-inference/pkgs/container/text-generation-inference

docker run --gpus all --shm-size 1g -p 8080:80 -v $volume:/data ghcr.io/huggingface/text-generation-inference:1.4 --model-id $model

- Gpus a 0 para usar CPU
- $volume -> especificamos donde guardar los "pesos", si volume no existe lo crea.
- $model -> el modelo de generacion de texto que queremos cargar

ejemplo de ejecucion:
docker run --gpus 0 --shm-size 1g -p 8080:80 -v data_test:/data ghcr.io/huggingface/text-generation-inference:1.4 --model-id HuggingFaceH4/zephyr-7b-beta

ejemplo de consulta con CURL:
curl 127.0.0.1:8080/generate \
    -X POST \
    -d '{"inputs":"What is Deep Learning?","parameters":{"max_new_tokens":20}}' \
    -H 'Content-Type: application/json'