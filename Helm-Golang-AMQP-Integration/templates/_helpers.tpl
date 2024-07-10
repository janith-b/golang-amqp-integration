{{- define "GolangAMQP.RabbitMQ.podLabels" -}}
app: {{ .Values.rabbitMQ.labels.app }}
deployment: {{ .Values.rabbitMQ.deploymentName }}
helmChart: {{ lower .Chart.Name }}
{{- end -}}

{{- define "GolangAMQP.producer.podLabels" -}}
app: {{.Values.producer.labels.app}}
deployment: {{ .Values.producer.deploymentName}}
helmChart: {{ lower .Chart.Name }}
{{- end -}}

{{- define "GolangAMQP.consumer.podLabels" -}}
app: {{.Values.consumer.labels.app}}
deployment: {{ .Values.consumer.deploymentName}}
helmChart: {{ lower .Chart.Name }}
{{- end -}}



{{- define "GolangAMQP.producer.envConfigMapName" -}}
name: {{ lower .Chart.Name }}-{{ .Values.producer.env.configMapNamePrefix }}-env-config
{{- end -}}

{{- define "GolangAMQP.RabbitMQ.serviceName" -}}
name: {{lower .Chart.Name }}-{{ .Values.rabbitMQ.service.serviceNamePrefix }}-service
{{- end -}}



{{- define "GolangAMQP.producer.env" -}}
envFrom:
- configMapRef:
  {{ include "GolangAMQP.producer.envConfigMapName" . | indent 2 }}
{{- end -}}

{{- define "GolangAMQP.producer.healthCheck" -}}
httpGet:
  path: {{ .Values.producer.healthPath }}
  port: {{ .Values.producer.env.API_BIND_PORT }}
initialDelaySeconds: {{ .Values.producer.initialDelaySeconds }}
periodSeconds: {{ .Values.producer.periodSeconds }}
{{- end -}}