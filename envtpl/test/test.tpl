[agent]
  ## Default data collection interval for all inputs
  interval = "{{ .INTERVAL | default "10s" }}"
  round_interval = {{ .ROUND_INTERVAL | default "true"  }}
  hostname = "{{ .HOSTNAME }}"

{{ if eq .OUTPUT_INFLUXDB_ENABLED "true" -}}
[[outputs.influxdb]]
{{- else -}}
# InfluxDB output is disabled
{{- end }}

{{ if .OUTPUT_KAFKA_ENABLED -}}
[[outputs.kafka]]
{{- else -}}
# kafka output is disabled
{{- end }}

Environment variables starting with TAG_:
{{ range $key, $value := environment "TAG_"  }}{{ $key }}="{{ $value }}"
{{ end -}}
