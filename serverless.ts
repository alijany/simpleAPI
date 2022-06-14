
import type { Serverless } from 'serverless/aws';

const serverlessConfiguration: Serverless = {
  app: 'aws-simple-api',
  service: 'aws-simple-api',
  provider: {
    name: 'aws',
    stage: 'dev',
    environment: {
      TABLE_PREFIX: '${self:service}-${opt:stage, self:provider.stage}-'
    },
    iam: {
      role: {
        statements: [{
          Effect: 'Allow',
          Action: [
            // 'dynamodb:DescribeTable',
            // 'dynamodb:Query',
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