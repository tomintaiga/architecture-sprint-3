package ru.yandex.practicum.smarthome.dto;

import java.util.List;

import lombok.Data;
import lombok.Builder;

@Builder
@Data
public class TelemetryRecordDto {
    private final String device;
    private final String serial;
    private final List<TelemetryValueDto> data;
}
