package ru.yandex.practicum.smarthome.kafka;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
//import util.properties packages
import java.util.Properties;

//import simple producer packages
import org.apache.kafka.clients.producer.Producer;

//import KafkaProducer packages
import org.apache.kafka.clients.producer.KafkaProducer;

//import ProducerRecord packages
import org.apache.kafka.clients.producer.ProducerRecord;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import ru.yandex.practicum.smarthome.dto.TelemetryRecordDto;
import ru.yandex.practicum.smarthome.dto.TelemetryValueDto;

public class TemperatureProducer {
    public static final String ENV_TOPIC_NAME = "TOPIC_NAME";
    public static final String ENV_TOPIC_NAME_DEFAULT = "temperature";
    public static final String ENV_TOPIC_KAFKA_ADDRESS = "KAFKA";
    public static final String ENV_TOPIC_KAFKA_ADDRESS_DEFAULT = "kafka:9092";

    private Producer<String, String> producer;
    private String topicName;

    public TemperatureProducer() {
        this.topicName = System.getenv().getOrDefault(ENV_TOPIC_NAME, ENV_TOPIC_NAME_DEFAULT);

        // create instance for properties to access producer configs
        Properties props = new Properties();

        // Assign address id
        props.put("bootstrap.servers",
                System.getenv().getOrDefault(ENV_TOPIC_KAFKA_ADDRESS, ENV_TOPIC_KAFKA_ADDRESS_DEFAULT));

        // Set acknowledgements for producer requests.
        props.put("acks", "all");

        // If the request fails, the producer can automatically retry,
        props.put("retries", 0);

        // Specify buffer size in config
        props.put("batch.size", 16384);

        // Reduce the no of requests less than 0
        props.put("linger.ms", 1);

        // The buffer.memory controls the total amount of memory available to the
        // producer for buffering.
        // props.put("buffer.memory", 33554432);

        props.put("key.serializer",
                "org.apache.kafka.common.serialization.StringSerializer");

        props.put("value.serializer",
                "org.apache.kafka.common.serialization.StringSerializer");

        this.producer = new KafkaProducer<String, String>(props);
    }

    public void Send(Long id, double temperature) throws JsonProcessingException {
        // Create message object
        TelemetryValueDto value = new TelemetryValueDto("temperature", String.valueOf(temperature));
        List<TelemetryValueDto> values = List.of(value);
        var recordDto = TelemetryRecordDto.builder().device("unknown").serial(String.valueOf(id)).data(values).build();

        // Dump message object to json
        var mapper = new ObjectMapper();
        var json = mapper.writeValueAsString(recordDto);

        // Send message
        ProducerRecord<String, String> kafkaRecord = new ProducerRecord<>(this.topicName, json);
        producer.send(kafkaRecord);
    }
}
