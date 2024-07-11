{{- define "GolangAMQP.RabbitMQ.podLabels" -}}
app: {{ .Values.rabbitMQ.deployment.labels.app }}
deployment: {{ include "GolangAMQP.RabbitMQ.deploymentName" .}}
helmChart: {{ lower .Chart.Name }}
{{- end -}}

{{- define "GolangAMQP.producer.podLabels" -}}
app: {{.Values.producer.deployment.labels.app}}
deployment: {{ include "GolangAMQP.producer.deploymentName" .}}
helmChart: {{ lower .Chart.Name }}
{{- end -}}

{{- define "GolangAMQP.consumer.podLabels" -}}
app: {{.Values.consumer.labels.app}}
deployment: {{ .Values.consumer.deploymentName}}
helmChart: {{ lower .Chart.Name }}
{{- end -}}



{{- define "GolangAMQP.producer.deploymentName" -}}
{{lower .Chart.Name }}-{{ .Values.producer.deployment.namePrefix }}-deployment
{{- end -}}

{{- define "GolangAMQP.producer.envConfigMapName" -}}
{{ lower .Chart.Name }}-{{ .Values.producer.deployment.env.configMapNamePrefix }}-env-configmap
{{- end -}}

{{- define "GolangAMQP.producer.ingressName" -}}
{{lower .Chart.Name }}-{{ .Values.producer.ingress.namePrefix }}-ingress
{{- end -}}

{{- define "GolangAMQP.producer.serviceName" -}}
{{lower .Chart.Name }}-{{ .Values.producer.service.serviceNamePrefix }}-service
{{- end -}}



{{- define "GolangAMQP.RabbitMQ.deploymentName" -}}
{{lower .Chart.Name }}-{{ .Values.rabbitMQ.deployment.namePrefix }}-deployment
{{- end -}}

{{- define "GolangAMQP.RabbitMQ.serviceName" -}}
{{lower .Chart.Name }}-{{ .Values.rabbitMQ.service.serviceNamePrefix }}-service
{{- end -}}

{{- define "GolangAMQP.RabbitMQ.ingressName" -}}
{{lower .Chart.Name }}-{{ .Values.rabbitMQ.ingress.namePrefix }}-ingress
{{- end -}}


{{- define "GolangAMQP.persistentVolumes.pvName" -}}
{{lower .Chart.Name }}-{{ .Values.persistentVolume.namePrefix }}-pv
{{- end -}}

{{- define "GolangAMQP.persistentVolumes.pvcName" -}}
{{lower .Chart.Name }}-{{ .Values.persistentVolume.claimNamePrefix }}-pvc
{{- end -}}


{{- define "GolangAMQP.producer.env" -}}
envFrom:
- configMapRef:
    name: {{ include "GolangAMQP.producer.envConfigMapName" .}}
{{- end -}}

{{- define "GolangAMQP.producer.healthCheck" -}}
httpGet:
  path: {{ .Values.producer.deployment.healthPath }}
  port: {{ .Values.producer.deployment.env.API_BIND_PORT }}
initialDelaySeconds: {{ .Values.producer.deployment.initialDelaySeconds }}
periodSeconds: {{ .Values.producer.deployment.periodSeconds }}
{{- end -}}