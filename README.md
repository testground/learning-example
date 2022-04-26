# Testground learning example

This project is intended as a practical example of a "real" project, with its own internal business logic, dependencies, etc. It is referenced as a dependency by the [Learning-Example-Testground] project.

The logic and behaviors within this project are intended to be as straightforward as possible, acting as a reference / guide on how to implement these behaviors and test them using Testground.

## Getting started

```
git clone https://github.com/testground/learning-example
cd learning-example

docker-compose build && docker-compose up

# Assuming everything works, you should see the following message
Server started and listening at port 8081
```

## Tech

- [Docker]
- [Redis] 
- [RabbitMq]

## License

Dual-licensed: [MIT](./LICENSE-MIT), [Apache Software License v2](./LICENSE-APACHE), by way of the
[Permissive License Stack](https://protocol.ai/blog/announcing-the-permissive-license-stack/).


[//]: # (Reference links)

   [Learning-Example-Testground]: <https://github.com/testground/learning-example-tg>
   [Redis]: <https://www.redis.com/>
   [Docker]: <https://www.docker.com/>
   [RabbitMq]: <https://www.rabbitmq.com/>
