package ru.yandex.practicum.smarthome.dto;

import lombok.Data;

@Data
public class TelemetryValueDto {
    private final String name;
    private final String value;
}
