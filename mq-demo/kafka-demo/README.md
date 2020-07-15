# README

先启动 zookeeper, 再启动 kafka

/usr/local/apache-zookeeper-3.5.6-bin/bin/zkServer.sh start

启动 kafka
nohup /usr/local/kafka_2.11-2.4.0/bin/kafka-server-start.sh /usr/local/kafka_2.11-2.4.0/config/server.properties >/usr/local/kafka_2.11-2.4.0/nohup.out 2>&1 &

停止 kafka
/usr/local/kafka_2.11-2.4.0/bin/kafka-server-stop.sh

创建一个topic
/usr/local/kafka_2.11-2.4.0/bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test

消费 topic
/usr/local/kafka_2.11-2.4.0/bin/kafka-console-consumer.sh --bootstrap-server 117.51.148.112:9092 --topic test --from-beginning

生产消息
/usr/local/kafka_2.11-2.4.0/bin/kafka-console-producer.sh --broker-list 117.51.148.112:9092 --topic test
