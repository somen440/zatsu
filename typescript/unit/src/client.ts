import { obj } from './obj';

export const clientObj = (arg: string): string => `${arg}_${obj.prop()}`
