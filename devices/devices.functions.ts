import type { Serverless } from 'serverless/aws';

const deviceFunctions: NonNullable<Serverless['functions']> = {
    devices: {
        handler: 'devices',
        package: {
            patterns: ['bin/devices'] // TODO
        },
        events: [
            { http: { method: 'post', path: '/api/devices', cors: true } },
            { http: { method: 'get', path: '/api/devices/{id}', cors: true } },
        ]
    }
}

export default deviceFunctions