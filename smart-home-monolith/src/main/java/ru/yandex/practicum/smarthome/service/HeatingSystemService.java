package ru.yandex.practicum.smarthome.service;

import ru.yandex.practicum.smarthome.dto.HeatingSystemDto;

public interface HeatingSystemService {
    Long newHeatingSystem(HeatingSystemDto heatingSystemDto);
    HeatingSystemDto getHeatingSystem(Long id);
    HeatingSystemDto updateHeatingSystem(Long id, HeatingSystemDto heatingSystemDto);
    void turnOn(Long id);
    void turnOff(Long id);
    void setTargetTemperature(Long id, double temperature);
    Double getCurrentTemperature(Long id);
}