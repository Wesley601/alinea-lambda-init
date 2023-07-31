package main

const IndexFile = `{{- if .HasApiGateway }}
import { APIGatewayEvent } from 'aws-lambda';
{{- end }}
{{- if .HasSqs }}
import { SQSEvent } from 'aws-lambda';
{{- end }}
{{- if .HasEventBridge }}
import { EventBridgeEvent } from 'aws-lambda';
{{- end }}

import { {{ .HandlerName }} } from './{{ .LambdaFileName }}';

{{ if .HasApiGateway }}export const handler = async (event: APIGatewayEvent) => {
  const result = await {{ .HandlerName }}();

  return {
    statusCode: result.success ? 200 : 400,
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
    body: JSON.stringify(result.data),
  };
};{{ end }}
{{ if .HasSqs }}export const handlerSQS = async (event: SQSEvent) => {
  await Promise.all(
    event.Records.map(({ body }) => {
      const body = JSON.parse(body);

      console.log(body);

      return {{ .HandlerName }}(body);
    }),
  );

  return true;
};{{ end }}
{{ if .HasEventBridge }}export const handlerEventBridge = async (event: EventBridgeEvent<string, any>) => {
  const { detail } = event;

  const result = await {{ .HandlerName }}(detail);

  return result.success;
};{{ end}}
`

const HandlerFile = `export const {{ .HandlerName }} = () => {
  console.log('Hello World!');

  return {
    success: true,
    data: {
      message: 'Hello World!',
    },
  };
};
`

const SpecFile = `import {
  connectMemory,
  closeDB,
  clearDB,
} from '{{ .DomainName }}/common/test/mongo-test-server';
import { {{ .HandlerName }} } from './{{ .LambdaFileName }}';

describe('{{ .SpecName }}', () => {
  beforeAll(async () => await connectMemory());
  beforeEach(async () => await clearDB());
  afterAll(async () => await closeDB());
});
`

const PackageFile = `{
  "name": "{{ .DomainName }}/{{ .LambdaFileName }}",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "devDependencies": {},
  "dependencies": {
    "{{ .DomainName }}/common": "*"
  },
  "scripts": {
    "test": "vitest",
    "build": "tsc --project tsconfig.build.json"
  },
  "keywords": [],
  "author": "",
  "license": "ISC"
}
`
