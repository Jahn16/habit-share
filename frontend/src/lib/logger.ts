import { createLogger, transports, format } from 'winston';

export const logger = createLogger({
	level: 'info',
	format: format.json(),
	defaultMeta: { service: 'user-service' },
	transports: [new transports.Console({ format: format.simple() })]
});
