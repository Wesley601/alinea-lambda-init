package main

const IndexFile = `import { APIGatewayEvent } from 'aws-lambda';
import { {{ .HandlerName }} } from './{{ .LambdaFileName }}';

export const handler = async (event: APIGatewayEvent) => {
  const result = await {{ .HandlerName }}Handler();

  return {
    statusCode: result.success ? 200 : 400,
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
    body: JSON.stringify(result.data),
  };
};`

const HandlerFile = `export const {{ .HandlerName }} = () => {
  console.log('Hello World!');

  return {
    success: true,
    data: {
      message: 'Hello World!',
    },
  };
};`

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
});`

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
}`
