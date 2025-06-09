
{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "gokv.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
Default annotations for jnnkrdb github projects.
*/}}
{{- define "gokv.defaultAnnotations" -}}
jnnkrdb.de/source: "github.com/jnnkrdb/gokv"
{{- end }}

{{/*
Common labels 
*/}}
{{- define "gokv.labels" -}}
jnnkrdb.de/chart: {{ include "gokv.chart" . }}
helm.sh/chart: {{ include "gokv.chart" . }}
{{ include "gokv.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels 
*/}}
{{- define "gokv.selectorLabels" -}}
jnnkrdb.de/service: gokv
jnnkrdb.de/instance: {{ .Release.Name }}
{{- end }}


