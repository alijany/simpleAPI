import type { Serverless } from 'serverless/aws';

const deviceFunctions: NonNullable<Serverless['functions']> = {
    devices: {
        handler: 'bin/main',
        events: [
            { http: { method: 'post', path: '/api/devices', cors: true } },
            { http: { method: 'get', path: '/api/devices/{id}', cors: true } },
            { http: { method: 'get', path: '/api/temp/{id+}', cors: true } },
        ]
    }
}

export default deviceFunctions