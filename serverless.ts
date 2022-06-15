
import type { Serverless } from 'serverless/aws';
// import deviceFunctions from './devices/devices.functions';

const serverlessConfiguration: Serverless = {
  app: 'aws-simple-api',
  service: 'aws-simple-api',
  functions: {
    devices: {
      handler: 'bin/main',
      events: [
        { http: { method: 'post', path: '/api/devices', cors: true } },
        { http: { method: 'get', path: '/api/devices/{id+}', cors: true } },
      ]
    }
  },
  package: {
    exclude: ["./**"],
    include: ["./bin/**"]
  },
  provider: {
    name: 'aws',
    stage: 'dev',
    runtime: "go1.x",
    region: 'us-east-1',
    environment: {
      TABLE_PREFIX: '${self:service}-${opt:stage, self:provider.stage}-'
    },
    iam: {
      role: {
        statements: [{
          Effect: 'Allow',
          Action: [
            'dynamodb:DescribeTable',
            'dynamodb:CreateTable',
            // 'dynamodb:Scan',
            'dynamodb:GetItem',
            'dynamodb:PutItem',
            // 'dynamodb:UpdateItem',
            // 'dynamodb:DeleteItem',

          ],
          // TODO
          Resource: 'arn:aws:dynamodb:${opt:region, self:provider.region}:*:table/${self:provider.environment.TABLE_PREFIX}*'
        }]
      }
    }
  },
};

module.exports = serverlessConfiguration;